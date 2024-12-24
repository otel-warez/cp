package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("[USAGE] cp src dest")
	}

	data, err := os.ReadFile(os.Args[0])
	if err != nil {
		log.Fatal(fmt.Sprintf("error copying file: %v", err))
	}

	err = os.WriteFile(os.Args[1], data, 0400)

	if err != nil {
		log.Fatal(fmt.Sprintf("error copying file: %v", err))
	}
}
