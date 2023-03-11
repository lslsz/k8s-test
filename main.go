package main

import (
    "fmt"
    "net/http"
    "os"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    hostname,_ := os.Hostname()
    fmt.Fprintf(w, "Hello %s\n", hostname)
}

func main() {
    http.HandleFunc("/", HelloHandler)
    http.ListenAndServe(":8080", nil)
}

