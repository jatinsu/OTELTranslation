// not really useful anymore, just use runme.sh since this literally just runs runme.sh
package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {

	scriptPath := "./runme.sh"

	// Checks if the script file exists
	_, err := os.Stat(scriptPath)
	if os.IsNotExist(err) {
		log.Fatalf("Script file does not exist: %s", scriptPath)
	}

	// Create the command to execute the shell script
	cmd := exec.Command("sh", scriptPath)

	// Redirect the command's output to the current process's output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to run script: %v", err)
	}
}
