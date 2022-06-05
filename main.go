package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const headerToken = "X-Token"
const headerVersion = "X-Version"

func main() {
	http.HandleFunc("/", rootHandler)
	error := http.ListenAndServe(":80", nil)
	if error != nil {
		log.Fatal(error)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	var remoteAddr string = r.RemoteAddr
	var xToken string = r.Header.Get(headerToken)
	w.Header().Add(headerToken, xToken)
	w.Header().Add(headerVersion, os.Getenv("VERSION"))

	temp := strings.Split(r.RequestURI, "?")
	var uri string = temp[0]
	var method string = r.Method

	var statusCode int = http.StatusOK
	var text string = "OK"
	switch method {
	case http.MethodGet:
		switch uri {
		case "/healthz":
			statusCode = http.StatusOK
			text = "OK"
		default:
			statusCode = http.StatusBadRequest
			text = "What?!"
		}
	default:
		statusCode = http.StatusMethodNotAllowed
		text = "Oops!"
	}

	// 输出
	w.WriteHeader(statusCode)
	io.WriteString(w, text)
	log.Printf("[%d] URI: %s, remoteAddr: %s", statusCode, uri, remoteAddr)
}
