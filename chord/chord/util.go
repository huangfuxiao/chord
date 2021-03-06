/*
 *  Brown University, CS138, Spring 2017
 *
 *  Purpose: Utility functions to help with dealing with ID hashes in Chord.
 */

package chord

import (
	"bytes"
	"crypto/sha1"
	"math/big"
)

// Hash a string to its appropriate size.
func HashKey(key string) []byte {
	h := sha1.New()
	h.Write([]byte(key))
	v := h.Sum(nil)
	return v[:KEY_LENGTH/8]
}

// Convert a []byte to a big.Int string (useful for debugging/logging)
func HashStr(keyHash []byte) string {
	keyInt := big.Int{}
	keyInt.SetBytes(keyHash)
	return keyInt.String()
}

func EqualIds(a, b []byte) bool {
	return bytes.Equal(a, b)
}

// Example of how to do math operations on []byte IDs.
// You may not need this function.
func AddIds(a, b []byte) []byte {
	aInt := big.Int{}
	aInt.SetBytes(a)

	bInt := big.Int{}
	bInt.SetBytes(b)

	sum := big.Int{}
	sum.Add(&aInt, &bInt)
	return sum.Bytes()
}

/*
 * On this crude ascii Chord ring, X is between (A : B)
 *
 *    ___
 *   /   \-A
 *  |     |
 * B-\   /-X
 *    ---
 */
func Between(nodeX, nodeA, nodeB []byte) bool {

	if EqualIds(nodeA, nodeB) {
		return true
	}
	xInt := big.Int{}
	xInt.SetBytes(nodeX)

	aInt := big.Int{}
	aInt.SetBytes(nodeA)

	bInt := big.Int{}
	bInt.SetBytes(nodeB)
	if aInt.Cmp(&bInt) == -1 {
		if aInt.Cmp(&xInt) == -1 && xInt.Cmp(&bInt) == -1 {
			return true
		}
	} else {
		if aInt.Cmp(&xInt) == -1 || xInt.Cmp(&bInt) == -1 {
			return true
		}
	}

	return false
}

// Returns true if X is between (A : B]
func BetweenRightIncl(nodeX, nodeA, nodeB []byte) bool {

	return Between(nodeX, nodeA, nodeB) || EqualIds(nodeX, nodeB)
}
