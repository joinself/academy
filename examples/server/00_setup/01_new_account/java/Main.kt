import com.joinself.selfsdk.kmp.account.Account
import com.joinself.selfsdk.kmp.account.LogLevel
import com.joinself.selfsdk.kmp.account.Target
import java.io.File
import java.util.concurrent.Semaphore

val signal = Semaphore(0)

fun main() {
    println("🆕 New Account Creation Example")
    println("===============================")

    // Check if account already exists
    println("🔍 Checking for existing account...")
    if (File("./storage").exists()) {
        println("⚠️  Account already exists in ./storage/")
        println("💡 Delete ./storage/ to create fresh account, or use '../02_existing_account'")
        return
    }

    // Create new account
    println("🔧 Creating new Self account...")
    val account = createAccount()
    
    // Display account information
    println("\n📋 New Account Information")
    println("==========================")
    val address = getAccountAddress(account)
    println("🆔 Account DID: $address")
    println("📬 Inbox Address: $address")
    println("🔐 Status: Encrypted and ready for secure communication")
    
    println("\n💡 Your DID can be shared with others for connections")
    
    println("\n📁 Storage: ./storage/ (encrypted)")
    println("🔄 Use '../02_existing_account' to reload this account")
    
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
        onConnect = { 
            println("✅ Connected to Self network")
            println("✅ New account created successfully!")
            signal.release()
        },
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

    signal.acquire() // Simple wait for connection
    return account
}

fun getAccountAddress(account: Account): String {
    var address = ""
    account.inboxOpen { status, addr -> 
        if (status.success()) address = addr.encodeHex()
        signal.release()
    }
    signal.acquire()
    return address
} 
