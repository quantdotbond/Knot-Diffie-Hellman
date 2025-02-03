# Knot-Based Diffie-Hellman Prototype

## Overview
This repo project currently is a **prototype** for a **knot-based Diffie-Hellman (DH) key exchange**. The Key Exchange proposes that instead of using standard modular exponentiation, it utilizes **knot operations** as a semigroup and computes a **finite type invariant** (currently a placeholder for the **Kontsevich invariant**). This prototype serves as a foundation for integrating real mathematical operations related to knot theory.

## Project Structure
The project is organized into the following files:

```
myproject/
├── go.mod                     # Go module configuration
├── main.go                     # Runs the knot-based DH key exchange
├── knotdh/
│   ├── dh.go                   # Implements the DH protocol using knots
│   ├── knot_operations.go       # Knot operations
│   ├── invariant.go             # Computes finite type invariant
│   ├── utils.go                 # Utility functions (currently empty)
```

### File Descriptions
- **`main.go`** → Runs the DH protocol and verifies key agreement between two parties, Alice and Bob.
- **`knotdh/dh.go`** → Implements the key exchange logic, replacing exponentiation with knot operations.
- **`knotdh/knot_operations.go`** → Defines knot operations (e.g., ConnectedSum as a placeholder for the real group operation).
- **`knotdh/invariant.go`** → Computes a placeholder invariant (eventually to be replaced by the Kontsevich invariant).
- **`knotdh/utils.go`** → Reserved for utility functions in future expansions.

## Next Steps
### **1. Replace Placeholders with Fully Functional Knot Operations**
- Implement a **proper semigroup operation** for knots instead of using the DH exponentiation.
- Define a **structured representation of knots**

### **2. Implement the Kontsevich Invariant**
- Implement a library for the **Kontsevich integral**.
- Use it to define a **finite type invariant** to derive a unique shared secret for the shared key.

### **3. Improve Security & Performance**
- Test with larger knots and optimize the computations.
- Explore cryptographic properties of the knot-based system.

### **4. Documentation & Testing**
- Use proper standard **unit tests** for the knot operations.

