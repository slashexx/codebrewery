package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"sync"
)

type CodeRequest struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

type CodeResponse struct {
	Output string `json:"output"`
	Error  string `json:"error"`
}

var mu sync.Mutex

func executeCode(language, code string) (string, string) {
	mu.Lock()
	defer mu.Unlock()

	cmd := exec.Command("docker", "run", "--rm", "-i", language, "sh", "-c", code)
	output, err2 := cmd.CombinedOutput()
	return string(output), string(err2.Error())
}

func runHandler(w http.ResponseWriter, r *http.Request) {
	var req CodeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	output, err := executeCode(req.Language, req.Code)
	response := CodeResponse{
		Output: output,
		Error:  err,
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/execute", runHandler)
	fmt.Print("code-execution-service running on port 8081\n")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
