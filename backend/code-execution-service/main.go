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
	Language string `json:"language"`
	Code string `json:"code"`
}

type Response struct {
	Output string `json:"output"`
	Error  string `json:"error,omitempty"`
}

func executeCode(language string, code string) (string, error) {
	var cmd *exec.Cmd
	var tmpFile *os.File
	var err error

	// Create a temporary file
	tmpFile, err = os.CreateTemp("", "*."+language)
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

	
	switch language {
	case "go":
		cmd = exec.Command("go", "run", tmpFile.Name())
	case "python":
		cmd = exec.Command("python3", tmpFile.Name())
	case "c":
		execCmd := exec.Command("gcc", tmpFile.Name(), "-o", tmpFile.Name()[:len(tmpFile.Name())-2]) 
		if err := execCmd.Run(); err != nil {
			cmdOutput, _ := execCmd.CombinedOutput() 
			return string(cmdOutput), err
		}
		cmd = exec.Command(tmpFile.Name()[:len(tmpFile.Name())-2]) 
	case "cpp":
		execCmd := exec.Command("g++", tmpFile.Name(), "-o", tmpFile.Name()[:len(tmpFile.Name())-4]) 
		if err := execCmd.Run(); err != nil {
			cmdOutput, _ := execCmd.CombinedOutput() 
			return string(cmdOutput), err
		}
		cmd = exec.Command(tmpFile.Name()[:len(tmpFile.Name())-4]) // Execute the compiled binary
	case "java":
		execCmd := exec.Command("javac", tmpFile.Name())
		if err := execCmd.Run(); err != nil {
			
			cmdOutput, _ := execCmd.CombinedOutput()
			return string(cmdOutput), err
		}
		className := tmpFile.Name()[:len(tmpFile.Name())-5] 
		cmd = exec.Command("java", className)
		
		
		cmdOutput, err := cmd.CombinedOutput()
		if err != nil {
			return string(cmdOutput), err
		}
		return string(cmdOutput), nil
	
	default:
		return "", fmt.Errorf("unsupported language: %s", language)
	}

	// Execute the command and capture output
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

    output, err := executeCode(req.Language, req.Code)
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
