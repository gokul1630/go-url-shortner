package main

import (
	"crypto/rand"
	"encoding/json"
	"math/big"
	"net/http"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Data struct {
	Url string `json:"url"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})

	http.HandleFunc("/url", handleNewUrl)

	http.ListenAndServe(":3000", nil)

}

func handleNewUrl(w http.ResponseWriter, r *http.Request) {

	var decodedUrl *Data
	if r.Method == "POST" {
		json.NewDecoder(r.Body).Decode(decodedUrl)
	}

	data := Data{
		Url: generateUrl(10),
	}

	marshal, ok := json.Marshal(data)

	if ok != nil {
		panic(ok)
	}

	w.Write([]byte(marshal))
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
