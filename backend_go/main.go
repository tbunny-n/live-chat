package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/charmbracelet/log"

	"github.com/tbunny-n/twitch_chat/api"
	"github.com/tbunny-n/twitch_chat/tui"
)

var addr = flag.String("addr", ":8080", "http service address")
var fullstack = flag.Bool("fullstack", false, "serve fullstack app in terminal only (default: false)")
var debug = flag.Bool("debug", false, "enable debug logging (default: false)")

func main() {
	flag.Parse()

	// ! Set up logging
	if *debug {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug logging enabled")
	}

	// ! Serve fullstack app in terminal if flag is set
	if *fullstack {
		go tui.StartChatroomInterface()
	}

	// Set up websocket hub
	hub := api.NewHub()
	go hub.Run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		api.ServeWs(hub, w, r)
	})

	// Start up server
	log.Info("Listening on", "port", *addr)
	server := &http.Server{
		Addr:              *addr,
		ReadHeaderTimeout: 3 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Debug(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "index.html")
}
