package main

import (
	"fmt"
	"io"
	"net/http"
)

const (
	portNum        = 8000
	acceptablePath = "/objects/"
)

var dataStore = make(map[string][]byte)

func main() {
	http.HandleFunc(acceptablePath, objectsHandler)

	fmt.Printf("Starting server on port %v\n", portNum)
	err := http.ListenAndServe(fmt.Sprintf(":%v", portNum), nil)
	if err != nil {
		panic(err)
	}
}

func objectsHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len(acceptablePath):]

	switch r.Method {
	case http.MethodPut:
		body := make([]byte, r.ContentLength)
		_, err := r.Body.Read(body)
		if err != nil && err != io.EOF {
			http.Error(w, fmt.Sprintf("Failed to read request body : %v", err.Error()), http.StatusInternalServerError)
			return
		}
		dataStore[key] = body
		w.WriteHeader(http.StatusOK)

	case http.MethodGet:
		if data, ok := dataStore[key]; ok {
			w.WriteHeader(http.StatusOK)
			w.Write(data)
		} else {
			http.NotFound(w, r)
		}

	default:
		http.Error(w, fmt.Sprintf("Method not allowed : %v", r.Method), http.StatusMethodNotAllowed)
	}
}
