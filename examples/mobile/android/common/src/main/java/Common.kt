import com.joinself.sdk.models.Account
import com.joinself.sdk.models.ChatMessage
import com.joinself.sdk.models.PublicKey
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch

class Common {
    companion object {
        // connect with server by an inbox address, a group address is returned.
        suspend fun connect(account: Account, qrCode: ByteArray): PublicKey? {
            try {
                return account.connectWith(qrCode)
            } catch (ex: Exception) {
                return null
            }
        }

        fun notifyServerForRequest(account: Account, groupAddress: PublicKey, message: String) = CoroutineScope(Dispatchers.IO).launch {
            val chat = ChatMessage.Builder()
                .setMessage(message)
                .build()
            account.send(toAddress = groupAddress, chat)
        }
    }
}
