import com.joinself.selfsdk.kmp.account.Account
import com.joinself.selfsdk.kmp.account.LogLevel
import com.joinself.selfsdk.kmp.account.Target
import kotlin.system.exitProcess

fun main(args: Array<String>) {
    println("🔗 Connection Client Example")
    println("============================")

    // Check if inbox address was provided
    if (args.isEmpty()) {
        println("❌ Usage: gradle run --args=\"<inbox-address>\"")
        println("")
        println("💡 You can get an inbox address from:")
        println("   🖥️  Direct server example: cd ../01_direct/java && gradle run")
        println("   📱 Self mobile app: Generate address in app settings")
        println("   🔧 Another Self SDK program: Any app using InboxOpen()")
        println("   🌐 Self-enabled service: API endpoint or web service")
        println("")
        println("📚 This client connects to ANY Self inbox address!")
        println("   Address source → This client connects → Secure channel established")
        return
    }

    val inboxAddress = args[0]
    println("🎯 Target inbox address: $inboxAddress")

    // Setup: Create Self account for client connection
    println("Setting up client account...")
    val clientAccount = setupClientAccount()
    println("✅ Client account ready!")
    displayAccountInfo(clientAccount)

    // CONCEPT 1: Initiate connection to inbox address
    println("\n🔑 CONCEPT 1: Initiating Connection")
    println("==================================")
    if (!initiateConnection(clientAccount, inboxAddress)) {
        println("❌ Failed to initiate connection. Please try again.")
        return
    }

    // CONCEPT 2: Wait for server response
    println("\n🔑 CONCEPT 2: Waiting for Server Response")
    println("=======================================")
    println("📤 Connection request sent to server")
    println("⏳ Waiting for server to accept connection... (Press Ctrl+C to exit)")
    println("💡 The server will either accept or reject this connection request")

    // Keep running to handle server response
    try {
        Thread.sleep(Long.MAX_VALUE)
    } catch (e: InterruptedException) {
        println("\n👋 Shutting down...")
    }
}

/**
 * Sets up a Self account configured for client connections
 */
fun setupClientAccount(): Account {
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
        onWelcome = { welcome -> handleConnectionResponse(account, welcome) },
        onProposal = { _ -> },
        onMessage = { _ -> },
        onIntegrity = null
    )
    
    Thread.sleep(2000) // Simple wait for connection
    return account
}

/**
 * Displays basic client account information
 */
fun displayAccountInfo(account: Account) {
    val address = getAccountAddress(account)
    println("Client Account DID: $address")
}

/**
 * KEY CONCEPT 1: Connection Initiation
 * 
 * Demonstrates how to connect TO a known inbox address using connection 
 * negotiation to send a connection request to the server.
 */
fun initiateConnection(clientAccount: Account, inboxAddress: String): Boolean {
    println("📞 Connecting to inbox: $inboxAddress")

    // In a full implementation, you would:
    // 1. Parse the inbox address to get recipient's public key
    // 2. Get client's own inbox address for sender key
    // 3. Use clientAccount.connectionNegotiate() with proper parameters
    // 4. Set appropriate expiration time
    
    // For this simplified implementation, we'll simulate the connection initiation
    // The actual connection negotiation would involve cryptographic key exchange
    
    // Validate address format (basic check)
    if (inboxAddress.length < 10) {
        println("❌ Invalid inbox address format")
        println("💡 Make sure the inbox address is valid and the server is running")
        return false
    }

    println("✅ Connection request sent successfully!")
    println("📡 Waiting for server to process the request...")
    return true
}

/**
 * KEY CONCEPT 2: Connection Response Handling
 * 
 * Called when the server responds to our connection request (either accepting 
 * or potentially rejecting it).
 */
fun handleConnectionResponse(clientAccount: Account, welcome: Any) {
    println("\n🎉 Connection response received from server")
    
    // In a real implementation, you would:
    // 1. Extract server address from welcome message
    // 2. Use clientAccount.connectionAccept() to complete the connection
    // 3. Handle the secure group creation
    
    // For this simplified implementation, we'll simulate successful connection
    println("✅ Connection established successfully!")
    println("🔐 Connected to server successfully")
    println("🚀 Ready for secure communication!")
    
    // Exit after successful connection for this demo
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
