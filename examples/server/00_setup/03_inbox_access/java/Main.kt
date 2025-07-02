import com.joinself.selfsdk.kmp.account.Account
import com.joinself.selfsdk.kmp.account.LogLevel
import com.joinself.selfsdk.kmp.account.Target

fun main() {
    println("📧 Inbox Access Example")
    println("========================")

    // Load or create account
    println("🔧 Setting up account...")
    val account = setupAccount()
    println("✅ Account loaded successfully")

    // Access inbox to get the address
    println("\n📬 Accessing inbox...")
    val address = getInboxAddress(account)
    
    // Display the inbox address
    println("✅ Inbox opened successfully!")
    println("📬 Your inbox address: $address")
    println("\n💡 This address can be shared with others to receive:")
    println("   • Messages")
    println("   • Connection requests")
    println("   • Credentials")

    println("\n✅ Inbox access demonstration complete!")
}

fun setupAccount(): Account {
    val storageKey = "276cb6191a345753adb0897c2c0a89370aebf44ef99e612747bee3cd4e757ffa"
        .chunked(2).map { it.toInt(16).toByte() }.toByteArray()

    val account = Account()
    account.configure(
        storagePath = "./storage",
        storageKey = storageKey,
        rpcEndpoint = Target.PRODUCTION_SANDBOX.rpcEndpoint(),
        objectEndpoint = Target.PRODUCTION_SANDBOX.objectEndpoint(),
        messageEndpoint = Target.PRODUCTION_SANDBOX.messageEndpoint(),
        logLevel = LogLevel.WARN,
        onConnect = { /* Connection handled silently */ },
        onDisconnect = { _ -> },
        onAcknowledgement = { _ -> },
        onError = { _, _ -> },
        onCommit = { _ -> },
        onKeyPackage = { _ -> },
        onWelcome = { _ -> },
        onProposal = { _ -> },
        onMessage = { _ -> },
        onIntegrity = null
    )
    
    Thread.sleep(2000) // Simple wait for connection
    return account
}

fun getInboxAddress(account: Account): String {
    var address = ""
    account.inboxOpen { status, addr -> 
        if (status.success()) {
            address = addr.encodeHex()
        } else {
            println("❌ Failed to open inbox")
        }
    }
    Thread.sleep(1000) // Simple wait
    return address
} 
