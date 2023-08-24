package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func LogError(err error) {
	fmt.Println("Error:", err)
}

func jsonResponse(w http.ResponseWriter, statusCode int, d any) {
	j, err := json.Marshal(d)
	if err != nil {
		LogError(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Contect-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(j)
}
