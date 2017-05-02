// Process Protection project main.go
package main

import (
	"fmt"
	"github.com/alepacheco/Process-Protection/process-protection"
)

func main() {
	fmt.Println("This program will protect its-self.")
	process-protection.Protect()
	fmt.Println("Try killing me with a unelevated tool.")
	for {
	}
}
