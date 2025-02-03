/*
Core package template for the DH protocol, using Knot operations and invariantes as demonstrations, the final version will come out of this file.
For the time being, we are using some placeholders while we implement the required libraries and code.
*/

package knotdh

import (
	"fmt"
	"math/big"
)

// Diffie-Hellman-like key exchange using knots structs.
type InstanceKnotDH struct {
	PrivateKnot Knot
	PublicKnot  Knot
	PrimeKnot   Knot
	BaseKnot    Knot
}

// Initializes a new key exchange instance, given prime and base knots.
func NewInstanceKnotDH(prime, base Knot) *InstanceKnotDH {
	return &InstanceKnotDH{
		PrimeKnot: prime,
		BaseKnot:  base,
	}
}

// Selects a random private knot.
func (kdh *InstanceKnotDH) GeneratePrivateKnot() {
	kdh.PrivateKnot = GenerateRandomKnot(kdh.PrimeKnot)
}

// Computes the public-key using knot operations.
func (kdh *InstanceKnotDH) ComputePublicKnot() {
	kdh.PublicKnot = KnotOperation(kdh.BaseKnot, kdh.PrivateKnot, kdh.PrimeKnot)
}

// Shared secret using the Kontsevich invariant.
func (kdh *InstanceKnotDH) ComputeSharedSecret(otherPublic Knot) *big.Int {
	sharedKnot := KnotOperation(otherPublic, kdh.PrivateKnot, kdh.PrimeKnot)
	return ComputeKontsevichInvariant(sharedKnot)
}

// For debugging
func (kdh *InstanceKnotDH) DebugInfo() {
	fmt.Println("Private Knot:", kdh.PrivateKnot.ID)
	fmt.Println("Public Knot:", kdh.PublicKnot.ID)
	fmt.Println("Prime Knot:", kdh.PrimeKnot.ID)
	fmt.Println("Base Knot:", kdh.BaseKnot.ID)
}

