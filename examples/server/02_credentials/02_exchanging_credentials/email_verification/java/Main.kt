import com.joinself.selfsdk.kmp.account.Account
import com.joinself.selfsdk.kmp.credential.Address
import com.joinself.selfsdk.kmp.credential.CredentialBuilder
import com.joinself.selfsdk.kmp.credential.VerifiableCredential
import com.joinself.selfsdk.kmp.error.SelfStatus
import com.joinself.selfsdk.kmp.event.*
import com.joinself.selfsdk.kmp.keypair.signing.PublicKey
import com.joinself.selfsdk.kmp.message.DiscoveryRequestBuilder
import com.joinself.selfsdk.kmp.time.Timestamp
import java.time.LocalDateTime
import java.time.format.DateTimeFormatter
import java.util.concurrent.Semaphore

/**
 * Main entry point for the email verification demo.
 *
 * This application acts as an "Email Service Provider" that:
 * 1. Generates a QR code for a user to scan with their Self mobile app.
 * 2. Waits for a secure connection to be established.
 * 3. Issues a verifiable credential for a verified email address to the connected mobile device.
 */
fun main() {
    println("📧 Email Credential Verification Demo")
    println("====================================")

    // A semaphore to signal when the entire workflow is complete, allowing the app to exit.
    val completionSignal = Semaphore(0)

    // 1. Create the email service provider's account with a callback to handle incoming connections.
    val emailService = createEmailServiceAccount(completionSignal)
    Common.displayAccountInfo(emailService, "Email Service Provider")

    // 2. Generate a QR code for the mobile app to scan.
    if (!generateEmailVerificationQR(emailService)) {
        println("❌ Failed to generate QR code. Please try again.")
        return
    }

    // 3. Wait for the mobile connection and subsequent credential issuance.
    println("\n📱 SCAN QR CODE with Self mobile app to verify email")
    println("⏳ Waiting for mobile device connection...")
    println("🔐 Once connected, an email verification credential will be created and sent.")

    // Wait here until the completionSignal is released in the callback flow.
    completionSignal.acquire()

    println("\n🏁 Demo completed. Press enter to exit.")
    readln()
}

/**
 * Sets up the email service provider's Self account and defines the callback
 * for handling incoming connections from mobile devices.
 *
 * @param completionSignal A semaphore that will be released when the workflow is finished.
 * @return The configured Account instance.
 */
fun createEmailServiceAccount(completionSignal: Semaphore): Account {
    println("Setting up email service provider...")

    val callbacks = object : Common.Callbacks {
        override fun onWelcome(account: Account, welcome: Welcome) {
            // This function is triggered when a mobile app scans the QR code.
            handleMobileConnection(account, welcome, completionSignal)
        }
    }

    val emailService = Common.setupAccount(
        storagePath = "email_service_storage",
        callbacks = callbacks
    )

    println("✅ Email service provider ready")
    return emailService
}

/**
 * Generates a scannable QR code that initiates the connection from a Self mobile app.
 *
 * @param emailService The account that will receive the connection.
 * @return True if the QR code was generated successfully, false otherwise.
 */
fun generateEmailVerificationQR(emailService: Account): Boolean {
    println("Generating QR code for mobile email verification...")

    return runCatching {
        val inboxAddress = Common.openInbox(emailService)
        if (inboxAddress == null) throw IllegalStateException("Failed to open inbox")

        val expires = Timestamp.now() + 1800 // 30 minutes

        // Generate a key package for secure out-of-band communication.
        val keyPackage = emailService.connectionNegotiateOutOfBand(inboxAddress, expires)

        // Build a discovery request message containing the key package.
        val discoveryRequest = DiscoveryRequestBuilder()
            .keyPackage(keyPackage)
            .expires(expires)
            .finish()

        // Create an anonymous message and encode it into a QR code.
        val anonymousMessage = AnonymousMessage.fromContent(discoveryRequest)
        anonymousMessage.setFlags(FlagSet(Flag.TARGET_SANDBOX))

        val qrCodeBytes = anonymousMessage.encodeQR(QrEncoding.UNICODE)
        if (qrCodeBytes.isEmpty()) throw IllegalStateException("Failed to generate QR code bytes")

        // Display the QR code and its expiration time.
        println("\n${qrCodeBytes.decodeToString()}")
        val expirationTime = System.currentTimeMillis() + (30 * 60 * 1000)
        println("⏱️  Expires: ${java.text.SimpleDateFormat("HH:mm:ss").format(java.util.Date(expirationTime))}")

        true
    }.getOrElse { error ->
        println("❌ Failed to generate QR code: ${error.message}")
        false
    }
}

