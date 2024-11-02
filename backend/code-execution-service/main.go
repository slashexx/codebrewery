package main

import (
	"encoding/json"
	"fmt"
	// "io"
	"net/http"
	"os"
	"os/exec"
)

type Request struct {
	Code string `json:"code"`
}

type Response struct {
	Output string `json:"output"`
	Error  string `json:"error,omitempty"`
}

func executeCode(code string) (string, error) {
	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "*.go")
	if err != nil {
		return "", err
	}
	defer func() {
		if err := os.Remove(tmpFile.Name()); err != nil {
			fmt.Println("Error removing temp file:", err)
		}
	}()

	// Write the code to the temporary file
	if _, err := tmpFile.Write([]byte(code)); err != nil {
		return "", err
	}
	tmpFile.Close()

	// Execute the code
	cmd := exec.Command("go", "run", tmpFile.Name())
	cmdOutput, err := cmd.CombinedOutput()

	return string(cmdOutput), err
}

func enableCors(w http.ResponseWriter) {
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080") 
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS") 
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func executeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Received %s request for %s\n", r.Method, r.URL.Path)
    enableCors(w) 

    if r.Method == http.MethodOptions {
		
        w.WriteHeader(http.StatusOK)
        return
    }

    var req Request
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    output, err := executeCode(req.Code)
    res := Response{Output: output}
    if err != nil {
        res.Error = err.Error()
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(res)
}


func main() {
	http.HandleFunc("/execute", executeHandler)
	fmt.Println("Code Execution Service running on port 8081\n")
	http.ListenAndServe(":8081", nil)
}
