// create module
// > cd f_taskkill>
// > go mod init commandtask

// run the program
// > go run main.go

package main

import (
	"fmt"
	"log"
	"os/exec"
	"sync"
	"time"
)

// killProcess attempts to forcefully terminate a process by its name.
// It takes a WaitGroup pointer to signal when it's done.
func killProcess(wg *sync.WaitGroup, processName string) {
	// Call Done() when the function exits. This is a Go idiom for
	// decrementing the WaitGroup counter, signaling that this goroutine is finished.
	defer wg.Done()

	fmt.Printf("Attempting to terminate process: %s\n", processName)

	// Build the taskkill command.
	cmd := exec.Command("taskkill", "-f", "-im", processName)

	// Use CombinedOutput() to run the command and capture
	// both standard output and standard error.
	output, err := cmd.CombinedOutput()

	if err != nil {
		// Log an error if the command fails to run.
		log.Printf("Error terminating %s: %v\n", processName, err)
		// Print the output for more details on the error.
		fmt.Printf("Command output for %s:\n%s\n", processName, output)
		return
	}

	// If successful, print a success message and the command's output.
	fmt.Printf("Successfully terminated process: %s\n", processName)
	fmt.Printf("Command output for %s:\n%s\n", processName, output)
}

func main() {
	// This loop will run indefinitely until you stop the program with Ctrl+C.
	for {
		fmt.Println("\n--- Starting a new concurrent termination attempt ---")

		// A WaitGroup is a counter that waits for goroutines to finish.
		var wg sync.WaitGroup

		// We are launching two goroutines, so we add 2 to the WaitGroup counter.
		wg.Add(2)

		// Launch the first goroutine to terminate "app.exe".
		go killProcess(&wg, "TiWorker.exe")

		// Launch the second goroutine to terminate "mobsync.exe".
		go killProcess(&wg, "mobsync.exe")

		// Wait() blocks the main function until the WaitGroup counter becomes 0.
		// This ensures the program doesn't exit before the goroutines have finished.
		fmt.Println("Waiting for processes to terminate...")
		wg.Wait()

		fmt.Println("All termination attempts are complete.")

		// Wait for 5 seconds before the next attempt.
		fmt.Println("Waiting 5 seconds before the next attempt...")
		time.Sleep(5 * time.Second)
	}
}
