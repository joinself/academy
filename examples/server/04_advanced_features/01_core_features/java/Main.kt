import com.joinself.selfsdk.kmp.account.Account
import com.joinself.selfsdk.kmp.error.SelfStatus
import com.joinself.selfsdk.kmp.event.*
import com.joinself.selfsdk.kmp.keypair.signing.PublicKey
import com.joinself.selfsdk.kmp.message.Chat
import com.joinself.selfsdk.kmp.message.ChatBuilder
import com.joinself.selfsdk.kmp.message.ContentType
import com.joinself.selfsdk.kmp.message.DiscoveryRequestBuilder
import com.joinself.selfsdk.kmp.time.Timestamp
import kotlinx.serialization.Serializable
import kotlinx.serialization.encodeToString
import kotlinx.serialization.json.Json
import java.time.Duration
import java.time.Instant
import java.time.ZoneId
import java.time.format.DateTimeFormatter
import java.util.concurrent.ConcurrentHashMap
import java.util.concurrent.Executors
import java.util.concurrent.Semaphore
import java.util.concurrent.TimeUnit
import java.util.concurrent.atomic.AtomicInteger

// Data classes for clean JSON serialization
@Serializable
data class UserProfile(val name: String, val role: String)

@Serializable
data class AppSettings(val darkMode: Boolean, val notifications: Boolean)

@Serializable
data class ConnectionInfo(
    val peerId: String,
    val connectedAt: String,
    val connectionType: String,
    val status: String
)

@Serializable
data class MessageInfo(
    val sender: String,
    val content: String,
    val timestamp: String,
    val messageId: Int,
    val type: String
)

/**
 * Manages the advanced features demonstration, including state and lifecycle.
 */
class AdvancedDemo {
    private lateinit var account: Account
    private val connections = ConcurrentHashMap<String, Boolean>()
    private val messageCount = AtomicInteger(0)
    private val startTime: Instant = Instant.now()
    private val doneSignal = Semaphore(0)
    private val json = Json { prettyPrint = true; ignoreUnknownKeys = true }
    private val scheduler = Executors.newSingleThreadScheduledExecutor()

    /**
     * Orchestrates the complete advanced features demonstration.
     */
    fun run() {
        println("🚀 Advanced Self SDK Features Demo (Kotlin SDK)")
        println("==============================================")
        println("This demo showcases advanced Self SDK capabilities using the core SDK directly.")
        println()

        try {
            setupAdvancedAccount()
            demonstrateAdvancedStorage()
            generateDiscoveryQR()
            setupGracefulShutdown()
            waitForInteractions()
        } catch (e: Exception) {
            println("❌ An unexpected error occurred: ${e.message}")
            e.printStackTrace()
        } finally {
            displaySummary()
            account.destroyAccount()
            scheduler.shutdownNow()
            println("\n✅ Advanced features demo completed!")
        }
    }


    /**
     * Creates an account with advanced configuration, including custom storage and callbacks.
     */
    private fun setupAdvancedAccount() {
        println("🔧 Setting up advanced account with core SDK...")

        account = Common.setupAccount(storagePath = "./advanced_demo_storage", callbacks = object :Common.Callbacks {
            override fun onWelcome(account: Account, welcome: Welcome) {
                onWelcomeMessage(welcome)
            }

            override fun onKeyPackage(account: Account, keyPackage: KeyPackage) {
                onKeyPackageMessage(keyPackage)
            }

            override fun onMessage(account: Account, msg: Message) {
                onMessageReceived(msg)
            }
        })
        val inbox = Common.openInbox(account)
        if (inbox == null) {
            throw IllegalStateException("Failed to open inbox for the advanced account.")
        }

        println("✅ Advanced account created successfully")
        println("🆔 Account DID: ${inbox.encodeHex()}")
        println()
    }

    /**
     * Demonstrates storing and retrieving various data types using the encrypted value store.
     */
    private fun demonstrateAdvancedStorage() {
        println("💾 Advanced Storage Capabilities")
        println("================================")
        println("🔹 Encrypted Storage Operations")

        val storageData = mapOf(
            "user_profile" to json.encodeToString(UserProfile("Advanced User", "Developer")),
            "app_settings" to json.encodeToString(AppSettings(true, true)),
            "session_token" to "token_${System.currentTimeMillis()}",
            "last_activity" to Instant.now().toString(),
            "connection_count" to "0"
        )

        storageData.forEach { (key, value) ->
            runCatching {
                account.valueStore(key, value.encodeToByteArray())
                println("   ✅ Stored '$key'")
            }.onFailure {
                println("   ❌ Failed to store '$key': ${it.message}")
            }
        }

        println("🔹 Data Retrieval Verification")
        runCatching {
            val profileData = account.valueLookup("user_profile")?.decodeToString()
            val profile = profileData?.let { json.decodeFromString<UserProfile>(it) }
            println("   📖 Retrieved profile: ${profile?.name} (${profile?.role})")

            val settingsData = account.valueLookup("app_settings")?.decodeToString()
            val settings = settingsData?.let { json.decodeFromString<AppSettings>(it) }
            println("   🎨 Dark mode: ${settings?.darkMode}, Notifications: ${settings?.notifications}")
        }.onFailure {
            println("   ❌ Failed to retrieve or parse stored data: ${it.message}")
        }
        println("✅ Storage operations completed successfully")
    }

