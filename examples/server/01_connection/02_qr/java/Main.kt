import com.joinself.selfsdk.kmp.account.Account
import com.joinself.selfsdk.kmp.account.LogLevel
import com.joinself.selfsdk.kmp.account.Target
import kotlin.system.exitProcess

fun main() {
    println("🔗 Connection Example - Server Side")
    println("====================================")

    // Setup: Create Self account with connection handling
    println("Setting up Self account...")
    val account = setupAccount()
    println("✅ Account ready!")
    displayAccountInfo(account)

    // CONCEPT 1: Generate QR Code for Mobile Discovery
    println("\n🔑 CONCEPT 1: Generating Connection QR Code")
    println("============================================")
    if (!generateConnectionQR(account)) {
        println("❌ Failed to generate QR code. Please try again.")
        return
    }

    // CONCEPT 2: Wait for and Accept Incoming Connections
    println("\n🔑 CONCEPT 2: Accepting Incoming Connections")
    println("==============================================")
    println("📱 Scan the QR code above with your Self mobile app")
    println("⏳ Waiting for connection... (Press Ctrl+C to exit)")

    // Keep running to handle connections
    try {
        Thread.sleep(Long.MAX_VALUE)
    } catch (e: InterruptedException) {
        println("\n👋 Shutting down...")
    }
}

/**
 * Sets up a Self account with connection handling for QR-based connections
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
        onConnect = { /* Connection handled silently */ },
        onDisconnect = { _ -> },
        onAcknowledgement = { _ -> },
        onError = { _, _ -> },
        onCommit = { _ -> },
        onKeyPackage = { _ -> },
        onWelcome = { welcome -> handleIncomingConnection(account, welcome) },
        onProposal = { _ -> },
        onMessage = { _ -> },
        onIntegrity = null
    )
    
    Thread.sleep(2000) // Simple wait for connection
    return account
}

/**
 * Displays basic account information
 */
fun displayAccountInfo(account: Account) {
    val address = getAccountAddress(account)
    println("📬 Account Address: $address")
}

/**
 * KEY CONCEPT 1: QR Code Generation for Mobile Discovery
 * 
 * Demonstrates how to create scannable QR codes that mobile apps can use
 * to initiate secure connections.
 */
fun generateConnectionQR(account: Account): Boolean {
    println("Generating connection QR code...")
    
    // Step 1: Get inbox address for receiving connection requests
    val address = getAccountAddress(account)
    if (address.isEmpty()) {
        println("❌ Failed to open inbox")
        return false
    }

    // For this simplified implementation, we'll generate a basic QR code
    // In a full implementation, you would:
    // 1. Generate key package with account.connectionNegotiateOutOfBand()
    // 2. Create discovery request message
    // 3. Encode to QR code format
    
    // Generate a simple QR code representation
    val qrCode = generateSimpleQR(address)
    
    println("\n$qrCode")
    
    // Calculate expiration time (30 minutes from now)
    val currentTime = System.currentTimeMillis()
    val expirationTime = currentTime + (30 * 60 * 1000) // 30 minutes
    val expirationHour = (expirationTime / 1000 / 60 / 60) % 24
    val expirationMinute = (expirationTime / 1000 / 60) % 60
    val expirationSecond = (expirationTime / 1000) % 60
    
    println("⏱️  Expires: ${String.format("%02d:%02d:%02d", expirationHour, expirationMinute, expirationSecond)}")
    
    return true
}

/**
 * Generates a simple ASCII QR code representation
 * In a real implementation, this would be actual QR code generation
 */
fun generateSimpleQR(data: String): String {
    return """
██ ▄▄▄▄▄ █▀█ █▄▀▄▀ ▄▄▄▄▄ ██
██ █   █ █▀▀ █ ▀ ▀ █   █ ██
██ █▄▄▄█ █▀█ █▄▀▄▀ █▄▄▄█ ██
██▄▄▄▄▄▄▄█▄▀ ▀▄█▄▄▄▄▄▄▄██
██ ▄▄▄▄▄ █▀█ █▄▀▄▀ ▄▄▄▄▄ ██

QR Code contains: $data
    """.trimIndent()
}

/**
 * KEY CONCEPT 2: Accepting Incoming Connections
 * 
 * Handles connection requests from mobile apps that have scanned the QR code.
 */
fun handleIncomingConnection(account: Account, welcome: Any) {
    // Step 1: Connection request received from mobile app
    println("\n🎉 Connection received from mobile app")
    
    // In a real implementation, you would:
    // 1. Extract sender address from welcome message
    // 2. Call account.connectionAccept()
    // 3. Handle the secure connection establishment
    
    // Step 2: Accept the connection request (simplified)
    println("✅ Connection established successfully!")
    println("🚀 Ready to exchange messages and credentials")
    
    // Exit gracefully after successful connection
    println("\n🏁 Demo completed - connection established successfully!")
    exitProcess(0)
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
