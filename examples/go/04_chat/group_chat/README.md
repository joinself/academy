# Group Chat Example 🟡

> **⚠️ Note**: This example currently uses a non-existent client API and needs to be rewritten to use the actual Self SDK APIs like the main chat example. The concepts and workflow described below represent the intended functionality.

A comprehensive demonstration of group chat capabilities using the Self SDK, showing multi-participant messaging, role-based permissions, and group management features.

## 🚀 Quick Start

```bash
go run main.go
```

The demo creates 3 clients (admin + 2 members) and demonstrates the complete group chat workflow automatically.

## 🎯 What You'll Learn

- **Multi-Client Architecture**: Managing multiple Self identities for group communication
- **Group Creation & Administration**: Setting up groups with admin privileges
- **Member Management**: Invitation workflows and role-based permissions
- **Group Messaging**: Broadcasting messages to all participants
- **Event-Driven Communication**: Real-time group activity notifications
- **Administrative Features**: Group name/description updates and member management

## 🏗️ How Group Chat Works

### Step 1: Multi-Client Setup
```go
// Create separate accounts for admin and members
admin, err := createAccount("./group_admin_storage")
member1, err := createAccount("./group_member1_storage") 
member2, err := createAccount("./group_member2_storage")
```

**What happens:**
- Each participant needs their own Self identity
- Separate encrypted storage for each client
- Independent connection to Self network
- Different roles and permissions per client

### Step 2: Group Creation
```go
// Admin creates group with description
group, err := admin.CreateGroup("Dev Team", "Daily standup and project discussions")
```

**Group Properties:**
- **Name**: Human-readable group identifier
- **Description**: Purpose and context for the group
- **Admin**: Creator with management privileges
- **Member List**: Initially contains only the admin
- **Group ID**: Unique identifier for message routing

### Step 3: Member Invitation Process
```go
// Admin sends invitation with custom message
err := admin.InviteToGroup(group.ID(), member.DID(), "Welcome to our dev team!")

// Member receives and can accept/reject
invitation.Accept() // or invitation.Reject()
```

**Invitation Flow:**
1. **Admin Sends**: Creates invitation with group context
2. **Network Delivery**: Encrypted invitation sent via Self network  
3. **Member Receives**: Gets notification of invitation
4. **Member Decides**: Can accept or reject invitation
5. **Group Updated**: Successful acceptance adds member to group

### Step 4: Group Message Broadcasting
```go
// Admin or any member sends to group
err := admin.SendToGroup(group.ID(), "Hello everyone!")

// All members receive the message
OnGroupMessage(func(msg GroupMessage) {
    fmt.Printf("From %s: %s", msg.Sender(), msg.Text())
})
```

**Message Flow:**
1. **Sender**: Any group member creates message
2. **Broadcast**: Self network delivers to all group members
3. **Encryption**: End-to-end encrypted per recipient
4. **Callbacks**: Each member's app receives notification
5. **Display**: UI shows message with sender and timestamp

## 📋 Expected Output

### Initial Setup
```
👥 Group Chat Demo
==================

🔧 Setting up group chat clients...
✅ All clients created successfully
👑 Admin: did:self:admin123...
👤 Member1: did:self:member1456...
👤 Member2: did:self:member2789...

📨 Setting up group event handlers...
✅ Group handlers configured
```

### Group Creation & Management
```
📋 Creating group chat...
✅ Group created: Dev Team (ID: group_abc123, Members: 1)

🔗 Establishing peer connections...
✅ Peer connections established

👥 Inviting members to group...
📤 Invitation sent to Member1
📧 [👤 Member1] Invitation from did:self:admin123... to join 'Dev Team'
✅ [👤 Member1] Joined group: Dev Team
```

### Group Messaging
```
💬 Demonstrating group messaging...
📤 Sent: "🎉 Hello everyone! Welcome to our dev team group."
📨 [15:04:05] did:self:admin123... in 'Dev Team': "🎉 Hello everyone! Welcome to our dev team group."

📤 Sent: "Daily standup in 5 minutes!"
📨 [15:04:06] did:self:admin123... in 'Dev Team': "Daily standup in 5 minutes!"
```

## 🔍 Code Architecture

### Function Breakdown

| Function | Purpose | Key Concepts |
|----------|---------|--------------|
| `createClients()` | Multi-account setup | Independent identities, storage separation |
| `setupGroupHandlers()` | Event configuration | Message callbacks, invitation handlers |
| `createGroup()` | Group establishment | Admin privileges, group properties |
| `inviteMembers()` | Member onboarding | Invitation workflow, acceptance process |
| `demonstrateGroupMessaging()` | Message broadcasting | Multi-participant communication |
| `demonstrateGroupManagement()` | Administrative features | Name/description updates, member listing |

### Key Group Chat Concepts

