# How This Works in Your Applications

### Traditional Authentication Flow
```
1. User enters username/password
2. Server checks database
3. Server creates session token
4. User is "logged in" to that specific app
```

**Problems:** Password security, account recovery, platform lock-in

### Self SDK Authentication Flow
```
1. User creates Self identity (one time setup)
2. App requests connection to user's DID
3. User approves connection cryptographically
4. Secure encrypted channel established
```

**Benefits:** No passwords, cryptographic security, works across apps

### Real-World Example: Healthcare App

#### Traditional Approach:
```
User → Creates account with username/password
     → Uploads medical records to your servers
     → Trusts you with sensitive health data
     → Loses access if account is compromised
```

#### Self SDK Approach:
```
User → Uses existing Self identity
     → Shares specific health credentials when needed
     → Maintains control over their data
     → Can revoke access at any time
``` 
