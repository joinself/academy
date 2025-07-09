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

<h1>Welcome to the Joinself Academy</h1>
<p class="subtitle">The definitive guide for developers building the future of digital identity.</p>

<div class="homepage-cta">
    <a href="examples/setup/" class="md-button md-button--primary">Get Started with Your First Self ID</a>
</div>

<div class="landing-grid">
    <div class="card">
        <h3><a href="concepts/overview/">Understand the Fundamentals</a></h3>
        <ul>
            <li><span class="icon">📄</span> <a href="concepts/decentralized-identity/">Decentralized Identity</a></li>
            <li><span class="icon">🔐</span> <a href="concepts/cryptographic-foundations/">Cryptographic Foundations</a></li>
            <li><span class="icon">🤝</span> <a href="concepts/secure-connections/">Secure Connections</a></li>
            <li><span class="icon">🎫</span> <a href="concepts/verifiable-credentials/">Verifiable Credentials</a></li>
            <li><span class="icon">🤫</span> <a href="concepts/message-layer-security/">Message Layer Security</a></li>
        </ul>
    </div>
    <div class="card">
        <h3><a href="examples/overview/">Learn by Building</a></h3>
        <ul>
            <li><span class="icon">🛠️</span> <a href="examples/setup/">Setup & Configuration</a></li>
            <li><span class="icon">🔗</span> <a href="examples/connections/">Connections & Discovery</a></li>
            <li><span class="icon">🎫</span> <a href="examples/credentials/">Verifiable Credentials</a></li>
            <li><span class="icon">💬</span> <a href="examples/chat/">Chat & Messaging</a></li>
            <li><span class="icon">⚙️</span> <a href="examples/advanced/">Advanced Features</a></li>
        </ul>
    </div>
    <div class="card">
        <h3><a href="solutions/overview/">Explore Real-World Solutions</a></h3>
        <ul>
            <li><span class="icon">🔑</span> <a href="solutions/authentication/">Authentication</a></li>
            <li><span class="icon">✍️</span> <a href="solutions/digital-signatures/">Digital Signatures</a></li>
            <li><span class="icon">✅</span> <a href="solutions/identity-verification/">Identity Verification</a></li>
        </ul>
    </div>
</div>

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
