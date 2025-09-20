// cmd set variables
// > set GOPATH=<root_dir>\github\Go\ ...
//   Go 1.18+ DOESNT'T require GOPATH,
//   but it's still good to organize your projects.
// > set GOBIN=<root_dir>\github\Go\bin

// create module
// > cd <project_dir>
// > go mod init main/my_project

// run the program
// > go run .
// or
// > go run hi.go
// or
// > go run src/main/hi.go

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hi")
}
