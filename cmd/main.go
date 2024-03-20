package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/oreillyross/hello-go/handlers/rest"
)

func main() {

	addr := ":8080"

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", rest.TranslateHandler)
	log.Printf("Listening on %v", addr)
	log.Fatal(http.ListenAndServe(addr, mux))

}

type Resp struct {
	Language    string `json:"languge"`
	Translation string `json:"translation"`
}
