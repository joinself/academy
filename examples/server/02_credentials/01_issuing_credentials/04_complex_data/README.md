# Complex Data Structure Credentials

🎯 **What you'll learn:**
- How to structure complex nested data in credentials
- Arrays and collections in claims
- Hierarchical data organization
- Real-world data modeling patterns
- Advanced claim structuring techniques

📚 **Prerequisites:** Complete `../01_basic/`, `../02_multi_claim/`, and `../03_with_evidence/` examples first.

## Overview

This example demonstrates creating credentials with complex nested data structures using the Self SDK. You'll learn to organize sophisticated organizational data with nested objects, arrays, and hierarchical structures while maintaining cryptographic integrity.

## Key Concepts

### Nested Objects
Organize related data in hierarchical structures:
```go
"position": map[string]interface{}{
    "title":      "Senior Software Engineer",
    "department": "Engineering",
    "level":      "L5",
    "manager":    "jane.smith@techcorp.com",
}
```

### Arrays and Collections
Store multiple values of the same type:
```go
"permissions": []string{
    "read:repositories",
    "write:code", 
    "deploy:staging",
    "review:pull-requests",
}
```

### Complex Arrays
Arrays containing objects with multiple properties:
```go
"certifications": []map[string]interface{}{
    {
        "name":       "AWS Solutions Architect",
        "level":      "Professional", 
        "issueDate":  "2023-06-15",
        "verified":   true,
    },
}
```

## Code Breakdown

### 1. Nested Object Structure
```go
"contact": map[string]interface{}{
    "email":    "john.doe@techcorp.com",
    "phone":    "+1-555-0123",
    "office":   "Building A, Floor 3, Desk 42",
    "timezone": "America/New_York",
    "address": map[string]interface{}{  // Deeply nested
        "street":  "123 Tech Street",
        "city":    "San Francisco",
        "state":   "CA",
        "zipCode": "94105",
        "country": "United States",
    },
}
```

**What this demonstrates:**
- **Hierarchical organization**: Contact info contains nested address
- **Logical grouping**: Related contact data stays together
- **Scalable structure**: Easy to add more contact details
- **Data integrity**: All nested data is cryptographically signed

### 2. Benefits Package Modeling
```go
"benefits": map[string]interface{}{
    "healthInsurance": true,
    "retirement401k":  true,
    "paidTimeOff":     25,
    "stockOptions":    1000,
    "remoteWork":      true,
    "wellness": map[string]interface{}{  // Nested benefits category
        "gymMembership":    true,
        "mentalHealth":     true,
        "annualWellness":   "$1000",
        "flexibleSchedule": true,
    },
}
```

**Key patterns:**
- **Mixed data types**: Booleans, numbers, strings in same structure
- **Categorical nesting**: Wellness benefits grouped separately
- **Extensible design**: Easy to add new benefit categories
- **Clear semantics**: Self-documenting structure

### 3. Certification Arrays
```go
"certifications": []map[string]interface{}{
    {
        "name":       "AWS Solutions Architect",
        "level":      "Professional",
        "issueDate":  "2023-06-15",
        "expiryDate": "2026-06-15",
        "verified":   true,
        "provider":   "Amazon Web Services",
    },
    {
        "name":       "Kubernetes Administrator",
        "level":      "Certified",
        "issueDate":  "2023-09-20",
        "expiryDate": "2026-09-20",
        "verified":   true,
        "provider":   "Cloud Native Computing Foundation",
    },
}
```

**Benefits of this approach:**
- **Standardized structure**: Each certification has consistent fields
- **Queryable data**: Easy to filter by provider, level, or expiry
- **Audit trail**: Issue and expiry dates for compliance
- **Verification status**: Boolean flag for validation state

