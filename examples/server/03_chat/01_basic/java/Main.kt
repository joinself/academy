import com.joinself.selfsdk.kmp.account.Account
import com.joinself.selfsdk.kmp.account.LogLevel
import com.joinself.selfsdk.kmp.account.Target
import com.joinself.selfsdk.kmp.event.Message
import com.joinself.selfsdk.kmp.keypair.signing.PublicKey
import com.joinself.selfsdk.kmp.message.Chat
import com.joinself.selfsdk.kmp.message.ChatBuilder
import com.joinself.selfsdk.kmp.message.ContentType
import java.time.LocalTime
import java.time.format.DateTimeFormatter

fun main() {
    println("💬 Simple Chat Demo")
    println("===================")

    // Create and configure Self account
    println("🔧 Setting up Self account...")
    val account = Common.setupAccount(callbacks = object: Common.Callbacks {
        override fun onMessage(account: Account, msg: Message) {
            handleMessage(account, msg)
        }
    })
    println("✅ Self account created successfully")
    println("✅ Connected to Self network")

    Common.displayAccountInfo(account, "Chat Account")

    println("✅ Chat demo ready! Press Ctrl+C to exit.")

    // Keep running to handle incoming messages
    try {
        Thread.sleep(Long.MAX_VALUE)
    } catch (e: InterruptedException) {
        println("\n👋 Shutting down...")
    }
}

/**
 * Processes all incoming messages
 */
fun handleMessage(account: Account, message: Message) {
    val timestamp = getCurrentTimestamp()
    val type = message.content().contentType()
    when(type) {
        ContentType.CHAT -> {
            handleChatMessage(account, message, timestamp)
        }
        else -> {
            println("📨 [$timestamp] Received message $type from ${message.fromAddress().encodeHex()}")
        }
    }
}

/**
 * Processes incoming chat messages and sends responses
 */
fun handleChatMessage(account: Account, msg: Message, timestamp: String) {
    val chat = Chat.decode(msg.content())
    println("📨 [$timestamp] ${msg.fromAddress().encodeHex()}: ${chat.message()}")

    val response = generateResponse(chat.message(), timestamp)
    sendResponse(account, msg.fromAddress(), response, timestamp)
}

/**
 * Creates appropriate responses based on the message content
 */
fun generateResponse(messageText: String, timestamp: String): String {
    val message = messageText.lowercase().trim()
    
    return when {
        message.contains("hello") || message.contains("hi") -> {
            "👋 Hello! Message received at $timestamp"
        }
        message.contains("how are you") -> {
            "🤖 I'm doing great! I'm a Self SDK chat demo."
        }
        message.contains("help") -> {
            "💡 I'm a simple chat bot. Try saying 'hello', 'how are you', or send any message for an echo!"
        }
        message.contains("time") -> {
            "🕐 Current time is $timestamp"
        }
        else -> {
            "🔄 Echo: $messageText"
        }
    }
}

/**
 * Sends a chat response to the peer
 */
fun sendResponse(account: Account, toAddress: PublicKey, responseText: String, timestamp: String) {
    val chat = ChatBuilder()
        .message(responseText)
        .finish()
    val sendStatus = account.messageSend(toAddress, chat)
    if (sendStatus.success()) {
        println("📤 [$timestamp] Sent: \"$responseText\"\n")
    } else {
        println("❌ [$timestamp] Failed to send response: ${sendStatus.errorMessage()}")
    }
}

/**
 * Helper function to get current timestamp
 */
fun getCurrentTimestamp(): String {
    return LocalTime.now().format(DateTimeFormatter.ofPattern("HH:mm:ss"))
} 