**Roles & Permissions:**
- **Admin**: Group creator with full management rights
  - Create/delete groups
  - Invite/remove members  
  - Update group properties
  - Send messages
- **Members**: Participants with messaging rights
  - Send messages to group
  - Receive group messages
  - Leave group voluntarily

**Event Types:**
- **Group Messages**: Real-time message delivery to all members
- **Member Joined**: Notification when someone accepts invitation
- **Group Created**: Confirmation of successful group creation
- **Invitations**: Secure invitation delivery and response handling

**Group Properties:**
- **Persistent**: Groups survive client restarts
- **Encrypted**: All messages end-to-end encrypted per recipient
- **Scalable**: Support for multiple participants (implementation dependent)
- **Distributed**: No central server stores group data

## 🎓 What Just Happened

When you run this example:

1. **Multi-Identity Setup**: Three separate Self accounts are created with different roles
2. **Event Handler Configuration**: Each client registers callbacks for group activities
3. **Group Establishment**: Admin creates a named group with description
4. **Peer Discovery**: Clients establish direct connections (simulated)
5. **Member Invitation**: Admin sends encrypted invitations to potential members
6. **Acceptance Process**: Members receive and automatically accept invitations
7. **Group Messaging**: Messages are broadcast to all group participants
8. **Administrative Tasks**: Group properties are updated to show management capabilities

**Security Features:**
- Each message encrypted separately for each recipient
- Admin privileges enforced by Self network protocol
- No central server stores group membership or messages
- Private keys remain on each participant's device

## 🔧 Customization Ideas

### Add More Participants
```go
// Create additional members
member3, err := createAccount("./group_member3_storage")
member4, err := createAccount("./group_member4_storage")

// Invite them to existing groups
admin.InviteToGroup(group.ID(), member3.DID(), "Welcome to the team!")
```

### Implement Member Removal
```go
// Admin removes member from group
err := admin.RemoveFromGroup(group.ID(), member.DID())

// Handle removal events
OnMemberRemoved(func(groupID, memberDID string) {
    fmt.Printf("Member %s removed from group %s", memberDID, groupID)
})
```

### Create Multiple Groups
```go
// Admin can manage multiple groups
devTeam := admin.CreateGroup("Dev Team", "Development discussions")
projectGroup := admin.CreateGroup("Project Alpha", "Alpha project coordination")

// Members can be in multiple groups
admin.InviteToGroup(devTeam.ID(), member.DID(), "Join dev discussions")
admin.InviteToGroup(projectGroup.ID(), member.DID(), "Join project team")
```

### Add Message Threading
```go
// Reply to specific messages
replyMsg := CreateReply(originalMessage.ID(), "Great idea!")
admin.SendToGroup(group.ID(), replyMsg)
```

## 💡 Troubleshooting

### Common Issues

**Group Creation Fails:**
- Verify admin account is properly connected to Self network
- Check storage permissions for group data persistence
- Ensure unique group names within admin's scope

**Invitation Problems:**
- Confirm peer connections exist between admin and members
- Verify member accounts are online and responsive
- Check that invitation messages aren't being filtered

**Message Delivery Issues:**
- Ensure all members have accepted group invitations
- Verify group callbacks are properly registered
- Check that sender has permission to post to group

**Permission Errors:**
- Only admins can invite members or update group properties
- Members cannot remove other members or modify group settings
- Verify role assignments are correct

### Debug Mode
```bash
# Enable verbose Self SDK logging
export SELF_LOG_LEVEL=debug
go run main.go
```

## 🚀 Next Steps

After mastering group chat, explore these advanced patterns:

| Example | Complexity | Skills Gained |
|---------|------------|---------------|
| **File Sharing in Groups** 🟠 | Attachment handling, large data distribution |
| **Group Credential Exchange** 🟠 | Identity verification within groups |
| **Multi-Group Management** 🟠 | Cross-group coordination, user dashboards |
| **Group Chat UI Components** 🔴 | Mobile integration, real-time interfaces |

## 🛠️ Prerequisites

- **Go**: Version 1.22 or later
- **Self SDK**: Core functionality for group operations
- **Storage**: Write permissions for multiple client storage directories
- **Network**: Stable internet for multi-client coordination
- **Understanding**: Basic chat messaging concepts from main example

## ⚡ Performance Considerations

- **Memory**: Linear growth with number of group members
- **Storage**: Each client stores group membership and message history
- **Network**: Message encryption scales with group size
- **CPU**: Group operations require cryptographic processing per member

## 🔄 Real-World Applications

**Team Communication:**
- Development team coordination
- Project status updates
- Daily standup discussions

**Community Building:**
- Interest-based discussion groups
- Event coordination and planning
- Knowledge sharing communities

**Business Use Cases:**
- Department communications
- Cross-functional project teams  
- Customer support channels
