package main

import (
	"encoding/json"
	"net/http"
)

func Error(w http.ResponseWriter, error string, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	output, err := json.Marshal(map[string]string{"error": error})
	if err != nil {
		panic(err)
	}
	w.Write(output)
}

func Respond(w http.ResponseWriter, data []byte, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
