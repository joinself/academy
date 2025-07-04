import com.joinself.selfsdk.kmp.account.Account
import com.joinself.selfsdk.kmp.error.SelfStatus
import com.joinself.selfsdk.kmp.event.*
import com.joinself.selfsdk.kmp.keypair.signing.PublicKey
import com.joinself.selfsdk.kmp.message.DiscoveryRequestBuilder
import com.joinself.selfsdk.kmp.time.Timestamp
import java.util.concurrent.Semaphore

fun main() {
    val signal = Semaphore(0)
    println("🔗 Connection Example - Server Side")
    println("====================================")

    // Setup: Create a Self account with connection handling
    println("Setting up Self account...")
    val account = Common.setupAccount(callbacks = object: Common.Callbacks {
        override fun onWelcome(account: Account, welcome: Welcome) {
            handleIncomingConnection(account, welcome)
            signal.release()
        }
    })
    println("✅ Account ready!")
    Common.displayAccountInfo(account,"Server Account")

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
    signal.acquire()
    println("\nPress enter to exit")
    readln()
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
    val address = Common.getAccountAddress(account)
    if (address.isEmpty()) {
        println("❌ Failed to open inbox")
        return false
    }

    val expires = Timestamp.now() + 3600

    // Step 2: Generate cryptographic key package for secure communication
    val keyPackage = account.connectionNegotiateOutOfBand(PublicKey.decodeHex(address), expires)

    // Step 3: Build discovery request message
    val discoveryRequest = DiscoveryRequestBuilder()
        .keyPackage(keyPackage)
        .expires(expires)
        .finish()

    // Step 4: Create anonymous message and encode to QR
    val anonymousMessage = AnonymousMessage.fromContent(discoveryRequest)
    anonymousMessage.setFlags(FlagSet(Flag.TARGET_SANDBOX))

    val qrCodeBytes = anonymousMessage.encodeQR(QrEncoding.UNICODE)
    if (qrCodeBytes.isEmpty()) {
        println("❌ Failed to generate QR code")
        return false
    }

    // Display the scannable QR code
    val qrCodeString = qrCodeBytes.decodeToString()
    println("\n$qrCodeString")
    
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
 * KEY CONCEPT 2: Accepting Incoming Connections
 * 
 * Handles connection requests from mobile apps that have scanned the QR code.
 * Key steps:
 *  1. Receive a Welcome message (connection request from mobile app)
 *  2. Accept the connection using ConnectionAccept
 *  3. Establish encrypted communication channel
 */
fun handleIncomingConnection(account: Account, welcome: Welcome) {
    val signal = Semaphore(0)
    // Step 1: Connection request received from mobile app
    println("\n🎉 Connection received from mobile app ${welcome.fromAddress().encodeHex()}")

    // Step 2: Accept the connection request
    // This is the critical call that establishes the secure connection
    account.connectionAccept(asAddress = welcome.toAddress(), welcome =  welcome.welcome()) { status: SelfStatus, groupAddress: PublicKey ->
        if (!status.success()) {
            println("❌ Failed to accept connection")
        }
        signal.release()
    }
    signal.acquire()

    // Step 3: Connection established - ready for secure communication
    println("✅ Connection established successfully!")
    println("🚀 Ready to exchange messages and credentials")
    
    // Exit gracefully after successful connection
    println("\n🏁 Demo completed - connection established successfully!")
}
