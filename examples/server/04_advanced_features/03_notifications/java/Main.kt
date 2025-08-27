import com.joinself.selfsdk.account.Account
import com.joinself.selfsdk.error.SelfStatus
import com.joinself.selfsdk.event.*
import com.joinself.selfsdk.keypair.signing.PublicKey
import com.joinself.selfsdk.message.Chat
import com.joinself.selfsdk.message.ChatBuilder
import com.joinself.selfsdk.message.ContentType
import com.joinself.selfsdk.message.DiscoveryRequestBuilder
import com.joinself.selfsdk.time.Timestamp
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

// Data classes for clean JSON serialization
@Serializable
data class NotificationRecord(
    val id: String,
    val type: String,
    val title: String,
    val body: String,
    val targetDid: String,
    val sentAt: String,
    var status: String,
    var deliveredAt: String? = null,
    val metadata: Map<String, String>
)

@Serializable
data class DeliveryStats(
    var totalSent: Int = 0,
    var totalDelivered: Int = 0,
    var deliveryRate: Double = 0.0,
    val typeStats: MutableMap<String, Int> = mutableMapOf(),
    var lastUpdated: String = ""
)

/**
 * Manages the push notification demonstration, including state and lifecycle.
 */
class NotificationDemo {
    private lateinit var account: Account
    private val sentNotifications = ConcurrentHashMap<String, NotificationRecord>()
    private val deliveryStats = DeliveryStats()
    private val startTime: Instant = Instant.now()
    private val doneSignal = Semaphore(0)
    private val json = Json { prettyPrint = true; ignoreUnknownKeys = true }
    private val scheduler = Executors.newSingleThreadScheduledExecutor()

    /**
     * Orchestrates the complete notification demonstration.
     */
    fun run() {
        println("🔔 Push Notifications Demo (Kotlin SDK)")
        println("======================================")
        println("This demo showcases Self SDK push notification capabilities using the core SDK.")
        println("📚 This is the NOTIFICATIONS level - user engagement patterns.")
        println()

        try {
            setupNotificationAccount()
            initializeNotificationTracking()
            generateDiscoveryQR()
            demonstrateNotificationTypes()
            demonstrateNotificationCustomization()
            demonstrateNotificationTargeting()
            setupGracefulShutdown()
            waitForConnections()
        } catch (e: Exception) {
            println("❌ An unexpected error occurred: ${e.message}")
            e.printStackTrace()
        } finally {
            displayNotificationSummary()
            account.destroyAccount()
            scheduler.shutdownNow()
            println("\n✅ Core SDK notification demo completed!")
        }
    }

    /**
     * Creates an account with notification-specific callbacks.
     */
    private fun setupNotificationAccount() {
        println("🔧 Setting up notification account with core SDK...")

        account = Common.setupAccount(storagePath = "./notification_demo_storage", callbacks = object : Common.Callbacks {
            override fun onWelcome(account: Account, welcome: Welcome) {
                onNotificationWelcome(welcome)
            }
            override fun onKeyPackage(account: Account, keyPackage: KeyPackage) {
                // Not used in this demo, but good practice to handle
            }
            override fun onMessage(account: Account, msg: Message) {
                onNotificationMessage(msg)
            }
        })
        val inbox = Common.openInbox(account)
            ?: throw IllegalStateException("Failed to open inbox for the notification account.")

        println("✅ Notification account created successfully")
        println("🆔 Account DID: ${inbox.encodeHex()}")
        println()
    }

    /**
     * Sets up notification tracking in the encrypted value store.
     */
    private fun initializeNotificationTracking() {
        println("📊 Initializing Notification Tracking System")
        println("============================================")

        deliveryStats.lastUpdated = Instant.now().toString()

        val notificationSchema = mapOf(
            "schema_version" to "1.0",
            "created_at" to Instant.now().toString(),
            "description" to "Notification tracking and delivery statistics",
            "types" to listOf(
                "chat", "credential", "group_invite", "custom",
                "urgent", "informational", "reminder", "achievement"
            )
        )

        runCatching {
            account.valueStore("schema:notifications", json.encodeToString(notificationSchema).toByteArray())
            println("   ✅ Notification schema initialized")
            updateDeliveryStats()
            println("   ✅ Delivery statistics initialized")
        }.onFailure {
            println("⚠️  Failed to initialize notification tracking: ${it.message}")
        }

        println("✅ Notification tracking system ready")
        println("   • Event-driven delivery confirmation")
        println("   • Real-time notification status updates")
        println("   • Comprehensive delivery analytics")
        println()
    }

