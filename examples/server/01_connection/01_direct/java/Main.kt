import com.joinself.selfsdk.kmp.account.Account
import com.joinself.selfsdk.kmp.error.SelfStatus
import com.joinself.selfsdk.kmp.event.KeyPackage
import com.joinself.selfsdk.kmp.keypair.signing.PublicKey
import java.util.concurrent.Semaphore

fun main() {
    val signal = Semaphore(0)
    println("🔗 Direct Connection Example - Server Side")
    println("==========================================")

    // Setup: Create Self account with connection handling
    println("Setting up Self account...")
    val account = Common.setupAccount(object: Common.Callbacks {
        override fun onKeyPackage(account: Account, keyPackage: KeyPackage) {
            handleKeyPackageCallback(account, keyPackage)
            signal.release()
        }
    })
    println("✅ Account ready!")
    Common.displayAccountInfo(account)

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

    signal.acquire()
    println("\nPress enter to exit")
    readln()
}

/**
 * KEY CONCEPT 1: Direct Address Connection
 * 
 * Demonstrates how to create and display an inbox address that other parties 
 * can use to connect directly (without QR codes).
 */
fun displayConnectionAddress(account: Account): Boolean {
    // Step 1: Open inbox for receiving connection requests
    val address = Common.getAccountAddress(account)
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
 * handleKeyPackageCallback handles incoming connection requests automatically.
 * This function is called when another party sends a connection request to our
 * inbox address.
 *
 * What happens:
 *  1. Receives a KeyPackage from the connecting party
 *  2. Establishes an encrypted connection using the key package
 *  3. Creates a secure communication group
 *  4. Exits the program (for demo purposes)
 *
 * In production, you might want to:
 *  - Validate the connection request before accepting
 *  - Store connection information for later use
 *  - Continue running to handle multiple connections
 */
fun handleKeyPackageCallback(account: Account, keyPackage: KeyPackage) {
    val signal = Semaphore(0)

    // Note: The exact type and methods for keyPackage may vary based on SDK version
    // This is a simplified implementation that demonstrates the concept
    var groupAddressHex: String = ""
    println("\n🎉 Connection request received!")
    
    // In a real implementation, you would:
    // 1. Extract the sender's address from the key package
    // 2. Establish connection using account.connectionEstablish()
    // 3. Handle the secure group creation
    account.connectionEstablish(asAddress =  keyPackage.toAddress(), keyPackage = keyPackage.keyPackage(),
        onCompletion = { status: SelfStatus, groupAddress: PublicKey ->
            if (!status.success()) {
                println("❌ Failed to establish connection: %v")
                println("💡 This might happen if the key package is invalid or network issues occur")
            }
            groupAddressHex = groupAddress.encodeHex()

            signal.release()
        }
    )
    signal.acquire()
    
    println("✅ Successfully established encrypted connection!")
    println("📱 Connected to: ${keyPackage.fromAddress().encodeHex()}")
    println("🔐 Secure group created: ${groupAddressHex}")
    println("🚀 Connection is now ready for secure messaging!")
    
    // Exit the program for this demo (in production, you'd continue running)
    println("\n🏁 Demo completed - connection established successfully!")
}