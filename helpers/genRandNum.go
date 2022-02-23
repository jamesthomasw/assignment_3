package helpers

import (
	"crypto/rand"
	"math/big"
)

func GenRandNum(min, max int64) int64 {

	bg := big.NewInt(max - min)

	n, err := rand.Int(rand.Reader, bg)
	if err != nil {
		panic(err)
	}

	return n.Int64() + min
}