### 4. Project History with Technology Arrays
```go
"projects": []map[string]interface{}{
    {
        "name":         "Payment Gateway Redesign",
        "role":         "Lead Developer",
        "startDate":    "2023-01-01",
        "endDate":      "2023-06-30",
        "status":       "Completed",
        "technologies": []string{"Go", "PostgreSQL", "Redis", "Docker"},
    },
}
```

**Advanced patterns:**
- **Arrays within objects**: Technologies list for each project
- **Status tracking**: Project lifecycle management
- **Role documentation**: Responsibility and authority levels
- **Timeline preservation**: Complete project history

## Running the Example

```bash
cd examples/go/02_credentials/01_issuing_credentials/04_complex_data
go run main.go
```

## Expected Output

```
Complex Credential Issuance Demo
================================
Issuer: 1:ABC123...
Holder: 1:DEF456...

Creating organization credential with complex data...
✅ Organization credential issued (Type: [VerifiableCredential Organisation])
   Employee: EMP-2024-001 - Senior Software Engineer
   Data structure: 5 permissions, 2 certifications, 2 projects
✅ Demo completed successfully!
```

## What Just Happened

1. **Created complex nested structure**: Multi-level hierarchical data organization
2. **Used arrays for collections**: Multiple permissions, certifications, and projects
3. **Nested objects for categories**: Position, contact, benefits, and address groupings
4. **Maintained data integrity**: All nested data cryptographically signed together
5. **Demonstrated real-world modeling**: Enterprise employee credential structure

## Data Structure Patterns

### 1. Hierarchical Organization
```go
// ✅ Good: Logical hierarchy
"employee": map[string]interface{}{
    "personal": map[string]interface{}{
        "name": "John Doe",
        "contact": map[string]interface{}{
            "email": "john@company.com",
        },
    },
    "employment": map[string]interface{}{
        "position": "Engineer",
        "department": "Tech",
    },
}
```

### 2. Collection Management
```go
// ✅ Good: Consistent object structure in arrays
"skills": []map[string]interface{}{
    {"name": "Go", "level": "Expert", "years": 5},
    {"name": "Python", "level": "Advanced", "years": 3},
}

// ❌ Avoid: Inconsistent structures
"skills": []interface{}{
    "Go Programming",           // String
    map[string]interface{}{     // Object
        "name": "Python",
        "level": "Advanced",
    },
}
```

### 3. Type Consistency
```go
// ✅ Good: Consistent types within categories
"scores": map[string]interface{}{
    "technical":     85,    // All numbers
    "communication": 92,
    "leadership":    78,
}

// ✅ Good: Logical type usage
"flags": map[string]interface{}{
    "isActive":     true,   // All booleans
    "isVerified":   false,
    "hasAccess":    true,
}
```

## Advanced Patterns

### Dynamic Array Building
```go
// Build permissions based on role
var permissions []string
if role == "senior" {
    permissions = append(permissions, "admin:team-resources")
}
permissions = append(permissions, "read:repositories", "write:code")

claims["permissions"] = permissions
```

### Conditional Nested Objects
```go
// Add manager info only for non-senior roles
position := map[string]interface{}{
    "title": "Software Engineer",
    "level": "L4",
}

if level != "executive" {
    position["manager"] = "jane.smith@company.com"
}

claims["position"] = position
```

### Array Validation Patterns
```go
// Validate certification structure before adding
certifications := []map[string]interface{}{}

for _, cert := range userCertifications {
    if cert.IsValid() && cert.IsVerified() {
        certifications = append(certifications, map[string]interface{}{
            "name":     cert.Name,
            "verified": true,
            "issuer":   cert.Provider,
        })
    }
}

claims["certifications"] = certifications
```

## Best Practices

### 1. Depth Management
```go
// ✅ Good: Reasonable nesting depth (2-3 levels)
"employee": {
    "contact": {
        "address": {
            "street": "123 Main St"
        }
    }
}

// ❌ Avoid: Excessive nesting (4+ levels)
"data": {
    "employee": {
        "personal": {
            "contact": {
                "address": {
                    "details": {
                        "street": "123 Main St"
                    }
                }
            }
        }
    }
}
```

