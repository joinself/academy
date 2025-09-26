import com.joinself.selfsdk.account.Account
import com.joinself.selfsdk.account.LogLevel
import com.joinself.selfsdk.account.Target
import com.joinself.selfsdk.error.SelfStatus
import com.joinself.selfsdk.event.KeyPackage
import com.joinself.selfsdk.event.Message
import com.joinself.selfsdk.event.Welcome
import com.joinself.selfsdk.keypair.signing.PublicKey
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
        fun setupAccount(storagePath: String = "./storage", callbacks: Callbacks? = null): Account {
            val signal = Semaphore(0)
            val storageKey = "276cb6191a345753adb0897c2c0a89370aebf44ef99e612747bee3cd4e757ffa"
                .chunked(2).map { it.toInt(16).toByte() }.toByteArray()

            val account = Account()
            account.configure(
                storagePath = storagePath,
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
                onDropped = {_ ->},
                onIntegrity = null
            )

            signal.acquire() // Simple wait for connection
            return account
        }

        fun openInbox(account: Account): PublicKey? {
            val signal = Semaphore(0)
            var inboxAddress: PublicKey? = null
            account.inboxOpen(expires = 0L) { status, addr ->
                if (status.success()) inboxAddress = addr
                signal.release()
            }
            signal.acquire() // wait for opening inbox complete
            return inboxAddress
        }

        fun getAccountAddress(account: Account): String {
            val signal = Semaphore(0)
            var address = ""
            account.inboxOpen(expires = 0L) { status, addr ->
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
