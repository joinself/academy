# Controlling Credential Disclosure

> **🛠️ Practical Implementation:** [Credential Examples](../examples/credentials.md#phase-25-master-presentation-request-patterns)  
> **📖 Learn the concepts first:** [Verifiable Credentials Concepts](verifiable-credentials.md)

## What You'll Learn

- How to control when a credential should be shared using presentation requests.
- The difference between unconditional and conditional credential retrieval.
- How to use operators to create privacy-preserving conditional requests.
- **Working code** to implement both request modes in your applications.
- Real-world decision-making for designing credential requests.

**Time Investment**: 10 minutes to understand disclosure control.  
**Immediate Result**: A clear strategy for designing credential requests.

---

## The Core Concept: Conditional Disclosure

When you request a credential, you don't just ask for it by type; you can also specify **conditions** under which the user's device should share it. The user's device evaluates these conditions locally. This gives you two powerful modes of operation for every request:

1.  **Unconditional Retrieval**: "Please share this credential."
2.  **Conditional Retrieval**: "Please share this credential *only if* it meets these criteria."

In both cases, a successful response contains the **full credential data**. The difference is the logic that determines whether a response is sent at all.

---

## Two Modes of Credential Retrieval

### **Mode 1: Unconditional Retrieval** (`nil` parameter)

This is the simplest mode. You request a credential by type, and the user's device will always return it if the user consents. This is the most common approach when you need guaranteed data access for your application to function.

**Behavior:**
```
Verifier → "Please share your email credential" → Holder
Holder → Returns complete credential with all claims → Verifier
Result: Verifier always receives the full credential data (if user has it and approves).
```

**Implementation:**
```go
import "github.com/joinself/self-go-sdk/credential"

// Request complete credential data, no conditions.
content, err := message.NewCredentialPresentationRequest().
    Type(credential.PresentationTypeContactDetails).
    Details(credential.CredentialTypeEmail, nil).  // ← nil means no conditions.
    Finish()
```

### **Mode 2: Conditional Retrieval** (with operator parameters)

This mode adds a privacy-preserving check. You specify conditions that the credential must meet. The user's device evaluates these conditions locally *before* sharing. This is ideal for eligibility checks where you only want data from users who meet certain criteria.

**Behavior:**
```
Verifier → "Share credential X only if condition Y is met" → Holder
Holder → Evaluates condition locally → Returns full credential OR nothing.
Result: Verifier gets complete credential data only if the condition is satisfied.
```

**Implementation:**
```go
// Request credential only if user is 18 or older.
eighteenYearsAgo := time.Now().AddDate(-18, 0, 0).Format("2006-01-02")

content, err := message.NewCredentialPresentationRequest().
    Type(credential.PresentationTypePassport).
    Details(
        []string{"DateOfBirthCredential"}, // This is a custom credential type
        []*message.CredentialPresentationDetailParameter{
            message.NewCredentialPresentationDetailParameter(
                message.OperatorLowerThan, // This is the condition.
                "dateOfBirth",
                eighteenYearsAgo,
            ),
        },
    ).
    Finish()
```

---

## Mode Comparison

The two modes offer different trade-offs between guaranteed data access and user privacy.

| Aspect | Unconditional Retrieval | Conditional Retrieval |
|--------|-------------------------|-----------------------|
| **Data Exposure** | Full credential content | Full content, but only if condition is met and user approves |
| **User Privacy** | Standard | Enhanced (no data shared if ineligible or user declines) |
| **Response Behavior** | Always prompts user for approval (if they have the credential) | Prompts user **only if** they are eligible |
| **Verifier Logic** | Simple (always expect a response, approved or denied) | Must handle timeouts (when user is ineligible or declines) |

---

## Technical Foundations & Best Practices

### The User is Always in Control

Before diving into the technical details, it's critical to understand one core principle: **the user must always explicitly approve any credential sharing.**

The "automation" in conditional retrieval is only in the **pre-filtering** of requests. The user's device automatically filters out requests where the user is ineligible, saving them from seeing irrelevant prompts. However, for any eligible request, the user is prompted and must consent before any data is sent.

**The flow is always:**

1.  **Request Sent**: Verifier sends a conditional or unconditional request.
2.  **Device Evaluates (if conditional)**: The Self app checks if the conditions are met.
3.  **User Prompted (if eligible)**: The user sees a prompt like, "Site X is asking for your Passport. Allow?"
4.  **User Approves**: The user must tap "Approve" or "Share".
5.  **Data Sent**: Only after user approval is the credential shared.

This user-centric design ensures privacy and control are always maintained.

### How Conditional Retrieval Works

1.  **Local Evaluation**: The user's device receives the request and sees the conditions (e.g., `dateOfBirth` must be lower than `eighteenYearsAgo`). It performs this check against the credential data stored locally on the device.
2.  **Decision Making**: Based on the evaluation, the device decides whether ask the user for confirmation or not.
3.  **Conditional Response**: The full credential is sent, but *only if* the condition was satisfied and the user accepts to share it.
4.  **Verifier Handling**: The verifier must be prepared for two outcomes: receiving the full credential (after user approval) or receiving no response within a timeout period (which implies the condition was not met or the user declined).

**Key Behavioral Properties:**

- **Eligibility Gating**: Only eligible users are prompted to share their credentials.
- **Binary Outcome**: The result is either the full credential data (with consent) or no response at all.
- **Privacy Protection**: Ineligible users do not expose any data, not even that they failed a check.
- **User-Centric Control**: The user makes the final sharing decision for all eligible requests.

### The Relationship Between `Type` and `Details`

It's important to understand how `Type()` and `Details()` work together. They are separate but powerful when combined.

1.  **`Type(...)`**: This sets the high-level **intent** or **purpose** of the presentation. You are telling the recipient, "This is a passport presentation." It acts as a label for the "box" of credentials you are requesting.
2.  **`Details(...)`**: This defines the **rules** for what must go *inside* the box. It specifies which credentials (e.g., `PassportCredential`) and which claims (e.g., `dateOfBirth`) should be checked.

The `PresentationType` itself does not contain the details, but it provides the semantic wrapper for a presentation that is built according to the rules in your `Details` call.

#### Example: Requesting an 18+ Passport Presentation

Here’s how you combine them to request a presentation labeled as a "Passport Presentation" that *only* contains a passport credential if the holder is over 18.

```go
import (
    "time"
    "github.com/joinself/self-go-sdk/credential"
    "github.com/joinself/self-go-sdk/message"
)

// Create the presentation request.
content, err := message.NewCredentialPresentationRequest().
    // Set the INTENT: "I'm asking for a passport presentation."
    Type(credential.PresentationTypePassport).

    // Set the RULES for the content:
    // "Look for a 'PassportCredential' and only include it if the 'dateOfBirth'
    // claim shows the person is older than 18."
    Details(
        credential.CredentialTypePassport, // The credential type to look for.
        []*message.CredentialPresentationDetailParameter{
            message.NewCredentialPresentationDetailParameter(
                message.OperatorLowerThan, // The condition.
                "dateOfBirth",             // The claim to check.
                eighteenYearsAgo,          // The value to compare against.
            ),
        },
    ).
    Finish()
```

### Predefined Types for Interoperability

To promote consistency, the SDK provides predefined constants for both presentation and credential types. You should use these whenever possible.

#### Presentation Types

<div data-github-embed="https://github.com/joinself/self-go-sdk/blob/main/credential/presentation.go#L18-L24"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

#### Credential Types

<div data-github-embed="https://github.com/joinself/self-go-sdk/blob/main/credential/credential.go#L22-L32"
     data-style="github-dark-dimmed"
     data-show-border="true"
     data-show-line-numbers="true"
     data-show-file-meta="true"
     data-show-full-path="true"
     data-show-copy="true"></div>

> **Note:** The base type `VerifiablePresentation` is always included in presentations. The second type provides specific semantic context for the presentation's purpose. You are in control and can always extend this list with your own custom types, for example `[]string{"VerifiablePresentation", "MyCustomLoanApplicationPresentation"}`.

### Available Operators for Conditions

```go
// Comparison operators for numerical/date fields.
message.OperatorEquals        // field == value
message.OperatorNotEquals     // field != value  
message.OperatorGreaterThan   // field > value
message.OperatorLowerThan     // field < value

// Existence operators.
message.OperatorNotEquals     // Check if a field exists by using an empty value.
```

---

## Summary

You've learned the fundamental mechanism for controlling credential disclosure:

### Unconditional Retrieval

- **When**: You always need the data for your application to function.
- **Privacy**: Standard data sharing model.
- **Result**: You always get the full credential if the user has it.

### Conditional Retrieval

- **When**: You only want data from users who meet specific criteria.
- **Privacy**: Enhanced privacy, as ineligible users share nothing.
- **Result**: You get the full credential, but only if the condition is met.

> **The mode you choose determines:**
>
> - User privacy and what data you process.
> - Your application's logic for handling responses vs. timeouts.
> - Your compliance with data minimization principles.

---

## Next Steps

**Ready to implement these requests?** Continue with:

- **[Credential Examples](../examples/credentials.md#phase-25-master-presentation-request-patterns)** - Implement both modes in working code.
- **[Age Verification Solution](../../examples/solutions/age-verifier/)** - See the conditional pattern in a production-ready example.
- **[Email Verification Example](../../examples/server/02_credentials/02_exchanging_credentials/email_verification/)** - See the unconditional pattern in action.

---

> **Congratulations!** You now understand how to precisely control credential disclosure. Choose unconditional retrieval when you always need data, and conditional retrieval when you need to protect the privacy of ineligible users. Both approaches leverage the same powerful, cryptographically secure foundation. 
