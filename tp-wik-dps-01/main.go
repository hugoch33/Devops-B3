package main

import (
	"log"
	"net/http"

	"tp-wik-dps-01/config"
	"tp-wik-dps-01/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", handlers.PingHandler)

	addr := ":" + config.GetPort()

	log.Printf("Server listening on %s\n", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
