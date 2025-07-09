---
hide:
  - navigation
  - toc
---

<div class="custom-homepage">

<style>
  .custom-homepage h1, .custom-homepage .subtitle {
    text-align: center;
  }
  .custom-homepage .subtitle {
    font-size: 1.2rem;
    font-style: italic;
    color: var(--md-default-fg-color--light);
  }
  .custom-homepage .card ul {
    padding-left: 0;
  }
  .custom-homepage .card li a {
    font-weight: 500;
  }
  .custom-homepage .final-cta {
    text-align: center;
    margin-top: 3rem;
  }
</style>

# Welcome to the Self Academy

The Self Academy is your comprehensive resource for learning how to build next-generation applications with decentralized identity. Whether you're new to the concepts or an experienced developer, you'll find everything you need to get started and build production-ready systems.

## 🚀 Get Started with Solutions

Our new solution guides are the best place to start. They provide a direct path from a business use case to the code you need to write.

- [**Authentication**](./solutions/authentication.md): Build passwordless login flows.
- [**Identity Verification**](./solutions/identity-verification.md): Verify user identities with cryptographic proof.
- [**Digital Signatures**](./solutions/digital-signatures.md): Create legally binding, biometric-backed signatures.

## 📚 Core Learning Paths

- **[Concepts](./concepts/overview.md)**: Start here to understand the core theories and principles behind decentralized identity, verifiable credentials, and secure connections.

- **[Examples](./examples/overview.md)**: Dive into runnable code examples that demonstrate specific features and use cases for our server and mobile SDKs.


<div class="try-it-out">
    <h2><span class="icon">⚡</span> Try It Out</h2>
    <p>Get a feel for the Self SDK right here. The code below shows how to create a brand new, cryptographically secure Self identity with just a few lines of Go.</p>
    <p><strong>Run this full example yourself:</strong> <a href="../examples/server/00_setup/01_new_account/go/main.go"><code>examples/server/00_setup/01_new_account/go/main.go</code></a></p>
    <div class="code-wrapper">
        <div class="code-header">Go</div>
        <div class="highlight">
<pre><code class="language-go">// From: examples/server/00_setup/01_new_account/go/main.go

// Create a new Self SDK account with a randomly generated key and default storage.
selfAccount, err := account.New(&account.Config{
    StorageKey:  generateStorageKey(),
    StoragePath: "./storage",
    Environment: account.TargetSandbox,
})
if err != nil {
    log.Fatal("failed to create new account: ", err)
}
defer selfAccount.Close()

// Open the account's inbox to get its unique address.
inbox, err := selfAccount.InboxOpen()
if err != nil {
    log.Fatal("failed to open inbox: ", err)
}

fmt.Println("✅ New account created successfully!")
fmt.Printf("🆔 Account DID: %s\n", inbox.String())
</code></pre>
        </div>
    </div>
</div>

<div class="final-cta">
    <h3>🌟 Ready to build the future of digital identity?</h3>
    <p><a href="examples/setup/"><strong>Start your journey here.</strong></a></p>
</div>

</div>
