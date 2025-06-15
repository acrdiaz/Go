/*
create a folder for your Go code

	mkdir -p ~/go-projects/hello-go
	cd ~/go-projects/hello-go

initialize a module

	go mod init hello-go

install useful tools that help with code quality, formatting, and testing:

	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/stretchr/testify@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest

run

	cd challenge/AA/hello-go
	go run main.go fizzbuzz.go
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(Salute())

	mySlice := FizzBuzz(15)

	fmt.Println(mySlice)
}

func Salute() string {
	return "🚀 You're ready to become a Go master!"
}
