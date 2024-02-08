package main

import (
    "fmt"
    "log"
    "net/http"
)

const portNum string = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Homepage")
}

func Info(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Info page")
}

func main() {
    log.Println("Starting our simple http server.")

    http.HandleFunc("/", Home)
    http.HandleFunc("/info", Info)

    log.Println("Started on port", portNum)
    fmt.Println("To close connection CTRL+C")

    err := http.ListenAndServe(portNum, nil)
    if err != nil {
        log.Fatal(err)
    }
}
