/**
 * Demonstrates comprehensive credential issuance capabilities using the Self SDK.
 *
 * This example combines all previous concepts into a single demonstration:
 * - Basic credential issuance (e.g., email verification).
 * - Multi-claim credentials for richer data (e.g., user profiles).
 * - Credentials with linked file evidence (e.g., certifications).
 * - Credentials with complex, nested data structures (e.g., organizational roles).
 * - Creation of Verifiable Presentations to share credentials.
 *
 * 🎯 What you'll learn:
 * • A complete overview of the credential issuance lifecycle.
 * • Issuing various types of credentials, from simple to complex.
 * • Attaching file evidence to credentials using asset management.
 * • Structuring complex nested data and arrays within claims.
 * • Creating Verifiable Presentations to package and share credentials.
 */
import com.joinself.selfsdk.kmp.account.Account
import com.joinself.selfsdk.kmp.asset.BinaryObject
import com.joinself.selfsdk.kmp.credential.Address
import com.joinself.selfsdk.kmp.credential.CredentialBuilder
import com.joinself.selfsdk.kmp.credential.PresentationBuilder
import com.joinself.selfsdk.kmp.credential.VerifiableCredential
import com.joinself.selfsdk.kmp.credential.VerifiablePresentation
import com.joinself.selfsdk.kmp.time.Timestamp
import java.time.LocalDate
import java.time.format.DateTimeFormatter
import java.util.concurrent.Semaphore

fun main() {
    println("Comprehensive Credential Issuance Demo")
    println("======================================")

    // Initialize issuer and holder accounts
    println("\nSetting up Self accounts...")
    val issuer = Common.setupAccount(storagePath = "./comprehensive_issuer_storage")
    val holder = Common.setupAccount(storagePath = "./comprehensive_holder_storage")
    println("✅ Accounts ready!")

    println()
    Common.displayAccountInfo(issuer, "Issuer")
    Common.displayAccountInfo(holder, "Holder")

    println("\nCREDENTIAL ISSUANCE EXAMPLES")
    println("============================")

    // Demonstrate each type of credential issuance
    demonstrateBasicCredential(issuer, holder)
    demonstrateProfileCredential(issuer, holder)
    demonstrateCustomCredentialWithEvidence(issuer, holder)
    demonstrateOrganizationCredential(issuer, holder)

    println("\n✅ All examples completed successfully!")
    println("\nPress enter to exit")
    readln()
}

/**
 * Demonstrates issuing a basic credential with a few simple claims.
 */
@OptIn(ExperimentalStdlibApi::class)
fun demonstrateBasicCredential(issuer: Account, holder: Account) {
    println("\n1. Basic Email Credential")
    println("=========================")

    val issuerAddress = Common.openInbox(issuer) ?: return
    val holderAddress = Common.openInbox(holder) ?: return

    val claims = mapOf(
        "emailAddress" to "john.doe@example.com",
        "verified" to true,
        "verificationDate" to LocalDate.now().format(DateTimeFormatter.ISO_LOCAL_DATE)
    )

    runCatching {
        val unsignedCredential = CredentialBuilder()
            .credentialType(arrayOf("EmailCredential"))
            .credentialSubject(Address.key(holderAddress))
            .credentialSubjectClaims(claims)
            .issuer(Address.key(issuerAddress))
            .validFrom(Timestamp.now())
            .signWith(issuerAddress, Timestamp.now())
            .finish()

        val emailCredential = issuer.credentialIssue(unsignedCredential)
        println("✅ Email credential issued (Type: ${emailCredential.credentialType().toList()})")
        println("   Email: john.doe@example.com (verified)")
    }.getOrElse { error ->
        println("❌ Failed to issue basic credential: ${error.message}")
    }
}

/**
 * Demonstrates issuing a credential with multiple top-level claims.
 */
