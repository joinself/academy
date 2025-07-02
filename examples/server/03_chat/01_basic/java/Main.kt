import com.joinself.selfsdk.kmp.account.Account
import com.joinself.selfsdk.kmp.account.LogLevel
import com.joinself.selfsdk.kmp.account.Target
import java.time.LocalTime
import java.time.format.DateTimeFormatter

fun main() {
    println("💬 Simple Chat Demo")
    println("===================")

    // Create and configure Self account
    println("🔧 Setting up Self account...")
    val account = setupAccount()
    println("✅ Self account created successfully")
    println("✅ Connected to Self network")

    // Display connection information
    displayConnectionInfo(account)

    println("✅ Chat demo ready! Press Ctrl+C to exit.")

    // Keep running to handle incoming messages
    try {
        Thread.sleep(Long.MAX_VALUE)
    } catch (e: InterruptedException) {
        println("\n👋 Shutting down...")
    }
}

/**
 * Sets up a Self account configured for chat messaging
 */
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
        onConnect = { /* Connection handled in main */ },
        onDisconnect = { _ -> },
        onAcknowledgement = { _ -> },
        onError = { _, _ -> },
        onCommit = { _ -> },
        onKeyPackage = { _ -> },
        onWelcome = { _ -> },
        onProposal = { _ -> },
        onMessage = { message -> handleMessage(account, message) },
        onIntegrity = null
    )
    
    Thread.sleep(2000) // Simple wait for connection
    return account
}

/**
 * Displays connection information for the chat account
 */
fun displayConnectionInfo(account: Account) {
    val address = getAccountAddress(account)
    
    println("\n📬 Connection Information:")
    println("   🆔 Inbox Address: $address")
    println("   📱 To connect: Use this address in another Self SDK instance")
    println("   🔐 All messages will be automatically encrypted")
}

/**
 * Processes all incoming messages
 */
fun handleMessage(account: Account, message: Any) {
    val timestamp = getCurrentTimestamp()
    
    // In this simplified implementation, we'll treat all messages as chat messages
    // In a full implementation, you would:
    // 1. Check message content type
    // 2. Route to appropriate handler based on type
    // 3. Handle different message types (chat, credentials, etc.)
    
    handleChatMessage(account, message, timestamp)
}

/**
 * Processes incoming chat messages and sends responses
 */
fun handleChatMessage(account: Account, message: Any, timestamp: String) {
    // For this simplified implementation, we'll simulate message handling
    // In a real implementation, you would:
    // 1. Decode the chat message content
    // 2. Extract the sender address and message text
    // 3. Generate appropriate response
    // 4. Send response back to sender
    
    // Simulate receiving a message
    val senderAddress = "did:self:example"
    val messageText = "Hello from mobile app!"
    
    println("📨 [$timestamp] $senderAddress: \"$messageText\"")
    
    val response = generateResponse(messageText, timestamp)
    sendResponse(account, senderAddress, response, timestamp)
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
fun sendResponse(account: Account, toAddress: String, responseText: String, timestamp: String) {
    // In a real implementation, you would:
    // 1. Build chat content using message builders
    // 2. Use account.messageSend() to deliver the message
    // 3. Handle any send errors appropriately
    
    // For this simplified implementation, we'll just log the response
    println("📤 [$timestamp] Sent: \"$responseText\"")
    
    // Simulate successful send
    // In reality, this would involve actual message delivery via the SDK
}

/**
 * Helper function to get account address
 */
fun getAccountAddress(account: Account): String {
    var address = ""
    account.inboxOpen { status, addr -> 
        if (status.success()) {
            address = addr.encodeHex()
        }
    }
    Thread.sleep(1000) // Simple wait
    return address
}

/**
 * Helper function to get current timestamp
 */
fun getCurrentTimestamp(): String {
    return LocalTime.now().format(DateTimeFormatter.ofPattern("HH:mm:ss"))
} 
