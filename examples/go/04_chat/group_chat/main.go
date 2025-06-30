// Package main demonstrates group chat functionality using the Self SDK
package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("👥 Group Chat Demo")
	fmt.Println("==================")

	// Create multiple clients (admin and members)
	admin, member1, member2 := createClients()
	defer admin.Close()
	defer member1.Close()
	defer member2.Close()

	fmt.Printf("👑 Admin: %s\n", admin.DID())
	fmt.Printf("👤 Member1: %s\n", member1.DID())
	fmt.Printf("👤 Member2: %s\n", member2.DID())
	fmt.Println()

	// Set up group event handlers
	setupGroupHandlers(admin, member1, member2)

	// Create a group chat
	group := createGroup(admin)

	// Establish peer connections
	establishConnections(admin, member1, member2)

	// Invite members to the group
	inviteMembers(admin, group, member1, member2)

	// Demonstrate group messaging
	demonstrateGroupMessaging(admin, group)

	// Show group management features
	demonstrateGroupManagement(admin, group)

	fmt.Println("✅ Group chat demo completed! Press Ctrl+C to exit.")

	// Keep running to demonstrate ongoing group capabilities
	select {}
}

// createClients sets up the admin and member clients for group chat
func createClients() (*client.Client, *client.Client, *client.Client) {
	fmt.Println("🔧 Setting up group chat clients...")

	admin, err := client.NewSimplified("./group_admin_storage")
	if err != nil {
		log.Fatal("Failed to create admin client:", err)
	}

	member1, err := client.NewSimplified("./group_member1_storage")
	if err != nil {
		log.Fatal("Failed to create member1 client:", err)
	}

	member2, err := client.NewSimplified("./group_member2_storage")
	if err != nil {
		log.Fatal("Failed to create member2 client:", err)
	}

	fmt.Println("✅ All clients created successfully")
	return admin, member1, member2
}

// setupGroupHandlers configures event handlers for all group activities
func setupGroupHandlers(admin, member1, member2 *client.Client) {
	fmt.Println("📨 Setting up group event handlers...")
	setupClientHandlers(admin, "👑 Admin")
	setupClientHandlers(member1, "👤 Member1")
	setupClientHandlers(member2, "👤 Member2")
	fmt.Println("✅ Group handlers configured")
	fmt.Println()
}

// setupClientHandlers configures group event handlers for a specific client
func setupClientHandlers(selfClient *client.Client, clientName string) {
	// Handle incoming group messages
	selfClient.GroupChats().OnGroupMessage(func(msg client.GroupChatMessage) {
		timestamp := time.Now().Format("15:04:05")
		fmt.Printf("📨 [%s] %s in '%s': \"%s\"\n",
			timestamp, msg.From(), msg.GroupName(), msg.Text())
	})

	// Handle group invitations
	selfClient.GroupChats().OnGroupInvite(func(invitation *client.GroupChatInvitation) {
		fmt.Printf("📧 [%s] Invitation from %s to join '%s'\n",
			clientName, invitation.InviterDID, invitation.GroupName)

		err := invitation.Accept()
		if err != nil {
			fmt.Printf("❌ Failed to accept: %v\n", err)
		} else {
			fmt.Printf("✅ [%s] Joined group: %s\n", clientName, invitation.GroupName)
		}
	})

	// Handle member join events
	selfClient.GroupChats().OnMemberJoined(func(groupID string, member *client.GroupMember) {
		fmt.Printf("👋 [%s] %s joined group (%s)\n", clientName, member.DID, member.Role)
	})

	// Handle group creation events
	selfClient.GroupChats().OnGroupCreated(func(group *client.GroupChat) {
		fmt.Printf("🎉 [%s] Group created: %s\n", clientName, group.Name())
	})
}

// createGroup demonstrates group creation with admin privileges
func createGroup(admin *client.Client) *client.GroupChat {
	fmt.Println("📋 Creating group chat...")

	group, err := admin.GroupChats().CreateGroup("Dev Team", "Daily standup and project discussions")
	if err != nil {
		log.Fatal("Failed to create group:", err)
	}

	fmt.Printf("✅ Group created: %s (ID: %s, Members: %d)\n",
		group.Name(), group.ID(), group.MemberCount())
	fmt.Println()

	return group
}

// establishConnections simulates peer discovery between clients
func establishConnections(admin, member1, member2 *client.Client) {
	fmt.Println("🔗 Establishing peer connections...")
	time.Sleep(2 * time.Second)
	fmt.Println("✅ Peer connections established")
	fmt.Println()
}

// inviteMembers demonstrates the group invitation process
func inviteMembers(admin *client.Client, group *client.GroupChat, member1, member2 *client.Client) {
	fmt.Println("👥 Inviting members to group...")

	err := admin.GroupChats().InviteToGroup(group.ID(), member1.DID(), "Welcome to our dev team group!")
	if err != nil {
		log.Printf("Failed to invite Member1: %v", err)
	} else {
		fmt.Println("📤 Invitation sent to Member1")
	}

	time.Sleep(1 * time.Second)

	err = admin.GroupChats().InviteToGroup(group.ID(), member2.DID(), "Join our daily discussions!")
	if err != nil {
		log.Printf("Failed to invite Member2: %v", err)
	} else {
		fmt.Println("📤 Invitation sent to Member2")
	}

	time.Sleep(3 * time.Second)
	fmt.Println()
}

// demonstrateGroupMessaging shows group message broadcasting
func demonstrateGroupMessaging(admin *client.Client, group *client.GroupChat) {
	fmt.Println("💬 Demonstrating group messaging...")

	messages := []string{
		"🎉 Hello everyone! Welcome to our dev team group.",
		"Let's use this for daily standups and project updates.",
		"Daily standup in 5 minutes!",
		"Great work everyone! 🚀",
	}

	for _, msg := range messages {
		err := admin.GroupChats().SendToGroup(group.ID(), msg)
		if err != nil {
			fmt.Printf("❌ Failed to send message: %v\n", err)
		} else {
			fmt.Printf("📤 Sent: \"%s\"\n", msg)
		}
		time.Sleep(1 * time.Second)
	}
	fmt.Println()
}

// demonstrateGroupManagement shows group administration features
func demonstrateGroupManagement(admin *client.Client, group *client.GroupChat) {
	fmt.Println("⚙️ Demonstrating group management...")

	err := group.UpdateName("Dev Team - Sprint 1")
	if err != nil {
		log.Printf("Failed to update group name: %v", err)
	} else {
		fmt.Println("✅ Group name updated")
	}

	time.Sleep(1 * time.Second)

	err = group.UpdateDescription("Sprint 1 planning and daily standups")
	if err != nil {
		log.Printf("Failed to update description: %v", err)
	} else {
		fmt.Println("✅ Group description updated")
	}

	adminGroups := admin.GroupChats().ListGroups()
	fmt.Printf("📋 Admin manages %d group(s)\n", len(adminGroups))
	fmt.Println()
}
