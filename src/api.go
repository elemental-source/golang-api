package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

func index(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "Bem vindo, %s!", request.URL.Path[1:])
}

func sobre(responseWriter http.ResponseWriter, request *http.Request) {

	m := Message{"Bem vindo programador!"}
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	responseWriter.Write(b)
}

type Message struct {
	Text string
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/sobre", sobre)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
