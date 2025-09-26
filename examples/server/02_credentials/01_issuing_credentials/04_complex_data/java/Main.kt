/**
 * Demonstrates complex credential issuance using the Self SDK.
 *
 * This is the COMPLEX level of credential issuance examples.
 * Prerequisites: Complete the basic, multi-claim, and evidence examples first.
 *
 * This example shows:
 * - Complex nested objects in claims
 * - Arrays and collections in credentials
 * - Hierarchical data organization
 * - Real-world organizational data modeling
 * - Advanced claim structuring techniques
 *
 * 🎯 What you'll learn:
 * • How to structure complex nested data in credentials
 * • Using arrays and collections in claims
 * • Hierarchical data organization
 * • Real-world data modeling patterns
 * • Advanced claim structuring
 *
 * 📚 Next steps:
 * • Explore the advanced example, which combines all features.
 */

import com.joinself.selfsdk.account.Account
import com.joinself.selfsdk.credential.Address
import com.joinself.selfsdk.credential.CredentialBuilder
import com.joinself.selfsdk.time.Timestamp

fun main() {
    println("Complex Credential Issuance Demo")
    println("================================")

    // Create issuer and holder accounts
    println("\nSetting up Self accounts...")
    val issuer = Common.setupAccount(storagePath = "./complex_issuer_storage")
    val holder = Common.setupAccount(storagePath = "./complex_holder_storage")
    println("✅ Accounts ready!")

    println()
    Common.displayAccountInfo(issuer, "Issuer")
    Common.displayAccountInfo(holder, "Holder")

    // Create a credential with a complex, nested data structure
    createOrganizationCredential(issuer, holder)

    println("\n✅ Demo completed successfully!")
    println("\nPress enter to exit")
    readln()
}

/**
 * Creates an organization credential with complex nested data, including maps and lists.
 * @param issuer The account issuing the credential.
 * @param holder The account that will receive the credential.
 */
fun createOrganizationCredential(issuer: Account, holder: Account) {
    println("\nCreating organization credential with complex data...")

    // Get inbox addresses for the issuer and holder
    val issuerAddress = Common.openInbox(issuer) ?: run {
        println("❌ Failed to open issuer inbox")
        return
    }
    val holderAddress = Common.openInbox(holder) ?: run {
        println("❌ Failed to open holder inbox")
        return
    }

    // Create a complex claims structure with nested objects and arrays.
    // This demonstrates how to model real-world hierarchical data.
    val claims = mapOf<String, Any>(
        "organizationName" to "TechCorp Inc.",
        "employeeId" to "EMP-2024-001",
        "position" to mapOf(
            "title" to "Senior Software Engineer",
            "department" to "Engineering",
            "level" to "L5",
            "startDate" to "2024-01-15",
            "manager" to "jane.smith@techcorp.com",
            "team" to "Backend Infrastructure"
        ),
        "permissions" to listOf(
            "read:repositories",
            "write:code",
            "deploy:staging",
            "review:pull-requests",
            "admin:team-resources"
        ),
        "contact" to mapOf(
            "email" to "john.doe@techcorp.com",
            "phone" to "+1-555-0123",
            "office" to "Building A, Floor 3, Desk 42",
            "timezone" to "America/New_York",
            "address" to mapOf(
                "street" to "123 Tech Street",
                "city" to "San Francisco",
                "state" to "CA",
                "zipCode" to "94105",
                "country" to "United States"
            )
        ),
        "benefits" to mapOf(
            "healthInsurance" to true,
            "retirement401k" to true,
            "paidTimeOff" to 25,
            "stockOptions" to 1000,
            "remoteWork" to true,
            "wellness" to mapOf(
                "gymMembership" to true,
                "mentalHealth" to true,
                "annualWellness" to "$1000",
                "flexibleSchedule" to true
            )
        ),
        "certifications" to listOf(
            mapOf(
                "name" to "AWS Solutions Architect",
                "level" to "Professional",
                "issueDate" to "2023-06-15",
                "expiryDate" to "2026-06-15",
                "verified" to true,
                "provider" to "Amazon Web Services"
            ),
            mapOf(
                "name" to "Kubernetes Administrator",
                "level" to "Certified",
                "issueDate" to "2023-09-20",
                "expiryDate" to "2026-09-20",
                "verified" to true,
                "provider" to "Cloud Native Computing Foundation"
            )
        ),
        "projects" to listOf(
            mapOf(
                "name" to "Payment Gateway Redesign",
                "role" to "Lead Developer",
                "startDate" to "2023-01-01",
                "endDate" to "2023-06-30",
                "status" to "Completed",
                "technologies" to listOf("Go", "PostgreSQL", "Redis", "Docker")
            ),
            mapOf(
                "name" to "Microservices Migration",
                "role" to "Senior Engineer",
                "startDate" to "2023-07-01",
                "endDate" to "2024-01-31",
                "status" to "Completed",
                "technologies" to listOf("Go", "Kubernetes", "gRPC", "Prometheus")
            )
        )
    )

    // Build and issue the credential using the Self SDK
    runCatching {
        val unsignedCredential = CredentialBuilder()
            .credentialType("OrganizationCredential")
            .credentialSubject(Address.key(holderAddress))
            .credentialSubjectClaims(claims)
            .issuer(Address.key(issuerAddress))
            .validFrom(Timestamp.now())
            .signWith(issuerAddress, Timestamp.now())
            .finish()

        val orgCredential = issuer.credentialIssue(unsignedCredential)

        println("✅ Organization credential issued (Type: ${orgCredential.credentialType().toList()})")
        println("   - Employee: EMP-2024-001 - Senior Software Engineer")
        println("   - Data structure: 5 permissions, 2 certifications, 2 projects")

    }.getOrElse { error ->
        println("❌ Failed to issue credential: ${error.message}")
    }
}