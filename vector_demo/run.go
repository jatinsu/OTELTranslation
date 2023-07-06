package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	// Specify the path to the shell script file
	scriptPath := "runme.sh"

	// Check if the script file exists
	_, err := os.Stat(scriptPath)
	if os.IsNotExist(err) {
		log.Fatalf("Script file does not exist: %s", scriptPath)
	}

	// Create the command to execute the shell script
	cmd := exec.Command("sh", scriptPath)

	// Redirect the command's output to the current process's output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to run script: %v", err)
	}
}
