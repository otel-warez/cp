package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("[USAGE] cp src dest")
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(fmt.Sprintf("error copying file: %v", err))
	}

	err = os.WriteFile(os.Args[2], data, 0400)

	if err != nil {
		log.Fatal(fmt.Sprintf("error copying file: %v", err))
	}
}
