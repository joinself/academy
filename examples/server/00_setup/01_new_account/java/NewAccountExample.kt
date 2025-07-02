import com.joinself.selfsdk.kmp.account.Account
import com.joinself.selfsdk.kmp.account.LogLevel
import com.joinself.selfsdk.kmp.account.Target
import java.io.File

fun main() {
    println("🆕 New Account Creation Example")
    println("===============================")

    // Check if account already exists
    if (File("./storage").exists()) {
        println("⚠️  Account already exists in ./storage/")
        println("💡 Delete ./storage/ to create fresh account")
        return
    }

    // Create new account
    println("🔧 Creating new Self account...")
    val account = createAccount()
    
    // Display account info
    println("\n📋 Account Information:")
    val address = getAccountAddress(account)
    println("🆔 Account DID: $address")
    println("🔐 Status: Ready for secure communication")
    
    println("\n✅ New Self account ready!")
}

fun createAccount(): Account {
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
        onConnect = { println("✅ Connected to Self network") },
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

fun getAccountAddress(account: Account): String {
    var address = ""
    account.inboxOpen { status, addr -> 
        if (status.success()) address = addr.encodeHex() 
    }
    Thread.sleep(1000) // Simple wait
    return address
} 
