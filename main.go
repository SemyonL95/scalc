package main

import (
	"log"
	"os"
	"scalc/set"
)

func main() {
	args := os.Args[1:]

	result, err := set.ParseExpression(args)

	if err != nil {
		log.Fatal(err)
	}

	result.Display()
}