@OptIn(ExperimentalStdlibApi::class)
fun demonstrateProfileCredential(issuer: Account, holder: Account) {
    println("\n2. Profile Credential")
    println("=====================")

    val issuerAddress = Common.openInbox(issuer) ?: return
    val holderAddress = Common.openInbox(holder) ?: return

    val claims = mapOf(
        "firstName" to "John",
        "lastName" to "Doe",
        "displayName" to "John Doe",
        "profileLevel" to "verified",
        "country" to "United States",
        "registrationDate" to LocalDate.now().format(DateTimeFormatter.ISO_LOCAL_DATE)
    )

    runCatching {
        val unsignedCredential = CredentialBuilder()
            .credentialType(arrayOf("ProfileNameCredential"))
            .credentialSubject(Address.key(holderAddress))
            .credentialSubjectClaims(claims)
            .issuer(Address.key(issuerAddress))
            .validFrom(Timestamp.now())
            .signWith(issuerAddress, Timestamp.now())
            .finish()

        val profileCredential = issuer.credentialIssue(unsignedCredential)
        println("✅ Profile credential issued (Type: ${profileCredential.credentialType().toList()})")
        println("   Name: John Doe, Country: United States")
    }.getOrElse { error ->
        println("❌ Failed to issue profile credential: ${error.message}")
    }
}

/**
 * Demonstrates issuing a credential with linked file evidence and creating a presentation.
 */
@OptIn(ExperimentalStdlibApi::class)
fun demonstrateCustomCredentialWithEvidence(issuer: Account, holder: Account) {
    println("\n3. Custom Credential with Evidence")
    println("==================================")

    val issuerAddress = Common.openInbox(issuer) ?: return
    val holderAddress = Common.openInbox(holder) ?: return

    // 1. Create and upload evidence asset
    val certificateData = """
    Certificate of Completion
    Advanced Kotlin Programming Course
    Student: John Doe
    Grade: A+
    Date: ${LocalDate.now().format(DateTimeFormatter.ISO_LOCAL_DATE)}
    """.trimIndent().toByteArray()

    val evidenceObj = runCatching {
        val signal = Semaphore(0)
        val obj = BinaryObject.create("application/pdf", certificateData)
        issuer.objectUpload(obj, true) { status ->
            signal.release()
        }
        signal.acquire()
        obj
    }.getOrElse { error ->
        println("❌ Failed to create or upload evidence: ${error.message}")
        return
    }
    println("✅ Evidence uploaded: ${evidenceObj.id()?.toHexString()} (${certificateData.size} bytes)")

    // 2. Create credential that links to the evidence
    val claims = mapOf(
        "certificationType" to "Professional Development",
        "courseName" to "Advanced Kotlin Programming",
        "completionDate" to LocalDate.now().format(DateTimeFormatter.ISO_LOCAL_DATE),
        "evidenceId" to (evidenceObj.id()?.toHexString() ?: ""),
        "grade" to "A+",
        "institution" to "Self SDK Academy",
        "courseHours" to 40
    )

    val customCredential = runCatching {
        val unsignedCredential = CredentialBuilder()
            .credentialType(arrayOf("VerifiableCredential", "CertificationCredential"))
            .credentialSubject(Address.key(holderAddress))
            .credentialSubjectClaims(claims)
            .issuer(Address.key(issuerAddress))
            .validFrom(Timestamp.now())
            .signWith(issuerAddress, Timestamp.now())
            .finish()
        issuer.credentialIssue(unsignedCredential)
    }.getOrElse { error ->
        println("❌ Failed to issue certification credential: ${error.message}")
        return
    }
    println("✅ Certification credential issued (Type: ${customCredential.credentialType().toList()})")

    // 3. Create a verifiable presentation containing the new credential
    val presentation = createPresentation(issuer, customCredential)
    if (presentation != null) {
        println("✅ Presentation created (Type: ${presentation.presentationType().toList()})")
        println("   Course: Advanced Kotlin Programming, Grade: A+")
    }
}

/**
 * Demonstrates issuing a credential with a complex, hierarchical data structure.
 */
