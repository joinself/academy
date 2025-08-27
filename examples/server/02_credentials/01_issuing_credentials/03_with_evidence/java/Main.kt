/**
 * Demonstrates credential issuance with evidence using the Self SDK.
 *
 * This is the EVIDENCE level of credential issuance examples.
 * Prerequisites: Complete the basic and multi-claim examples first.
 *
 * This example shows:
 * - Creating custom credential types
 * - Attaching file evidence to credentials
 * - Asset management and upload functionality
 * - Creating verifiable presentations
 * - Linking evidence to credential claims
 *
 * 🎯 What you'll learn:
 * • How to attach evidence files to credentials
 * • Asset management and secure storage
 * • Creating verifiable presentations
 * • Linking evidence to claims with hashes
 * • Custom credential types
 *
 * 📚 Next steps:
 * • Explore complex nested data structures in credentials.
 * • See the advanced example for all features combined.
 */

import com.joinself.selfsdk.account.Account
import com.joinself.selfsdk.asset.BinaryObject
import com.joinself.selfsdk.credential.Address
import com.joinself.selfsdk.credential.CredentialBuilder
import com.joinself.selfsdk.credential.PresentationBuilder
import com.joinself.selfsdk.credential.VerifiableCredential
import com.joinself.selfsdk.time.Timestamp
import java.time.LocalDate
import java.time.format.DateTimeFormatter
import java.util.concurrent.Semaphore

fun main() {
    println("Evidence-Based Credential Issuance Demo")
    println("=======================================")

    // Create issuer and holder accounts
    println("\nSetting up Self accounts...")
    val issuer = Common.setupAccount(storagePath = "./evidence_issuer_storage")
    val holder = Common.setupAccount(storagePath = "./evidence_holder_storage")
    println("✅ Accounts ready!")

    println()
    Common.displayAccountInfo(issuer, "Issuer")
    Common.displayAccountInfo(holder, "Holder")

    // Create credentials with evidence
    createCertificationWithEvidence(issuer, holder)

    println("\n✅ Demo completed successfully!")
    println("\nPress enter to exit")
    readln()
}

/**
 * Orchestrates the creation of evidence, a credential, and a presentation.
 */
fun createCertificationWithEvidence(issuer: Account, holder: Account) {
    println("\nCreating certification with evidence...")

    // 1. Create and upload the evidence asset
    val evidence = createEvidence(issuer)
    if (evidence == null) {
        println("❌ Failed to create evidence, aborting.")
        return
    }

    // 2. Create a credential that references the evidence
    val cred = createCredentialWithEvidence(issuer, holder, evidence)
    if (cred == null) {
        println("❌ Failed to create credential, aborting.")
        return
    }

    // 3. Create a verifiable presentation containing the credential
    createPresentation(issuer, cred)
}

/**
 * Creates a file-based evidence object and uploads it to the issuer's storage.
 * @param issuer The account that will own and upload the evidence.
 * @return The created SelfObject on success, or null on failure.
 */
@OptIn(ExperimentalStdlibApi::class)
fun createEvidence(issuer: Account): BinaryObject? {
    println("Creating evidence asset...")

    val certificateData = """
    Certificate of Completion
    Advanced Kotlin Programming Course

    Student: John Doe
    Course: Advanced Kotlin Programming with Self SDK
    Institution: Self SDK Academy
    Grade: A+
    Credits: 40 hours
    Date: ${LocalDate.now().format(DateTimeFormatter.ISO_LOCAL_DATE)}

    This certificate verifies that the above-named student has
    successfully completed the Advanced Kotlin Programming course
    with distinction.

    Instructor: Jane Smith
    Director: Dr. Alice Johnson
    """.trimIndent()

    val evidence = runCatching {
        val signal = Semaphore(0)
        val obj = BinaryObject.create("plain/text", certificateData.toByteArray())
        issuer.objectUpload(obj, true) { status ->
            signal.release()
        }
        signal.acquire()
        obj
    }.getOrElse { error ->
        println("❌ Failed to create or upload evidence: ${error.message}")
        return null
    }

    println("✅ Evidence uploaded: ${evidence.id()?.toHexString()} (${certificateData.toByteArray().size} bytes)")
    return evidence
}

/**
 * Creates a verifiable credential that includes a reference to the evidence object.
 * @param issuer The account issuing the credential.
 * @param holder The account receiving the credential.
 * @param evidence The evidence object to link in the credential claims.
 * @return The issued VerifiableCredential on success, or null on failure.
 */
@OptIn(ExperimentalStdlibApi::class)
fun createCredentialWithEvidence(issuer: Account, holder: Account, evidence: BinaryObject): VerifiableCredential? {
    println("Building certification credential...")

    val issuerAddress = Common.openInbox(issuer) ?: return null
    val holderAddress = Common.openInbox(holder) ?: return null

    val claims = mapOf(
        "certificationType" to "Professional Development",
        "courseName" to "Advanced Kotlin Programming",
        "completionDate" to LocalDate.now().format(DateTimeFormatter.ISO_LOCAL_DATE),
        "grade" to "A+",
        "institution" to "Self SDK Academy",
        "courseHours" to 40,
        "instructor" to "Jane Smith",
        "evidenceId" to (evidence.id()?.toHexString() ?: "")
    )

    return runCatching {
        val unsignedCredential = CredentialBuilder()
            .credentialType(arrayOf("VerifiableCredential", "CertificationCredential"))
            .credentialSubject(Address.key(holderAddress))
            .credentialSubjectClaims(claims)
            .issuer(Address.key(issuerAddress))
            .validFrom(Timestamp.now())
            .signWith(issuerAddress, Timestamp.now())
            .finish()

        val customCredential = issuer.credentialIssue(unsignedCredential)
        println("✅ Certification credential issued (Type: ${customCredential.credentialType().toList()})")
        customCredential
    }.getOrElse { error ->
        println("❌ Failed to issue credential: ${error.message}")
        null
    }
}

/**
 * Creates a verifiable presentation containing the issued credential.
 * In this example, the issuer creates a presentation for itself.
 * @param issuer The account that will hold the presentation.
 * @param cred The credential to include in the presentation.
 */
fun createPresentation(issuer: Account, cred: VerifiableCredential) {
    println("Creating verifiable presentation...")

    val issuerAddress = Common.openInbox(issuer) ?: return

    runCatching {
        val unsignedPresentation = PresentationBuilder()
            .presentationType(arrayOf("VerifiablePresentation", "CertificationPresentation"))
            .holder(Address.key(issuerAddress))
            .credentialAdd(cred)
            .finish()

        val presentation = issuer.presentationIssue(unsignedPresentation)
        println("✅ Presentation created (Type: ${presentation.presentationType().toList()}, Credentials: ${presentation.credentials().size})")
    }.getOrElse { error ->
        println("❌ Failed to issue presentation: ${error.message}")
    }
}