### 2. Array Uniformity
```go
// ✅ Good: Consistent object structure
"projects": []map[string]interface{}{
    {"name": "A", "status": "Complete", "budget": 1000},
    {"name": "B", "status": "Active", "budget": 2000},
}

// ❌ Avoid: Mixed structures
"projects": []interface{}{
    map[string]interface{}{"name": "A", "status": "Complete"},
    map[string]interface{}{"title": "B", "state": "Active"}, // Different keys
}
```

### 3. Null Value Handling
```go
// ✅ Good: Omit missing values or use explicit nulls
claims := map[string]interface{}{
    "name": "John Doe",
    // Don't include "middleName" if not available
}

// Or use explicit nulls when semantically important
claims["middleName"] = nil  // Explicitly no middle name
```

## Performance Considerations

### Size Optimization
- **Avoid redundancy**: Don't repeat data across nested structures
- **Use references**: Link to external data when appropriate
- **Compress arrays**: Consider pagination for large collections

### Memory Efficiency
```go
// ✅ Efficient: Build structure incrementally
claims := make(map[string]interface{})
claims["basic"] = basicInfo
claims["contact"] = buildContactInfo()
claims["projects"] = buildProjectList()

// ❌ Inefficient: Large inline structures
claims := map[string]interface{}{
    // ... hundreds of lines of nested data
}
```

## Troubleshooting

### Common Issues

**"Failed to build credential"**
- Check for circular references in nested structures
- Ensure all values are JSON-serializable types
- Verify map keys are strings

**"Invalid data type in claims"**
- Use `interface{}` for mixed-type collections
- Ensure consistent types within logical groups
- Avoid complex custom types

**"Credential too large"**
- Consider breaking into multiple credentials
- Use evidence references instead of embedding large data
- Implement pagination for large arrays

### Debugging Complex Structures
```go
// Debug nested structure
import "encoding/json"

claimsJSON, _ := json.MarshalIndent(claims, "", "  ")
fmt.Printf("Claims structure:\n%s\n", claimsJSON)
```

## Production Patterns

### Schema Validation
```go
// Define expected structure
type EmployeeCredential struct {
    OrganizationName string `json:"organizationName"`
    EmployeeID       string `json:"employeeId"`
    Position         struct {
        Title      string `json:"title"`
        Department string `json:"department"`
        Level      string `json:"level"`
    } `json:"position"`
    Permissions []string `json:"permissions"`
}

// Validate before credential creation
func validateEmployeeData(claims map[string]interface{}) error {
    // Implementation for validation
    return nil
}
```

### Modular Construction
```go
func buildEmployeeCredential(employee Employee) map[string]interface{} {
    claims := make(map[string]interface{})
    
    claims["organizationName"] = employee.Company
    claims["employeeId"] = employee.ID
    claims["position"] = buildPositionInfo(employee)
    claims["permissions"] = buildPermissions(employee.Role)
    claims["contact"] = buildContactInfo(employee)
    
    return claims
}
```

## Next Steps

📚 **Continue learning:**
- `../05_comprehensive/` - All features combined
- `../../02_exchanging_credentials/` - Using complex credentials in presentations

🔍 **Related concepts:**
- JSON Schema validation for credential structures
- Query patterns for nested credential data
- Performance optimization for large structures

## When to Use Complex Data Structures

### Perfect for:
- **Employee records**: Complete organizational context
- **Academic transcripts**: Multiple courses, grades, and achievements
- **Medical records**: Patient history, treatments, and diagnoses  
- **Financial portfolios**: Multiple accounts, transactions, and holdings

### Consider simpler alternatives for:
- **Single facts**: Use basic credentials for simple assertions
- **Frequently changing data**: Real-time info better suited for APIs
- **Large datasets**: Consider external storage with credential references
- **Different security levels**: Separate credentials for different access needs 