@OptIn(ExperimentalStdlibApi::class)
fun demonstrateOrganizationCredential(issuer: Account, holder: Account) {
    println("\n4. Organization Credential (Complex Data)")
    println("=========================================")

    val issuerAddress = Common.openInbox(issuer) ?: return
    val holderAddress = Common.openInbox(holder) ?: return

    val claims = mapOf<String, Any>(
        "organizationName" to "Acme Corporation",
        "employeeID" to "EMP-12345",
        "position" to "Senior Software Engineer",
        "department" to "Engineering",
        "startDate" to "2020-01-15",
        "address" to mapOf(
            "street" to "123 Main Street",
            "city" to "San Francisco",
            "state" to "CA",
            "zipCode" to "94105",
            "country" to "United States"
        ),
        "skills" to listOf(
            "Kotlin Programming",
            "Microservices Architecture",
            "Docker & Kubernetes",
            "Cloud Computing",
            "API Design"
        ),
        "certifications" to listOf(
            mapOf(
                "name" to "AWS Solutions Architect",
                "issuer" to "Amazon Web Services",
                "issueDate" to "2021-03-15",
                "expiryDate" to "2024-03-15",
                "level" to "Professional"
            ),
            mapOf(
                "name" to "Kubernetes Administrator",
                "issuer" to "Cloud Native Computing Foundation",
                "issueDate" to "2021-08-20",
                "expiryDate" to "2024-08-20",
                "level" to "Certified"
            )
        ),
        "projects" to listOf(
            mapOf(
                "name" to "Self SDK Integration",
                "role" to "Lead Developer",
                "duration" to "6 months",
                "technologies" to listOf("Kotlin", "Self SDK", "REST APIs"),
                "status" to "Completed"
            ),
            mapOf(
                "name" to "Microservices Migration",
                "role" to "Senior Engineer",
                "duration" to "12 months",
                "technologies" to listOf("Docker", "Kubernetes", "gRPC"),
                "status" to "In Progress"
            )
        ),
        "performanceRating" to mapOf(
            "overall" to "Exceeds Expectations",
            "technical" to "Expert",
            "leadership" to "Strong",
            "reviewDate" to LocalDate.now().format(DateTimeFormatter.ISO_LOCAL_DATE),
            "reviewer" to "Jane Smith, Engineering Manager"
        )
    )

    runCatching {
        val unsignedCredential = CredentialBuilder()
            .credentialType(arrayOf("VerifiableCredential", "EmploymentCredential"))
            .credentialSubject(Address.key(holderAddress))
            .credentialSubjectClaims(claims)
            .issuer(Address.key(issuerAddress))
            .validFrom(Timestamp.now())
            .signWith(issuerAddress, Timestamp.now())
            .finish()

        val orgCredential = issuer.credentialIssue(unsignedCredential)
        println("✅ Organization credential issued (Type: ${orgCredential.credentialType().toList()})")
        println("   Employee: EMP-12345 - Senior Software Engineer")
        println("   Organization: Acme Corporation, Engineering Dept.")
    }.getOrElse { error ->
        println("❌ Failed to issue organization credential: ${error.message}")
    }
}

/**
 * Helper function to create a Verifiable Presentation from a credential.
 * @param account The account that will hold and sign the presentation.
 * @param cred The credential to include in the presentation.
 * @return The created VerifiablePresentation, or null on failure.
 */
fun createPresentation(account: Account, cred: VerifiableCredential): VerifiablePresentation? {
    println("   Creating verifiable presentation...")
    val holderAddress = Common.openInbox(account) ?: return null

    return runCatching {
        val unsignedPresentation = PresentationBuilder()
            .presentationType(arrayOf("VerifiablePresentation"))
            .holder(Address.key(holderAddress))
            .credentialAdd(cred)
            .finish()
        account.presentationIssue(unsignedPresentation)
    }.getOrElse { error ->
        println("   ❌ Failed to create presentation: ${error.message}")
        null
    }
}