// main.go
package main

import (
    "bytes"
    "encoding/json"
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

    // Forward the request to the Code Execution Service
    resp, err := http.Post("http://code-execution-service:8080/execute", "application/json", bytes.NewBuffer(req))
    if err != nil {
        http.Error(w, "Failed to execute code", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    // Return the response from the Code Execution Service
    w.WriteHeader(resp.StatusCode)
    var result map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&result)
    json.NewEncoder(w).Encode(result)
}

func main() {
    http.HandleFunc("/execute", executeCode)
    log.Fatal(http.ListenAndServe(":8081", nil))
}
