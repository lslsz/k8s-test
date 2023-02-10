package main

import (
    "fmt"
    "net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello v1")
}

func main() {
    http.HandleFunc("/", HelloHandler)
    http.ListenAndServe(":8080", nil)
}
