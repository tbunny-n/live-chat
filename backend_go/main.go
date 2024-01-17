package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func handleConnections(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", handleConnections)

	// Start server on port 8080
	log.Println("http server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
