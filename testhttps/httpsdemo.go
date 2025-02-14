package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// ResponseString  常量
const ResponseString = "KFC"

func main1() {
	http.HandleFunc("/", rootHandler)
	http.ListenAndServeTLS("localhost:8080", "cert.pem", "key.pem", nil)
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", fmt.Sprint(len(ResponseString)))
	log.Println(time.Now())
	w.Write([]byte(ResponseString))
}

func main() {
	h := http.FileServer(http.Dir("."))
	http.ListenAndServeTLS(":8001", "cert.pem", "key.pem", h)
}
