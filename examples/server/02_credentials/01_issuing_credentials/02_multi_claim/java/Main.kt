import com.joinself.selfsdk.account.Account
import com.joinself.selfsdk.credential.Address
import com.joinself.selfsdk.credential.CredentialBuilder
import com.joinself.selfsdk.time.Timestamp
import java.time.LocalDate
import java.time.format.DateTimeFormatter
import java.util.concurrent.Semaphore

fun main() {
    val signal = Semaphore(0)
    println("Multi-Claim Credential Issuance Demo")
    println("====================================")

    // Create issuer and holder accounts
    println("Setting up Self account...")
    val issuer = Common.setupAccount(storagePath = "./multi_issuer_storage")
    val holder = Common.setupAccount(storagePath = "./multi_holder_storage")
    println("✅ Account ready!")

    println()
    Common.displayAccountInfo(issuer, "Issuer")
    Common.displayAccountInfo(holder, "Holder")

    createProfileCredential(issuer, holder)
    createEducationCredential(issuer, holder)
    println("\n✅ Demo completed!")

    println("\nPress enter to exit")
    readln()
}

fun createProfileCredential(issuer: Account, holder: Account) {
    println("\nCreating profile credential...")

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
        "firstName" to "John",
        "lastName" to "Doe",
        "displayName" to "John Doe",
        "profileLevel" to "verified",
        "country" to "United States",
        "age" to 30,
        "isActive" to true,
        "registrationDate" to LocalDate.now().format(DateTimeFormatter.ISO_LOCAL_DATE)
    )

    val unsignedCredential = CredentialBuilder()
        .credentialType(arrayOf("ProfileNameCredential"))
        .credentialSubject(Address.key(holderAddress))
        .credentialSubjectClaims(claims)
        .issuer(Address.key(issuerAddress))
        .validFrom(Timestamp.now())
        .signWith(issuerAddress, Timestamp.now())
        .finish()

    val profileCredential = runCatching {
        issuer.credentialIssue(unsignedCredential)
    }.getOrElse { error ->
        println("❌ Failed to issue credential: ${error.message}")
        return
    }

    println("✅ Profile credential issued")
    println("   🔒 Type: ${profileCredential.credentialType().toList()}")
    println("   🆔 Subject: ${profileCredential.credentialSubject()}")
}

fun createEducationCredential(issuer: Account, holder: Account) {
    println("\nCreating education credential...")

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
        "institution" to "University of Technology",
        "degree" to "Bachelor of Science",
        "major" to "Computer Science",
        "graduationYear" to 2020,
        "gpa" to 3.8,
        "honors" to true,
        "creditsCompleted" to 120,
        "thesis" to "Machine Learning Applications",
        "graduationDate" to "2020-05-15",
    )

    val unsignedCredential = CredentialBuilder()
        .credentialType(arrayOf("VerifiableCredential", "EducationCredential"))
        .credentialSubject(Address.key(holderAddress))
        .credentialSubjectClaims(claims)
        .issuer(Address.key(issuerAddress))
        .validFrom(Timestamp.now())
        .signWith(issuerAddress, Timestamp.now())
        .finish()

    val educationCredential = runCatching {
        issuer.credentialIssue(unsignedCredential)
    }.getOrElse { error ->
        println("❌ Failed to issue credential: ${error.message}")
        return
    }

    println("✅ Education credential issued")
    println("   🔒 Type: ${educationCredential.credentialType().toList()}")
    println("   🆔 Subject: ${educationCredential.credentialSubject()}")
}