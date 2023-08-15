package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, this is a basic HTTP server in Go!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server listening on port 8080...")
    http.ListenAndServe(":8080", nil)
}
