// First ever comment from KDE let's go ?
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

func enableCors(w http.ResponseWriter) {
    w.Header().Set("Access-Control-Allow-Origin", "https://codebrewery.vercel.app") 
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func executeCode(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Received %s request for %s\n", r.Method, r.URL.Path)
    enableCors(w)

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

    // Marshal the request to JSON
    reqBody, err := json.Marshal(req)
    if err != nil {
        log.Printf("Error marshaling request: %v", err)
        http.Error(w, "Failed to marshal request", http.StatusInternalServerError)
        return
    }

    // Forward the request to the Code Execution Service
    resp, err := http.Post("https://codebrewery-code-execution-service.onrender.com/execute", "application/json", bytes.NewBuffer(reqBody))
    if err != nil {
        log.Printf("Error contacting code execution service: %v", err)
        http.Error(w, "Failed to execute code", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        log.Printf("Code execution service returned status: %d", resp.StatusCode)
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
