package main

import (
	"log"
	"net/http"
)

func handleSocket(w http.ResponseWriter, r *http.Request) {
	//serveWs(w, r)
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello World!"))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main () {
	http.Handle(
		"/websocket",
		http.HandlerFunc(handleSocket),
	)
	//http.HandleFunc("/websocket", handleSocket)

	log.Fatal(http.ListenAndServe(":5050", nil))
}