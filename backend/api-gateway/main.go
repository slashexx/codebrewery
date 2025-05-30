// First ever comment from KDE let's go ?
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type CodeRequest struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

func enableCors(w http.ResponseWriter) {
	// Allow requests from all origins temporarily for demonstration
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "false") // Changed to false since we're using "*" for origin
}

func executeCode(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	fmt.Printf("Received %s request for %s\n", r.Method, r.URL.Path)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var req CodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request: %v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	log.Printf("Request body: %+v", req)

	reqBody, err := json.Marshal(req)
	if err != nil {
		log.Printf("Error marshaling request: %v", err)
		http.Error(w, "Failed to marshal request", http.StatusInternalServerError)
		return
	}

	url := "https://codebrewery-code-execution-service.onrender.com/execute"
	log.Printf("Sending request to Code Execution Service: %s", url)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Printf("Error contacting code execution service: %v", err)
		http.Error(w, "Failed to execute code", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	log.Printf("Received response status: %d", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		log.Printf("Code execution service returned status: %d", resp.StatusCode)
		responseBody, _ := ioutil.ReadAll(resp.Body)
		log.Printf("Response body: %s", string(responseBody))
		http.Error(w, "Code execution service returned an error", http.StatusInternalServerError)
		return
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("Error decoding response: %v", err)
		http.Error(w, "Failed to decode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/execute", executeCode)
	fmt.Println("API Gateway running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
