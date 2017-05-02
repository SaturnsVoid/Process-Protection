// Process Protection project main.go
package main

import (
	"fmt"

	"github.com/SaturnsVoid/Process-Protection/process_protection"
)

func main() {
	fmt.Println("This program will protect its-self.")
	process_protection.Protect()
	fmt.Println("Try killing me with a unelevated tool.")
	for {
	}
}
