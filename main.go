package main

import (
	"log"
	"net/http"
)

func main() {
	hub := newHub()
	go hub.run()

	http.HandleFunc(
		"/websocket/",
		func(w http.ResponseWriter, r *http.Request) {
			serveWs(hub, w, r)
		},
	)

	log.Fatal(http.ListenAndServe(":5050", nil))
}
