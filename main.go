package main

import (
	"fmt"
	"net/http"
	"log"
	"log/slog"
	
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleHello)
	mux.HandleFunc("/ping", handleStatus)

	newRedisClient()
	
	log.Fatal(http.ListenAndServe(":8080", mux))
	
}

func handleHello(w http.ResponseWriter, _ *http.Request) {
	wc, err := w.Write([]byte("Hello world!\n"))
	if err != nil {
		slog.Error("error writing response", "err", err)
		return 
	}
	fmt.Printf("bytes written: %d\n", wc)
}

func handleStatus(w http.ResponseWriter, _ *http.Request) {
	status_code := http.StatusOK
	w.Write([]byte(string(status_code)))
	fmt.Printf("status code: %v", int(http.StatusOK))
}