// cmd/root.go
package main

import (
	"fmt"
	"os"

	"github.com/HyougaOgi/sgen/internal/build"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: sgen [build|serve]")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "build":
		if err := build.Run(); err != nil {
			fmt.Fprintln(os.Stderr, "build error:", err)
			os.Exit(1)
		}
	case "serve":
		fmt.Println("serve not implemented yet")

	case "test":
		fmt.Println("test function executed")
		build.Walk()
	default:
		fmt.Println("Unknown command:", os.Args[1])
		os.Exit(1)
	}
}
