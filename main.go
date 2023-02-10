package main

import (
	"crypto/rand"
	"math/big"
	"net/http"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})

	http.ListenAndServe(":3000", nil)

}

func generateUrl(n int) string {

	generatedString := make([]byte, n)

	for i := range generatedString {
		randomInt, ok := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))

		if ok != nil {
			panic(ok)
		}

		generatedString[i] = letters[randomInt.Int64()]

	}

	return string(generatedString)
}