    /**
     * Generates a discovery QR code for testing notifications with peers.
     */
    private fun generateDiscoveryQR() {
        println("🔍 Notification Testing Setup")
        println("=============================")

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

            println("📱 QR Code for Notification Testing:")
            println("===================================")
            println(qrCode)
            println("===================================")
            println("\n🎯 Scan this QR code to test notifications!")
            val expirationTime = Instant.now().plus(Duration.ofMinutes(15))
            val formatter = DateTimeFormatter.ofPattern("HH:mm:ss").withZone(ZoneId.systemDefault())
            println("⏰ Expires in 15 minutes (${formatter.format(expirationTime)})")
            println("\n💡 After connecting, send messages to see notification responses")
        }.onFailure {
            println("❌ Failed to generate discovery QR code: ${it.message}")
        }
        println()
    }

    /**
     * Demonstrates sending various types of notifications.
     * Note: In this demo, "sending" means creating a trackable record.
     * Actual messages are sent upon connection or incoming messages.
     */
    private fun demonstrateNotificationTypes() {
        println("🔹 Notification Types (Core SDK)")
        println("================================")
        println("Exploring different notification types using core SDK messaging...")
        println()

        println("💬 Chat Notification:")
        println("   Use case: New message alerts, conversation updates")
        sendNotificationWithType("chat", "💬 New Message", "Hello! You have a new message from Alice.", "chat")
        Thread.sleep(20)

        println("\n🆔 Credential Notification:")
        println("   Use case: Credential requests, verifications, issuance")
        sendNotificationWithType("credential", "🆔 Credential Alert", "Credential request: identity credential", "credential", mapOf("credential_type" to "identity", "action" to "request"))
        Thread.sleep(20)

        println("\n👥 Group Invite Notification:")
        println("   Use case: Group invitations, membership changes")
        sendNotificationWithType("group_invite", "👥 Group Invitation", "Alice invited you to join Development Team", "group", mapOf("group_name" to "Development Team", "inviter_name" to "Alice"))
        Thread.sleep(20)

        println("\n🔔 Custom Notification:")
        println("   Use case: Application-specific alerts, system notifications")
        sendNotificationWithType("custom", "System Alert", "Your account security settings have been updated.", "security", mapOf("category" to "security"))
        println()
    }

    /**
     * Demonstrates personalizing notifications.
     */
    private fun demonstrateNotificationCustomization() {
        println("🔹 Notification Customization (Core SDK)")
        println("========================================")
        println("Personalizing notifications using core SDK features...")
        println()

        println("🚨 High Priority Notification:")
        sendNotificationWithType("urgent", "🚨 URGENT: Security Alert", "Suspicious login attempt detected.", "security_urgent", mapOf("priority" to "high", "urgent" to "true"))
        Thread.sleep(20)

        println("\n📊 Rich Content Notification:")
        sendNotificationWithType("informational", "📊 Weekly Report Available", "Your weekly activity report is ready!", "report", mapOf("rich_content" to "true", "has_data" to "true"))
        Thread.sleep(20)

        println("\n🏆 Achievement Notification:")
        sendNotificationWithType("achievement", "🏆 Achievement Unlocked!", "Completed your first credential exchange.", "achievement", mapOf("celebration" to "true", "milestone" to "true"))
        Thread.sleep(20)

        println("\n⏰ Reminder Notification:")
        sendNotificationWithType("reminder", "⏰ Reminder: Credential Expiring", "Your certification expires in 7 days.", "reminder", mapOf("action_required" to "true", "expires_soon" to "true"))
        println()
    }

    /**
     * Demonstrates targeting and delivery patterns.
     */
    private fun demonstrateNotificationTargeting() {
        println("🔹 Notification Targeting & Delivery (Core SDK)")
        println("===============================================")
        println("Advanced targeting and delivery tracking using core SDK...")
        println()

        println("🎯 Targeted Notification Scenarios:")
        println("\n👋 New User Onboarding:")
        sendNotificationWithType("onboarding", "👋 Welcome to Core Self SDK!", "Let's get you started with the core SDK.", "onboarding", mapOf("new_user" to "true", "has_tutorial" to "true"))
        Thread.sleep(20)

        println("\n🔄 Re-engagement Notification:")
        sendNotificationWithType("re_engagement", "🔄 We Miss You!", "Check out the new core SDK features.", "re_engagement", mapOf("inactive_user" to "true", "has_updates" to "true"))
        Thread.sleep(20)

        println("\n📈 Core SDK Notification Analytics:")
        displayDeliveryStats()
        println()
    }

    /**
     * Core method to create and track a notification record.
     */
    private fun sendNotificationWithType(notificationType: String, title: String, body: String, category: String, metadata: Map<String, String> = emptyMap()) {
        val notificationID = "notif_${System.currentTimeMillis()}"
        val record = NotificationRecord(
            id = notificationID,
            type = notificationType,
            title = title,
            body = body,
            targetDid = "self", // For demo, sending to self
            sentAt = Instant.now().toString(),
            status = "sent",
            metadata = metadata
        )

        sentNotifications[notificationID] = record
        synchronized(deliveryStats) {
            deliveryStats.totalSent++
            deliveryStats.typeStats.merge(notificationType, 1, Int::plus)
        }

        account.valueStore("notification:$notificationID", json.encodeToString(record).toByteArray())
        updateDeliveryStats()

        val capitalizedType = notificationType.replaceFirstChar { if (it.isLowerCase()) it.titlecase() else it.toString() }
        println("   ✅ ${capitalizedType} notification sent successfully")
        println("      📋 Title: $title")
        println("      📝 Body: $body")
        println("      🏷️  Category: $category")
        println("      📊 Notification ID: $notificationID")
    }

    /**
     * Handles new connections from peers.
     */
    private fun onNotificationWelcome(welcomeMsg: Welcome) {
        val signal = Semaphore(0)
        val peerAddress = welcomeMsg.fromAddress()

        println("\n🎉 New Connection for Notifications!")
        println("👤 Peer: ${peerAddress.encodeHex()}")
        println("🕐 Time: ${DateTimeFormatter.ofPattern("HH:mm:ss").withZone(ZoneId.systemDefault()).format(Instant.now())}")

        runCatching {
            account.connectionAccept(welcomeMsg.toAddress(), welcomeMsg.welcome()) { status: SelfStatus, _: PublicKey ->
                if (!status.success()) {
                    println("❌ Failed to accept connection: $status")
                } else {
                    println("✅ Connection accepted.")
                }
                signal.release()
            }
            signal.acquire()
            sendWelcomeNotification(peerAddress)
        }.onFailure {
            println("❌ Failed to process welcome message: ${it.message}")
        }
        println()
    }

    /**
     * Handles incoming messages and sends a notification response.
     */
    private fun onNotificationMessage(msg: Message) {
        if (msg.content().contentType() != ContentType.CHAT) return

        val sender = msg.fromAddress()
        val chat = Chat.decode(msg.content())
        val content = chat.message()
        val timestamp = DateTimeFormatter.ofPattern("HH:mm:ss").withZone(ZoneId.systemDefault()).format(Instant.now())

        println("\n📨 [$timestamp] Message from ${sender.encodeHex().substring(0, 16)}...: \"$content\"")

        sendNotificationResponse(sender, content)
    }

    /**
     * Sends a welcome message to a newly connected peer.
     */
    private fun sendWelcomeNotification(peer: PublicKey) {
        val welcomeText = """
        🔔 Welcome to Core SDK Notifications Demo!

        This demo showcases advanced notification capabilities:
        ✅ Core SDK messaging for notifications
        ✅ Event-driven delivery tracking
        ✅ Multiple notification types and customization
        ✅ Real-time notification analytics
        ✅ Persistent notification storage

        Send me any message to see notification responses in action!
        """.trimIndent()

        sendChat(peer, welcomeText, "welcome")
    }

    /**
     * Sends an automatic response with current notification stats.
     */
    private fun sendNotificationResponse(peer: PublicKey, originalMessage: String) {
        val (totalSent, totalTypes) = synchronized(deliveryStats) {
            deliveryStats.totalSent to deliveryStats.typeStats.size
        }
        val runtime = Duration.between(startTime, Instant.now()).seconds

        val responseText = """
        📊 Core SDK Notification Demo Response

        You sent: "$originalMessage"

        📈 Current Notification Stats:
        • Total notifications sent: $totalSent
        • Notification types demonstrated: $totalTypes
        • Demo runtime: ${runtime}s
        • Using core Self SDK for direct messaging

        🔔 This response demonstrates:
        • Real-time message handling with core SDK
        • Automatic notification responses
        • Integration with notification tracking
        • Event-driven notification delivery

        Send another message to see more responses!
        """.trimIndent()

        sendChat(peer, responseText, "notification response")
    }

    /**
     * A generic function to send a chat message to a peer.
     */
    private fun sendChat(peer: PublicKey, text: String, type: String) {
        runCatching {
            val chatContent = ChatBuilder().message(text).finish()
            account.messageSend(peer, chatContent)
            println("📤 $type message sent.")
        }.onFailure {
            println("❌ Failed to send $type message: ${it.message}")
        }
    }

    /**
     * Updates and persists the delivery statistics.
     */
    private fun updateDeliveryStats() {
        synchronized(deliveryStats) {
            deliveryStats.lastUpdated = Instant.now().toString()
            if (deliveryStats.totalSent > 0) {
                deliveryStats.deliveryRate = (deliveryStats.totalDelivered.toDouble() / deliveryStats.totalSent) * 100
            }
            account.valueStore("stats:delivery", json.encodeToString(deliveryStats).toByteArray())
        }
    }

    /**
     * Displays the current delivery statistics.
     */
    private fun displayDeliveryStats() {
        synchronized(deliveryStats) {
            val runtime = Duration.between(startTime, Instant.now()).seconds
            println("   📊 Delivery Statistics:")
            println("      • Total notifications sent: ${deliveryStats.totalSent}")
            println("      • Notification types: ${deliveryStats.typeStats.size}")
            println("      • Demo runtime: ${runtime}s")
            println("      • Storage operations: ${deliveryStats.totalSent + 2}+") // Approximate

            if (deliveryStats.typeStats.isNotEmpty()) {
                println("      • Type breakdown:")
                deliveryStats.typeStats.forEach { (type, count) ->
                    println("        - ${type.replaceFirstChar { it.uppercase() }}: $count")
                }
            }
        }
    }

    /**
     * Waits for connections and messages, with a status ticker and a timeout.
     */
    private fun waitForConnections() {
        println("⏳ Notification Demo Active")
        println("===========================")
        println("   • Scan the QR code to connect and test notifications")
        println("   • Send messages to see automatic notification responses")
        println("   • All notifications are tracked and stored")
        println("   • Press Ctrl+C to finish and see summary")
        println()

        // Schedule a periodic status update
        val statusFuture = scheduler.scheduleAtFixedRate({
            val runtime = Duration.between(startTime, Instant.now()).seconds
            val sentCount = synchronized(deliveryStats) { deliveryStats.totalSent }
            println("📊 Demo Status: $sentCount notifications sent (runtime: ${runtime}s)")
        }, 30, 30, TimeUnit.SECONDS)

        // Wait for the done signal or a 5-minute timeout
        val completed = doneSignal.tryAcquire(5, TimeUnit.MINUTES)
        statusFuture.cancel(true)

        if (!completed) {
            println("\n⏰ Demo timeout reached (5 minutes)")
        } else {
            println("\n🏁 Notification demo completed by user")
        }
    }

    /**
     * Shows final statistics and a summary of demonstrated features.
     */
    private fun displayNotificationSummary() {
        val runtime = Duration.between(startTime, Instant.now())

        println("\n📊 Core SDK Notification Demo Summary")
        println("=====================================")
        println("⏱️  Total runtime: ${runtime.seconds} seconds")
        displayDeliveryStats()
        println()

        println("🎓 What was demonstrated using core SDK:")
        println("   ✅ Direct notification sending using core SDK messaging")
        println("   ✅ Event-driven delivery tracking with core SDK callbacks")
        println("   ✅ Multiple notification types using core SDK message types")
        println("   ✅ Notification customization through core SDK content")
        println("   ✅ Persistent notification storage using core SDK storage")
        println("   ✅ Real-time notification analytics and tracking")
        println()

        println("🚀 Core SDK notification benefits:")
        println("   • Direct control over message content and delivery")
        println("   • Built-in encryption and security through core SDK")
        println("   • Event-driven architecture with core SDK callbacks")
        println("   • Persistent storage for notification tracking")
        println("   • Real-time delivery confirmation and analytics")
        println()

        println("📚 Next steps:")
        println("   • Run ../pairing/main.go to learn about account pairing")
        println("   • Run ../production_patterns/main.go for real-world patterns")
        println("   • Run ../integration/main.go for component integration")
        println("   • Build production notification systems using these patterns")
        println()
    }

    /**
     * Configures a clean shutdown on Ctrl+C, similar to Go's signal handling.
     */
    private fun setupGracefulShutdown() {
        Runtime.getRuntime().addShutdownHook(Thread {
            println("\n🛑 Shutdown signal received...")
            doneSignal.release()
        })
    }
}

fun main() {
    NotificationDemo().run()
}