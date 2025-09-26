import com.joinself.selfsdk.account.Account
import com.joinself.selfsdk.credential.Address
import com.joinself.selfsdk.credential.CredentialBuilder
import com.joinself.selfsdk.credential.CredentialType
import com.joinself.selfsdk.time.Timestamp
import java.util.concurrent.Semaphore

fun main() {
    val signal = Semaphore(0)
    println("🎓 Basic Credential Issuance Demo")
    println("==========================================")

    // Setup: Create a Self account with connection handling
    println("Setting up Self account...")
    val issuer = Common.setupAccount(storagePath = "./basic_issuer")
    val holder = Common.setupAccount(storagePath = "./basic_holder")

    println("✅ Account ready!")
    Common.displayAccountInfo(issuer, "Issuer")
    Common.displayAccountInfo(holder, "Holder")

    createEmailCredential(issuer, holder)
    println("\n✅ Demo completed!")

    println("\nPress enter to exit")
    readln()
}

fun createEmailCredential(issuer: Account, holder: Account) {
    println("📧 Creating email credential...")

    val issuerAddress = Common.openInbox(issuer)
    if (issuerAddress == null) {
        println("Failed to open issuer inbox")
        return
    }
    val holderAddress = Common.openInbox(holder)
    if (holderAddress == null) {
        println("Failed to open holder inbox")
        return
    }
    val claims = mapOf(
        "emailAddress" to "alice@example.com",
        "verified" to true,
        "domain" to "example.com"
    )
    val unsignedCredential = CredentialBuilder()
        .credentialType(CredentialType.EMAIL)
        .credentialSubject(Address.key(holderAddress))
        .credentialSubjectClaims(claims)
        .issuer(Address.key(issuerAddress))
        .validFrom(Timestamp.now())
        .signWith(issuerAddress, Timestamp.now())
        .finish()

    val emailCredential = runCatching {
        issuer.credentialIssue(unsignedCredential)
    }.getOrElse { error ->
        println("❌ Failed to issue credential: ${error.message}")
        return
    }

    println("✅ Email credential created")
    println("   📧 Email: alice@example.com")
    println("   🔒 Type: ${emailCredential.credentialType().toList()}")
    println("   🆔 Subject: ${emailCredential.credentialSubject()}")
}