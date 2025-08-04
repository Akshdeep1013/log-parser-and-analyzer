package main

import (
	"fmt"
	"log-parser-and-analyzer/repository"
	"log-parser-and-analyzer/service"
)

func main() {
	logs, err := repository.Load()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	logInsight, err := service.GetLogInsight(logs)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if len(*logs) == 0 {
		fmt.Printf("\nLog file have 0 logs")
		return
	}
	fmt.Printf("\nFile has been loaded into memory wiht %d logs\n", len(*logs))
	fmt.Printf("\nLog Insight\n %+v\n\n", logInsight)
}
