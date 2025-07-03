import com.joinself.selfsdk.kmp.account.Account
import com.joinself.selfsdk.kmp.account.LogLevel
import com.joinself.selfsdk.kmp.account.Target
import com.joinself.selfsdk.kmp.event.KeyPackage
import com.joinself.selfsdk.kmp.event.Message
import com.joinself.selfsdk.kmp.event.Welcome
import java.util.concurrent.Semaphore

class Common {
    interface Callbacks {
        fun onConnect() {}
        fun onMessage(msg: Message) {}
        fun onKeyPackage(account: Account, keyPackage: KeyPackage) {}
        fun onWelcome(account: Account, welcome: Welcome) {}
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
                onDisconnect = { _ -> },
                onAcknowledgement = { _ -> },
                onError = { _, _ -> },
                onCommit = { _ -> },
                onKeyPackage = { keyPackage -> callbacks?.onKeyPackage(account, keyPackage)},
                onWelcome = { welcome -> callbacks?.onWelcome(account, welcome)},
                onProposal = { _ -> },
                onMessage = { msg -> callbacks?.onMessage(msg)},
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
            signal.acquire() // wait for openning inbox complete
            return address
        }

        /**
         * Displays basic account information
         */
        fun displayAccountInfo(account: Account) {
            val address = getAccountAddress(account)
            println("Server Account DID: $address")
        }
    }
}
