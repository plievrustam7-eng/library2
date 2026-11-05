package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type appStatus struct {
	State   string `json:"state"`
	Service string `json:"service_name"`
}

type metadata struct {
	Release   string `json:"release_version"`
	Timestamp string `json:"timestamp"`
}

func statusController(res http.ResponseWriter, req *http.Request) {
	payload := appStatus{
		State:   "operational",
		Service: "Library Core API",
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(res).Encode(payload)
}

func infoController(res http.ResponseWriter, req *http.Request) {
	payload := metadata{
		Release:   "2.0.1-beta",
		Timestamp: time.Now().Format(time.RFC1123),
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(res).Encode(payload)
}

func fallbackController(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
	res.WriteHeader(http.StatusNotFound)
	_, _ = res.Write([]byte("Error 404: Resource endpoint not found"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/status", statusController)
	mux.HandleFunc("/api/v1/about", infoController)
	mux.HandleFunc("/", fallbackController)

	port := ":8080"
	fmt.Fprintf(os.Stdout, "API запущен на порту %s\n", port)

	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Fprintf(os.Stderr, "Не удалось запустить веб-сервер: %v\n", err)
		os.Exit(1)
	}
}
