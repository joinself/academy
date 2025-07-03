import com.joinself.selfsdk.kmp.account.Account

fun main() {
    println("📧 Inbox Access Example")
    println("========================")

    // Load or create account
    println("🔧 Setting up account...")
    val account = Common.setupAccount()
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

    println("\nPress enter to exit")
    readln()
}

fun getInboxAddress(account: Account): String {
    var address = ""
    account.inboxOpen { status, addr -> 
        if (status.success()) {
            address = addr.encodeHex()
        } else {
            println("❌ Failed to open inbox")
        }
        signal.release()
    }
    signal.acquire()
    return address
} 
