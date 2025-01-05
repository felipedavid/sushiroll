package main

import (
    "net/http"
    "log/slog"
    "fmt"
)

func main() {
    addr := ":8080"

    mux := http.NewServeMux()

    mux.HandleFunc("/", helloWorldHandler)

    server := &http.Server{
        Addr: addr,
        Handler: mux,
    }

    slog.Info("Starting web server", "addr", addr)
    err := server.ListenAndServe()
    panic(err)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world :3")
}
