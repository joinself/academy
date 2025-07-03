import com.joinself.selfsdk.kmp.account.Account
import com.joinself.selfsdk.kmp.account.LogLevel
import com.joinself.selfsdk.kmp.account.Target
import java.io.File
import java.util.concurrent.Semaphore

val signal = Semaphore(0)

fun getAccountAddress(account: Account): String {
    var address = ""
    account.inboxOpen { status, addr ->
        if (status.success()) address = addr.encodeHex()
        signal.release()
    }
    signal.acquire() // wait for openning inbox complete
    return address
}