package utils

import (
	"crypto/rand"
	"math/big"
)

// Crypto/rand paketinden verilen parametreye göre random sayı üreten fonksiyon

func CreateRandomNumberGenerator(bit int) *big.Int {
	RandomCrypto, _ := rand.Prime(rand.Reader, bit)
	return RandomCrypto
}
