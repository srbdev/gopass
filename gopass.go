package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func generateToken() []byte {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	token := make([]byte, 6)

	for j := range token {
		r, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			fmt.Print(err)
		}
		token[j] = letters[r.Uint64()]
	}

	return token
}

func generateTokens() (string, string, string) {
	tokens := [3]string{}
	for i := range tokens {
		tokens[i] = string(generateToken())
	}

	return tokens[0], tokens[1], tokens[2]
}

func generate() string {
	first, second, third := generateTokens()
	return fmt.Sprintf("%s-%s-%s", first, second, third)
}

func main() {
	fmt.Println(generate())
}
