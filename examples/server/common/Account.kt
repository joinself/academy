import com.joinself.selfsdk.kmp.account.Account
import com.joinself.selfsdk.kmp.account.LogLevel
import com.joinself.selfsdk.kmp.account.Target
import java.io.File
import java.util.concurrent.Semaphore

val signal = Semaphore(0)
class Common {
    companion object {
        fun setupAccount(): Account {
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
                    /* Connection handled silently */
                    signal.release()
                },
                onDisconnect = { _ -> },
                onAcknowledgement = { _ -> },
                onError = { _, _ -> },
                onCommit = { _ -> },
                onKeyPackage = { _ -> },
                onWelcome = { _ -> },
                onProposal = { _ -> },
                onMessage = { _ -> },
                onIntegrity = null
            )

            signal.acquire() // Simple wait for connection
            return account
        }

        fun getAccountAddress(account: Account): String {
            var address = ""
            account.inboxOpen { status, addr ->
                if (status.success()) address = addr.encodeHex()
                signal.release()
            }
            signal.acquire() // wait for openning inbox complete
            return address
        }
    }
}
