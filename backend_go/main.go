package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/charmbracelet/log"

	"github.com/tbunny-n/twitch_chat/api"
	"github.com/tbunny-n/twitch_chat/tui"
)

var addr = flag.String("addr", ":8080", "http service address")

// TODO: Default these to false
var Fullstack = flag.Bool("fullstack", true, "serve fullstack app in terminal only (default: false)")
var Debug = flag.Bool("debug", false, "enable debug logging (default: false)")

func main() {
	flag.Parse()
	wsUrl := fmt.Sprintf("ws://localhost%s/ws", *addr) // WebSocket URL

	// ! Set up logging
	if *Debug {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug logging enabled")
	}

	// ! Serve fullstack app in terminal if flag is set
	if *Fullstack {
		go tui.StartChatroomInterface(wsUrl)

		if !*Debug {
			log.SetLevel(log.ErrorLevel)
		}
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
