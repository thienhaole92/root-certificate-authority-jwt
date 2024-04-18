package util

import (
	"crypto/rand"
	"math/big"
)

func GenerateSerial() (*big.Int, error) {
	limit := new(big.Int).Lsh(big.NewInt(1), 512)
	serial, err := rand.Int(rand.Reader, limit)
	return serial, err
}