    /**
     * Generates a discovery QR code for establishing secure connections with peers.
     */
    private fun generateDiscoveryQR() {
        println("🔍 Discovery & Connection Setup")
        println("===============================")

        runCatching {
            val inboxAddress = Common.openInbox(account) ?: throw IllegalStateException("Inbox not available")
            val expires = Timestamp.now() + Duration.ofMinutes(15).toMillis()

            val keyPackage = account.connectionNegotiateOutOfBand(inboxAddress, expires)
            val discoveryContent = DiscoveryRequestBuilder()
                .keyPackage(keyPackage)
                .expires(expires)
                .finish()

            val anonymousMsg = AnonymousMessage.fromContent(discoveryContent)
            anonymousMsg.setFlags(FlagSet(Flag.TARGET_SANDBOX))

            val qrCode = anonymousMsg.encodeQR(QrEncoding.UNICODE).toString(Charsets.UTF_8)

            println("📱 QR Code for Discovery:")
            println("========================")
            println(qrCode)
            println("========================")
            println("\n🎯 Scan this QR code with the Self app to connect!")
            val expirationTime = Instant.now().plus(Duration.ofMinutes(15))
            val formatter = DateTimeFormatter.ofPattern("HH:mm:ss").withZone(ZoneId.systemDefault())
            println("⏰ Expires in 15 minutes (${formatter.format(expirationTime)})")

            println("💡 What happens when someone connects:")
            println("   • Secure connection established")
            println("   • Welcome message sent automatically")
            println("   • Advanced storage updated with connection info")
            println("   • Demo showcases real-time message handling")
        }.onFailure {
            println("❌ Failed to generate discovery QR code: ${it.message}")
        }
        println()
    }

    /**
     * Handles new connection welcome messages from peers.
     */
    private fun onWelcomeMessage(welcomeMsg: Welcome) {
        val signal = Semaphore(0)
        val peerAddress = welcomeMsg.fromAddress()
        connections[peerAddress.encodeHex()] = true

        println("\n🎉 New connection established!")
        println("👤 Peer: ${peerAddress.encodeHex()}")
        println("🕐 Time: ${DateTimeFormatter.ofPattern("HH:mm:ss").withZone(ZoneId.systemDefault()).format(Instant.now())}")

        runCatching {
            account.connectionAccept(welcomeMsg.toAddress(), welcomeMsg.welcome()) { status: SelfStatus, groupAddress: PublicKey ->
                println("✅ Connection accepted.")
                signal.release()
            }
            signal.acquire()

            val connectionInfo = ConnectionInfo(
                peerId = peerAddress.encodeHex(),
                connectedAt = Instant.now().toString(),
                connectionType = "discovery",
                status = "active"
            )
            val storageKey = "connection_${peerAddress.encodeHex().substring(0, 8)}"
            account.valueStore(storageKey, json.encodeToString(connectionInfo).toByteArray())
            println("💾 Connection info stored in advanced storage.")

            account.valueStore("connection_count", connections.size.toString().toByteArray())

            sendAdvancedWelcomeMessage(peerAddress)
        }.onFailure {
            println("❌ Failed to process welcome message: ${it.message}")
        }
        println()
    }

    private fun onKeyPackageMessage(keyPackageMsg: KeyPackage) {
        println("🔑 Key package received from ${keyPackageMsg.fromAddress().encodeHex()}")
    }

    /**
     * Handles incoming messages, filtering for chat content and sending responses.
     */
    private fun onMessageReceived(msg: Message) {
        if (msg.content().contentType() != ContentType.CHAT) {
            println("🔍 DEBUG: Skipping non-chat message (type: ${msg.content().contentType()})")
            return
        }

        val count = messageCount.incrementAndGet()
        val sender = msg.fromAddress()
        val chat = Chat.decode(msg.content())
        val content = chat.message()

        println("📨 Chat Message #$count received")
        println("👤 From: ${sender.encodeHex()}")
        println("💬 Content: $content")

        runCatching {
            val messageInfo = MessageInfo(
                sender = sender.encodeHex(),
                content = content,
                timestamp = Instant.now().toString(),
                messageId = count,
                type = "chat"
            )
            val storageKey = "message_$count"
            account.valueStore(storageKey, json.encodeToString(messageInfo).toByteArray())
            println("💾 Message stored in advanced storage.")

            sendAdvancedResponse(sender, count)
        }.onFailure {
            println("⚠️ Failed to process incoming message: ${it.message}")
        }
        println()
    }

