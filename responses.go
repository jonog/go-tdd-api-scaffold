package main

import (
	"encoding/json"
	"net/http"
)

func Error(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	output, err := json.Marshal(map[string]string{"error": error})
	if err != nil {
		panic(err)
	}
	w.Write(output)
}

func Respond(w http.ResponseWriter, data []byte, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
