package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

var macAddr string

func startNas(w http.ResponseWriter, r *http.Request) {
	broadcastIP := os.Getenv("BROADCAST_IP")
	if broadcastIP == "" {
		broadcastIP = "255.255.255.255"
	}

	udpPort := os.Getenv("PORT")
	if udpPort == "" {
		udpPort = "9"
	}

	bcastInterface := os.Getenv("INTERFACE")

	wake(macAddr, broadcastIP, udpPort, bcastInterface)
}

func main() {
	macAddr = os.Getenv("MAC_ADDR")

	if macAddr == "" {
		fmt.Printf("error missing mac address\n")
		os.Exit(1)
	}

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("/app/static"))
	mux.Handle("/", fs)

	mux.HandleFunc("/api/start", startNas)

	err := http.ListenAndServe(":8090", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
