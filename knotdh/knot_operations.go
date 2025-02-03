/*
This file will give light towards a library implementation of Knot operations, while the team is working on it, we are using some placeholders until we arrive at a final version.
*/

package knotdh

import (
	"math/big"
	"math/rand"
	"time"
)

// Knot represents a mathematical knot structure with a semigroup-like operation.
type Knot struct {
	ID *big.Int // Unique identifier for the knot
}

// Creates a new random knot (private-key equivalent).
func GenerateRandomKnot(prime Knot) Knot {
	rand.Seed(time.Now().UnixNano())
	return Knot{ID: new(big.Int).Rand(rand.New(rand.NewSource(time.Now().UnixNano())), prime.ID)}
}

// Simulates the group operation in DH. Meanwhile, it uses the standard base^exp mod prime DH
func KnotOperation(base, exp, prime Knot) Knot {
	result := new(big.Int).Exp(base.ID, exp.ID, prime.ID)
	return Knot{ID: result}
}

// Equal checks if two knots are the same.
func (k1 Knot) Equal(k2 Knot) bool {
	return k1.ID.Cmp(k2.ID) == 0
}

