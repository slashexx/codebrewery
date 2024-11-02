// main.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type CodeRequest struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

func executeCode(w http.ResponseWriter, r *http.Request) {
	var req CodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Marshal the request to JSON
	reqBody, err := json.Marshal(req)
	if err != nil {
		http.Error(w, "Failed to marshal request", http.StatusInternalServerError)
		return
	}

	// Forward the request to the Code Execution Service
	resp, err := http.Post("http://code-execution-service:8080/execute", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		http.Error(w, "Failed to execute code", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Return the response from the Code Execution Service
	w.WriteHeader(resp.StatusCode)
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		http.Error(w, "Failed to decode response", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/execute", executeCode)
	fmt.Println("API Gateway running on port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
