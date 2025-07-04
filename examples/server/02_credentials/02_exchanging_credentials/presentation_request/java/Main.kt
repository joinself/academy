import com.joinself.selfsdk.kmp.account.Account
import com.joinself.selfsdk.kmp.credential.Address
import com.joinself.selfsdk.kmp.credential.CredentialBuilder
import com.joinself.selfsdk.kmp.credential.PresentationBuilder
import com.joinself.selfsdk.kmp.credential.VerifiableCredential
import com.joinself.selfsdk.kmp.credential.VerifiablePresentation
import com.joinself.selfsdk.kmp.keypair.signing.PublicKey
import com.joinself.selfsdk.kmp.time.Timestamp

/**
 * Represents a party in the credential exchange, holding an account and a
 * collection of their credentials.
 *
 * @param name A descriptive name for the party (e.g., "University Issuer").
 * @param account The underlying Self SDK account for this party.
 * @param credentials A map storing the verifiable credentials held by this party.
 */
data class ExchangeParty(
    val name: String,
    val account: Account,
    val credentials: MutableMap<String, VerifiableCredential> = mutableMapOf()
)

/**
 * Main entry point for the credential exchange demonstration.
 *
 * This example shows the foundational workflow for credential exchange:
 * 1. An **Issuer** (a university) issues multiple credentials to a **Holder** (a student).
 * 2. A **Verifier** (an employer) requests proof of education from the Holder.
 * 3. The **Holder** finds the required credentials in their wallet and creates a
 *    **Verifiable Presentation** to share with the Verifier.
 *
 * This demonstrates how issued credentials become the building blocks for secure data sharing.
 */
fun main() {
    println("Credential Exchange Demo")
    println("========================")

    // 1. Create the parties for the exchange scenario (an issuer and a holder)
    val (issuer, holder) = createExchangeParties()

    // 2. The issuer creates and sends several credentials to the holder
    issueCredentialsForExchange(issuer, holder)

    // 3. A verifier requests specific credentials, and the holder prepares a presentation
    demonstrateExchangeWorkflow(holder)

    println("\n✅ Exchange demo completed!")
    println("Press enter to exit")
    readln()
}

/**
 * Sets up the issuer and holder accounts and wraps them in the ExchangeParty data class.
 */
fun createExchangeParties(): Pair<ExchangeParty, ExchangeParty> {
    println("\nSetting up exchange parties...")

    // Create issuer party
    val issuerAccount = Common.setupAccount(storagePath = "./exchange_issuer_storage")
    val issuer = ExchangeParty("University Issuer", issuerAccount)

    // Create holder party
    val holderAccount = Common.setupAccount(storagePath = "./exchange_holder_storage")
    val holder = ExchangeParty("Student Holder", holderAccount)

    println()
    Common.displayAccountInfo(issuer.account, issuer.name)
    Common.displayAccountInfo(holder.account, holder.name)

    return Pair(issuer, holder)
}

/**
 * Issues a set of credentials from the issuer to the holder, populating the holder's "wallet".
 */
fun issueCredentialsForExchange(issuer: ExchangeParty, holder: ExchangeParty) {
    println("\nIssuing credentials for the holder...")

    val issuerAddress = Common.openInbox(issuer.account) ?: return
    val holderAddress = Common.openInbox(holder.account) ?: return

    // Issue and store an email credential
    createEmailCredential(issuer.account, issuerAddress, holderAddress)?.let {
        holder.credentials["email"] = it
        println("✅ Email credential issued")
    }

    // Issue and store a student ID credential
    createStudentCredential(issuer.account, issuerAddress, holderAddress)?.let {
        holder.credentials["student_id"] = it
        println("✅ Student ID credential issued")
    }

    // Issue and store a degree credential
    createDegreeCredential(issuer.account, issuerAddress, holderAddress)?.let {
        holder.credentials["degree"] = it
        println("✅ Degree credential issued")
    }

    println("🎓 Holder now has ${holder.credentials.size} credentials.")
}

/**
 * Simulates a verifier requesting credentials and the holder creating a presentation in response.
 */
