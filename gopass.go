package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func generateTokens() (string, string, string) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	tokens := [3]string{}

	for i := range tokens {
		token := make([]byte, 6)
		for j := range token {
			r, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
			if err != nil {
				fmt.Print(err)
			}

			token[j] = letters[r.Uint64()]
		}

		tokens[i] = string(token)
	}

	return tokens[0], tokens[1], tokens[2]
}

func generate() string {
	f, s, t := generateTokens()

	return fmt.Sprintf("%s-%s-%s", f, s, t)
}

func main() {
	fmt.Println(generate())
}
