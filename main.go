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
		ok := json.NewDecoder(r.Body).Decode(decodedUrl)
		err(ok)
	}

	data := Data{
		Url: generateUrl(10),
	}

	marshal, ok := json.Marshal(data)

	err(ok)

	w.Write([]byte(marshal))
}

func generateUrl(n int) string {

	generatedString := make([]byte, n)

	for i := range generatedString {
		randomInt, ok := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))

		err(ok)

		generatedString[i] = letters[randomInt.Int64()]

	}

	return string(generatedString)
}

func err(ok any) {
	if ok != nil {
		panic(ok)
	}
}
