// create module
// > cd b_runner
// > go mod init commandrunner

// run the program
// > go run main.go

package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// Define the command to run.
	// "net" is the command,
	// and "stop samplesvc" are its arguments.
	cmd := exec.Command("net", "stop", "wuauserv")

	// Run the command. The Output() method runs
	// the command and captures its output.
	output, err := cmd.Output()

	if err != nil {
		// If there's an error (e.g., the service doesn't exist), log it.
		log.Printf("Error running command: %v\n", err)
	} else {
		// If successful, print the command's output.
		fmt.Printf("Command output:\n%s\n", output)
	}
}
