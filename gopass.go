package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func generateToken(length int) []byte {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	token := make([]byte, length)
	for j := range token {
		r, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			fmt.Print(err)
		}
		token[j] = letters[r.Uint64()]
	}

	return token
}

func generateTokens(n int, tokenLength int) []string {
	tokens := make([]string, n)
	for i := range tokens {
		tokens[i] = string(generateToken(tokenLength))
	}

	return tokens
}

func main() {
	tokens := generateTokens(3, 6)
	fmt.Printf("%s-%s-%s", tokens[0], tokens[1], tokens[2])
}
