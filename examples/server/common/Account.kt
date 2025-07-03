import com.joinself.selfsdk.kmp.account.Account
import com.joinself.selfsdk.kmp.account.LogLevel
import com.joinself.selfsdk.kmp.account.Target
import com.joinself.selfsdk.kmp.error.SelfStatus
import com.joinself.selfsdk.kmp.event.KeyPackage
import com.joinself.selfsdk.kmp.event.Message
import com.joinself.selfsdk.kmp.event.Welcome
import com.joinself.selfsdk.kmp.keypair.signing.PublicKey
import java.util.concurrent.Semaphore

class Common {
    interface Callbacks {
        fun onConnect() {}
        fun onMessage(account: Account, msg: Message) {}
        fun onKeyPackage(account: Account, keyPackage: KeyPackage) {
            account.connectionEstablish(asAddress =  keyPackage.toAddress(), keyPackage = keyPackage.keyPackage(),
                onCompletion = { status: SelfStatus, groupAddress: PublicKey ->
                    if (status.success()) {
                        println("✅ Successfully established encrypted connection: ${groupAddress.encodeHex()}")
                    } else {
                        println("❌ Failed to establish connection: ${status.errorMessage()}")
                    }
                }
            )
        }
        fun onWelcome(account: Account, welcome: Welcome) {
            account.connectionAccept(asAddress = welcome.toAddress(), welcome =  welcome.welcome()) { status: SelfStatus, groupAddress: PublicKey ->
                if (status.success()) println("✅ Connection established successfully!")
                else println("❌ Failed to accept connection")
            }
        }
    }
    companion object {
        fun setupAccount(callbacks: Callbacks? = null): Account {
            val signal = Semaphore(0)
            val storageKey = "276cb6191a345753adb0897c2c0a89370aebf44ef99e612747bee3cd4e757ffa"
                .chunked(2).map { it.toInt(16).toByte() }.toByteArray()

            val account = Account()
            account.configure(
                storagePath = "./storage",
                storageKey = storageKey,
                rpcEndpoint = Target.PRODUCTION_SANDBOX.rpcEndpoint(),
                objectEndpoint = Target.PRODUCTION_SANDBOX.objectEndpoint(),
                messageEndpoint = Target.PRODUCTION_SANDBOX.messageEndpoint(),
                logLevel = LogLevel.WARN,
                onConnect = {
                    signal.release()
                    callbacks?.onConnect()
                },
                onDisconnect = { _ -> println("account disconnected")},
                onAcknowledgement = { _ -> },
                onError = { _, status -> println("account error ${status.code()}:${status.errorMessage()}")},
                onCommit = { _ -> },
                onKeyPackage = { keyPackage -> callbacks?.onKeyPackage(account, keyPackage)},
                onWelcome = { welcome -> callbacks?.onWelcome(account, welcome)},
                onProposal = { _ -> },
                onMessage = { msg -> callbacks?.onMessage(account, msg) },
                onIntegrity = null
            )

            signal.acquire() // Simple wait for connection
            return account
        }

        fun getAccountAddress(account: Account): String {
            val signal = Semaphore(0)
            var address = ""
            account.inboxOpen { status, addr ->
                if (status.success()) address = addr.encodeHex()
                signal.release()
            }
            signal.acquire() // wait for opening inbox complete
            return address
        }

        /**
         * Displays basic account information
         */
        fun displayAccountInfo(account: Account, name: String) {
            val address = getAccountAddress(account)
            if (address.isEmpty()) {
                println("Failed to open $name inbox")
            } else {
                println("📬 $name Address: $address")
            }
        }
    }
}