    private fun sendAdvancedWelcomeMessage(peer: PublicKey) {
        val welcomeText = """
        🚀 Welcome to the Advanced Self SDK Demo!

        This demo showcases advanced Self SDK features:
        ✅ Core SDK account management
        ✅ Encrypted storage with automatic security
        ✅ Real-time message handling
        ✅ Advanced connection tracking

        Connected at: ${DateTimeFormatter.ofPattern("HH:mm:ss").withZone(ZoneId.systemDefault()).format(Instant.now())}
        Active connections: ${connections.size}

        Send any message to see more features!
        """.trimIndent()

        sendChat(peer, welcomeText, "welcome")
    }

    private fun sendAdvancedResponse(peer: PublicKey, messageCount: Int) {
        val responseText = """
        📊 Advanced Demo Stats:

        Message #$messageCount processed ✅
        Active connections: ${connections.size}
        Demo runtime: ${Duration.between(startTime, Instant.now()).seconds}s
        Storage operations: Working ✅
        Real-time processing: Working ✅
        """.trimIndent()

        sendChat(peer, responseText, "response")
    }

    private fun sendChat(peer: PublicKey, text: String, type: String) {
        runCatching {
            val chatContent = ChatBuilder().message(text).finish()
            account.messageSend(peer, chatContent)
            println("📤 Advanced $type message sent.")
        }.onFailure {
            println("❌ Failed to send $type message: ${it.message}")
        }
    }

    /**
     * Configures a clean shutdown on Ctrl+C.
     */
    private fun setupGracefulShutdown() {
        Runtime.getRuntime().addShutdownHook(Thread {
            println("\n🛑 Shutdown signal received...")
            doneSignal.release()
        })
    }

    /**
     * Waits for connections and messages, with a status ticker and a timeout.
     */
    private fun waitForInteractions() {
        println("⏳ Waiting for connections and messages...")
        println("   • Scan the QR code above to connect")
        println("   • Press Ctrl+C or wait 5 minutes to finish the demo")
        println()

        // Schedule a periodic status update
        val statusFuture = scheduler.scheduleAtFixedRate({
            val runtime = Duration.between(startTime, Instant.now()).seconds
            println("📊 Status: ${connections.size} connections, ${messageCount.get()} messages (runtime: ${runtime}s)")
        }, 30, 30, TimeUnit.SECONDS)

        // Wait for the done signal or a 5-minute timeout
        val completed = doneSignal.tryAcquire(5, TimeUnit.MINUTES)
        statusFuture.cancel(true)

        if (!completed) {
            println("\n⏰ Demo timeout reached (5 minutes)")
        } else {
            println("\n🏁 Demo completed by user")
        }
    }

    /**
     * Shows final statistics and a summary of demonstrated features.
     */
    private fun displaySummary() {
        val runtime = Duration.between(startTime, Instant.now())
        val storageOps = connections.size * 2 + messageCount.get() + 5 // Approximate

        println("\n📊 Advanced Demo Summary")
        println("========================")
        println("⏱️  Total runtime: ${runtime.seconds} seconds")
        println("🔗 Connections made: ${connections.size}")
        println("📨 Messages processed: ${messageCount.get()}")
        println("💾 Storage operations: $storageOps+")
        println()
        println("🎓 What was demonstrated:")
        println("   ✅ Core SDK account creation and configuration")
        println("   ✅ Advanced encrypted storage with automatic security")
        println("   ✅ Discovery QR generation and connection handling")
        println("   ✅ Real-time event-driven message processing")
        println("   ✅ Graceful shutdown and resource management")

        println("🚀 Advanced features unlocked:")
        println("   • Direct core SDK usage")
        println("   • Encrypted storage for secure data persistence")
        println("   • Event-driven architecture with callbacks")
        println("   • Real-time connection and message handling")
        println("   • Production patterns for robust applications")
        println()

        println("📚 Next steps:")
        println("   • Explore individual subdirectories for deep dives")
        println("   • Build production applications using these patterns")
        println("   • Integrate advanced features into your own projects")
        println()

        println("✅ Advanced features demo completed!")
    }
}

fun main() {
    AdvancedDemo().run()
}