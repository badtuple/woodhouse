package main

import (
	"log"

	"../woodhouse/input"
	"../woodhouse/snippets"
)

func main() {
	inputChannel, err := input.ListenForInput()
	if err != nil {
		log.Fatalf("could not listen for input: %v", err)
	}

	snippets.MatchInputToSnippet(inputChannel)
}
