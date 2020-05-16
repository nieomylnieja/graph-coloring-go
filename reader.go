package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type DimacsReader struct {
}

func (r DimacsReader) Read() {
	file, err := os.Open("instances/fpsol2.i.1.col")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
