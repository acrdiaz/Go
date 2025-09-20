// create module
// > cd d_runner_concurrent
// > go mod init commandrunner

// run the program
// > go run main.go

package main

import (
	"fmt"
	"log"
	"os/exec"
	"sync"
)

// stopService attempts to stop a Windows service with the given name.
// It takes a WaitGroup pointer to signal when it's done.
func stopService(wg *sync.WaitGroup, serviceName string) {
	// Call Done() when the function exits. This is a Go idiom for
	// decrementing the WaitGroup counter, signaling that this goroutine is finished.
	defer wg.Done()

	fmt.Printf("Attempting to stop service: %s\n", serviceName)

	// Build the command to execute.
	cmd := exec.Command("net", "stop", serviceName)

	// Use CombinedOutput() to capture both standard output and standard error.
	output, err := cmd.CombinedOutput()

	if err != nil {
		// Log the error, but don't stop the program.
		log.Printf("Error stopping service %s: %v\n", serviceName, err)
		// Print the command's output for more details on the error.
		fmt.Printf("Command output for %s:\n%s\n", serviceName, output)
		return
	}

	// If successful, print a success message and the command's output.
	fmt.Printf("Successfully stopped service: %s\n", serviceName)
	fmt.Printf("Command output for %s:\n%s\n", serviceName, output)
}

func main() {
	// A WaitGroup is a counter that waits for goroutines to finish.
	var wg sync.WaitGroup

	// We are launching two goroutines, so we add 2 to the WaitGroup counter.
	wg.Add(2)

	// Launch the first goroutine to stop
	// The `go` keyword is all you need to run a function concurrently.
	go stopService(&wg, "wuauserv")

	// Launch the second goroutine to stop
	go stopService(&wg, "sysmain")

	// Wait() blocks the main function until the WaitGroup counter becomes 0.
	// This ensures the program doesn't exit before the goroutines have finished.
	fmt.Println("Waiting for services to stop...")
	wg.Wait()

	fmt.Println("All service stopping attempts are complete.")
}
