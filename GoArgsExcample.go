package main

import (
	"fmt"
	"log"
	"os"
)

// This is an excample of using Args in running go programm

func main() {
	fmt.Println("Prog was started")
	log.Println("Hello from log")
	args := os.Args[1:]

	// Info from log

	fmt.Println(args[0])
}
