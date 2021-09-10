package crng

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math/big"
)

func init() {
	if _, err := rand.Read(make([]byte, 32)); err != nil {
		panic(fmt.Errorf("crng: cannot initialize: %v", err))
	}
}

// Secret returns a random base64url encoded sequence of size bytes.
func Secret(size int) string {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		panic(fmt.Errorf("crng: cannot read crng: %v", err))
	}
	return base64.URLEncoding.EncodeToString(b)
}

// Int returns a random value in range [0, max).
func Int(max int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		panic(fmt.Errorf("crng: cannot read crng: %v", err))
	}
	return int(n.Int64())
}

// Incl returns a random value in range [0, max].
func Incl(max int) int {
	return Int(max + 1)
}

// Range returns a random value in range [from, to].
func Range(from, to int) int {
	if from > to {
		return 0
	}
	return from + Incl(to-from)
}

// Chance returns true with probability of percent, otherwise false.
func Chance(percent int) bool {
	r := Range(1, 100)
	if percent >= r {
		return true
	}
	return false
}

// CoinFlip returns true with probability of 50%, otherwise false.
func CoinFlip() bool {
	return Int(2) == 0
}
