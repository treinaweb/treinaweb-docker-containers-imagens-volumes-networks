package main

import (
	"fmt"
	"net/http"
	"os"
)

func pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/ping", pong)
	http.ListenAndServe(":"+port, nil)
}