fun demonstrateExchangeWorkflow(holder: ExchangeParty) {
    println("\nDemonstrating exchange workflow...")
    println("Scenario: An employer requests proof of education from the student holder.")

    // 1. A verifier requests credentials of specific types
    val requestedTypes = listOf("StudentCredential", "DegreeCredential")
    println("🔍 Verifier is requesting credential types: $requestedTypes")

    // 2. The holder searches their local credential store for matches
    val matchingCreds = holder.credentials.values.filter { cred ->
        val credTypes = cred.credentialType()
        // Check if any of the credential's types are in the requested list
        credTypes.any { it in requestedTypes }
    }

    if (matchingCreds.isEmpty()) {
        println("❌ No matching credentials found in holder's wallet.")
        return
    }

    println("✅ Found ${matchingCreds.size} matching credentials:")
    matchingCreds.forEach { cred ->
        // Find which of the credential's types was the one that matched
        val matchedType = cred.credentialType().first { it in requestedTypes }
        val identifier = cred.credentialSubjectClaims()["studentId"] ?: cred.credentialSubjectClaims()["degree"]
        println("   - '$identifier' ($matchedType)")
    }

    // 3. The holder creates a Verifiable Presentation containing the matching credentials
    val presentation = createCredentialPresentation(holder, matchingCreds)
    if (presentation != null) {
        println("✅ Presentation created with ${presentation.credentials().size} credentials.")
        println("   - Presentation Type: ${presentation.presentationType().toList()}")
        println("   - Holder: ${presentation.holder().address().encodeHex()}")
        println("✅ Exchange workflow completed successfully!")
    }
}

/**
 * Creates a Verifiable Presentation from a list of credentials, signed by the holder.
 * This is the standard way to package and share credentials with a verifier.
 */
fun createCredentialPresentation(holder: ExchangeParty, credentials: List<VerifiableCredential>): VerifiablePresentation? {
    if (credentials.isEmpty()) return null

    println("✍️  Creating Verifiable Presentation...")
    val holderAddress = Common.openInbox(holder.account) ?: return null

    return runCatching {
        val builder = PresentationBuilder()
            .presentationType(arrayOf("VerifiablePresentation", "ProofOfEducation"))
            .holder(Address.key(holderAddress))

        // Add all matching credentials to the presentation
        credentials.forEach { builder.credentialAdd(it) }

        val unsignedPresentation = builder.finish()

        // The holder's account signs and issues the presentation
        holder.account.presentationIssue(unsignedPresentation)
    }.getOrElse { error ->
        println("❌ Failed to create presentation: ${error.message}")
        null
    }
}

// --- Credential Creation Helper Functions ---

fun createEmailCredential(issuerAccount: Account, issuerAddress: PublicKey, holderAddress: PublicKey): VerifiableCredential? {
    return runCatching {
        val claims = mapOf(
            "emailAddress" to "student@university.edu",
            "verified" to true,
            "domain" to "university.edu"
        )

        val unsignedCredential = CredentialBuilder()
            .credentialType(arrayOf("VerifiableCredential", "EmailCredential"))
            .credentialSubject(Address.key(holderAddress))
            .credentialSubjectClaims(claims)
            .issuer(Address.key(issuerAddress))
            .validFrom(Timestamp.now())
            .signWith(issuerAddress, Timestamp.now())
            .finish()

        issuerAccount.credentialIssue(unsignedCredential)
    }.getOrElse {
        println("❌ Failed to issue email credential: ${it.message}")
        null
    }
}

fun createStudentCredential(issuerAccount: Account, issuerAddress: PublicKey, holderAddress: PublicKey): VerifiableCredential? {
    return runCatching {
        val claims = mapOf(
            "studentId" to "STU-2024-001",
            "enrollmentDate" to "2020-09-01",
            "status" to "enrolled",
            "program" to "Computer Science",
            "level" to "undergraduate"
        )

        val unsignedCredential = CredentialBuilder()
            .credentialType(arrayOf("VerifiableCredential", "StudentCredential"))
            .credentialSubject(Address.key(holderAddress))
            .credentialSubjectClaims(claims)
            .issuer(Address.key(issuerAddress))
            .validFrom(Timestamp.now())
            .signWith(issuerAddress, Timestamp.now())
            .finish()

        issuerAccount.credentialIssue(unsignedCredential)
    }.getOrElse {
        println("❌ Failed to issue student credential: ${it.message}")
        null
    }
}

fun createDegreeCredential(issuerAccount: Account, issuerAddress: PublicKey, holderAddress: PublicKey): VerifiableCredential? {
    return runCatching {
        val claims = mapOf(
            "degree" to "Bachelor of Science",
            "major" to "Computer Science",
            "graduationDate" to "2024-05-15",
            "gpa" to 3.8,
            "honors" to "magna cum laude",
            "institution" to "University of Technology"
        )

        val unsignedCredential = CredentialBuilder()
            .credentialType(arrayOf("VerifiableCredential", "DegreeCredential"))
            .credentialSubject(Address.key(holderAddress))
            .credentialSubjectClaims(claims)
            .issuer(Address.key(issuerAddress))
            .validFrom(Timestamp.now())
            .signWith(issuerAddress, Timestamp.now())
            .finish()

        issuerAccount.credentialIssue(unsignedCredential)
    }.getOrElse {
        println("❌ Failed to issue degree credential: ${it.message}")
        null
    }
}
