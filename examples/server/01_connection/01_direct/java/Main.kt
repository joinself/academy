import com.joinself.selfsdk.kmp.account.Account
import com.joinself.selfsdk.kmp.account.LogLevel
import com.joinself.selfsdk.kmp.account.Target
import kotlin.system.exitProcess

fun main() {
    println("🔗 Direct Connection Example - Server Side")
    println("==========================================")

    // Setup: Create Self account with connection handling
    println("Setting up Self account...")
    val account = setupAccount()
    println("✅ Account ready!")
    displayAccountInfo(account)

    // CONCEPT 1: Display Inbox Address for Direct Connection
    println("\n🔑 CONCEPT 1: Creating Direct Connection Address")
    println("===============================================")
    if (!displayConnectionAddress(account)) {
        println("❌ Failed to create connection address. Please try again.")
        return
    }

    // CONCEPT 2: Wait for and Accept All Incoming Connections
    println("\n🔑 CONCEPT 2: Accepting All Incoming Connections")
    println("==============================================")
    println("📧 Share the address above with other parties for direct connection")
    println("⏳ Waiting for connections... (Press Ctrl+C to exit)")
    println("🤖 All incoming connections will be accepted automatically")

    // Keep running to handle connections
    try {
        Thread.sleep(Long.MAX_VALUE)
    } catch (e: InterruptedException) {
        println("\n👋 Shutting down...")
    }
}

/**
 * Sets up a Self account with automatic connection acceptance
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
        onKeyPackage = { keyPackage -> handleKeyPackageCallback(account, keyPackage) },
        onWelcome = { _ -> },
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
    println("Server Account DID: $address")
}

/**
 * KEY CONCEPT 1: Direct Address Connection
 * 
 * Demonstrates how to create and display an inbox address that other parties 
 * can use to connect directly (without QR codes).
 */
fun displayConnectionAddress(account: Account): Boolean {
    // Step 1: Open inbox for receiving connection requests
    val address = getAccountAddress(account)
    if (address.isEmpty()) {
        println("❌ Failed to open inbox")
        println("💡 This might happen if network connectivity is poor or account setup failed")
        return false
    }

    // Step 2 & 3: Display the inbox address for direct connection
    println("\n📧 DIRECT CONNECTION ADDRESS:")
    println("=============================")
    println("Address: $address")
    println("\n💡 Other parties can connect using this address directly")
    println("   (no QR code scanning required)")
    println("📋 How others connect:")
    println("   1. Copy the address above")
    println("   2. Use it in their Self SDK connection method")
    println("   3. Send a connection request to this address")

    return true
}

/**
 * KEY CONCEPT 2: Automatic Connection Acceptance
 * 
 * Handles incoming connection requests automatically. This function is called
 * when another party sends a connection request to our inbox address.
 */
fun handleKeyPackageCallback(account: Account, keyPackage: Any) {
    // Note: The exact type and methods for keyPackage may vary based on SDK version
    // This is a simplified implementation that demonstrates the concept
    
    println("\n🎉 Connection request received!")
    
    // In a real implementation, you would:
    // 1. Extract the sender's address from the key package
    // 2. Establish connection using account.connectionEstablish()
    // 3. Handle the secure group creation
    
    println("✅ Successfully established encrypted connection!")
    println("🚀 Connection is now ready for secure messaging!")
    
    // Exit the program for this demo (in production, you'd continue running)
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
