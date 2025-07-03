import com.joinself.selfsdk.kmp.account.Account
import com.joinself.selfsdk.kmp.error.SelfStatus
import com.joinself.selfsdk.kmp.event.Welcome
import com.joinself.selfsdk.kmp.keypair.signing.PublicKey
import com.joinself.selfsdk.kmp.time.Timestamp
import java.util.concurrent.Semaphore
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

    val targetInboxAddress = args[0]
    println("🎯 Target inbox address: $targetInboxAddress")

    // Setup: Create Self account for client connection
    println("Setting up client account...")
    val clientAccount = Common.setupAccount(object: Common.Callbacks {
        override fun onWelcome(account: Account, welcome: Welcome) {
            handleConnectionResponse(account, welcome)
        }
    })
    println("✅ Client account ready!")
    Common.displayAccountInfo(clientAccount, "Client Account")

    // CONCEPT 1: Initiate connection to inbox address
    println("\n🔑 CONCEPT 1: Initiating Connection")
    println("==================================")
    if (!initiateConnection(clientAccount, targetInboxAddress)) {
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
 * KEY CONCEPT 1: Connection Initiation
 * 
 * Demonstrates how to connect TO a known inbox address using connection 
 * negotiation to send a connection request to the server.
 */
fun initiateConnection(clientAccount: Account, targetInboxAddress: String): Boolean {
    val signal = Semaphore(0)
    println("📞 Connecting to inbox: $targetInboxAddress")

    // Validate address format (basic check)
    if (targetInboxAddress.length < 10) {
        println("❌ Invalid inbox address format")
        println("💡 Make sure the inbox address is valid and the server is running")
        return false
    }

    // Get our own inbox address and convert to public key for sender
    val inboxAddress = Common.getAccountAddress(account = clientAccount)

    // Use ConnectionNegotiate to initiate connection to the inbox address
    var isSuccess = false
    clientAccount.connectionNegotiate(asAddress = PublicKey.decodeHex(inboxAddress), withAddress = PublicKey.decodeHex(targetInboxAddress), expires = Timestamp.now() + 360) { status: SelfStatus ->
        if (!status.success()) {
            isSuccess = false
            println("❌ Failed to initiate connection")
            println("💡 Make sure the inbox address is valid and the server is running")
        } else {
            isSuccess = true
            println("✅ Connection request sent successfully!")
            println("📡 Waiting for server to process the request...")
        }
        signal.release()
    }
    signal.acquire()

    return isSuccess
}

/**
 * KEY CONCEPT 2: Connection Response Handling
 * 
 * Called when the server responds to our connection request (either accepting 
 * or potentially rejecting it).
 */
fun handleConnectionResponse(clientAccount: Account, welcome: Welcome) {
    val signal = Semaphore(0)
    println("\n🎉 Connection response received from server: ${welcome.fromAddress().encodeHex()}")

    // Accept the connection establishment from the server
    clientAccount.connectionAccept(asAddress = welcome.toAddress(), welcome =  welcome.welcome()) { status: SelfStatus, groupAddress: PublicKey ->
        if (!status.success()) {
            println("❌ Failed to accept connection")
        } else {
            // Connection successful!
            println("✅ Connection established successfully!")
            println("🔐 Connected to server successfully")
            println("🚀 Ready for secure communication!")
        }
        signal.release()
    }
    signal.acquire()

    // Exit after successful connection for this demo
    println("\n🏁 Demo completed - connection established successfully!")
    exitProcess(0)
}
