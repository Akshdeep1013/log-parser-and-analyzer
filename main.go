package main

import (
	"fmt"
	"log-parser-and-analyzer/repository"
)

func main() {
	log, err := repository.Load()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("File has been loaded into memory wiht %d logs\n", len(*log))
}