/**
 * Handles the incoming connection request from a mobile device.
 * This function is called by the `onWelcome` callback.
 *
 * @param emailService The account that received the connection.
 * @param welcome The connection event details.
 * @param completionSignal The semaphore to release upon completion of the entire flow.
 */
fun handleMobileConnection(emailService: Account, welcome: Welcome, completionSignal: Semaphore) {
    println("\n📱 Mobile device connected: ${welcome.fromAddress().encodeHex()}")

    val acceptSignal = Semaphore(0)

    // Accept the connection request to establish a secure channel.
    emailService.connectionAccept(asAddress = welcome.toAddress(), welcome = welcome.welcome()) { status: SelfStatus, _ ->
        if (status.success()) {
            println("✅ Mobile connection established!")
            // Now that we're connected, create and issue the credential.
            demonstrateEmailCredentialCreation(emailService, welcome.fromAddress(), completionSignal)
        } else {
            println("❌ Failed to accept connection: ${status.errorMessage()}")
            completionSignal.release() // Release signal on failure to unblock the main thread.
        }
        acceptSignal.release()
    }

    // Wait for the connection acceptance to complete before this function returns.
    acceptSignal.acquire()
}

/**
 * Orchestrates the creation and issuance of the email verification credential.
 *
 * @param emailService The issuer's account.
 * @param holderAddress The public key of the recipient (the mobile device).
 * @param completionSignal The semaphore to release to signal the end of the demo.
 */
fun demonstrateEmailCredentialCreation(emailService: Account, holderAddress: PublicKey, completionSignal: Semaphore) {
    println("\nCreating email verification credential...")

    val emailCredential = createEmailVerificationCredential(emailService, holderAddress)

    if (emailCredential != null) {
        println("✅ Email verification credential created successfully!")
        println("📱 Credential details:")
        println("   • Email: user@example.com")
        println("   • Status: verified")
        println("   • Issuer: Example Email Service Provider")
        println("\n🎉 Email verification workflow completed!")
        println("💡 In a real application, this credential is now stored in the user's mobile wallet.")
    } else {
        println("❌ Failed to create and issue email credential.")
    }

    // Signal that the entire workflow is complete.
    completionSignal.release()
}

/**
 * Builds and issues the verifiable credential for the verified email address.
 *
 * @param issuerAccount The account issuing the credential.
 * @param holderAddress The public key of the credential recipient.
 * @return The issued VerifiableCredential, or null on failure.
 */
fun createEmailVerificationCredential(issuerAccount: Account, holderAddress: PublicKey): VerifiableCredential? {
    return runCatching {
        val issuerAddress = issuerAccount.keychainSigningCreate()

        // Define the claims for the email credential.
        val claims = mapOf<String, Any>(
            "emailAddress" to "user@example.com",
            "verified" to true,
            "verificationDate" to LocalDateTime.now().format(DateTimeFormatter.ISO_DATE_TIME),
            "domain" to "example.com",
            "verificationMethod" to "email_link_clicked",
            "issuerName" to "Example Email Service Provider"
        )

        // Use the CredentialBuilder to construct the credential.
        val unsignedCredential = CredentialBuilder()
            .credentialType(arrayOf("VerifiableCredential", "EmailCredential"))
            .credentialSubject(Address.key(holderAddress))
            .credentialSubjectClaims(claims)
            .issuer(Address.key(issuerAddress))
            .validFrom(Timestamp.now())
            .signWith(issuerAddress, Timestamp.now())
            .finish()

        // Issue the credential, which signs it and prepares it for delivery.
        issuerAccount.credentialIssue(unsignedCredential)

    }.getOrElse { error ->
        println("❌ Failed to build or issue email credential: ${error.message}")
        null
    }
}