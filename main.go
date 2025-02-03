/*
The Main function of the prototype is defined as a Key Exchange between two individuals, namely, Alice and Bod.
This is the general structure used, and to be used for a key-exchange and will be used for the final version of KnotDH for demonstrations
*/

package main

import (
	"fmt"
	"math/big"
	"knot/knotdh"
)

func main() {
	// Defining Knots
	prime := knotdh.Knot{ID: big.NewInt(23)} // Example small prime knot
	base := knotdh.Knot{ID: big.NewInt(5)}   // Example generator knot

	// Definig Alice's DH instance
	alice := knotdh.NewInstanceKnotDH(prime, base)
	alice.GeneratePrivateKnot()
	alice.ComputePublicKnot()

	// Defining Bob's DH instance
	bob := knotdh.NewInstanceKnotDH(prime, base)
	bob.GeneratePrivateKnot()
	bob.ComputePublicKnot()

	// Compute shared secrets
	aliceShared := alice.ComputeSharedSecret(bob.PublicKnot)
	bobShared := bob.ComputeSharedSecret(alice.PublicKnot)

	// Display results
	fmt.Println("Alice's Shared Secret:", aliceShared)
	fmt.Println("Bob's Shared Secret:", bobShared)

	// Verify key agreement
	if aliceShared.Cmp(bobShared) == 0 {
		fmt.Println("Key Exchange Successful!")
	} else {
		fmt.Println("Key Exchange Failed!")
	}
}

