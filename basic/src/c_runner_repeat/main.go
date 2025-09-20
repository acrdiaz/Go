// create module
// > cd c_runner_repeat
// > go mod init commandrunner

// run the program
// > go run main.go

package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func main() {
	// This is an infinite loop.
	// It will run forever until you stop it with Ctrl+C.
	for {
		// Define the command to run.
		cmd := exec.Command("net", "stop", "wuauserv")

		// Run the command and capture output.
		output, err := cmd.Output()

		if err != nil {
			log.Printf("Error running command: %v\n", err)
		} else {
			fmt.Printf("Command output:\n%s\n", output)
		}

		// Wait for second(s) before running the command again.
		time.Sleep(1 * time.Second)
	}
}
