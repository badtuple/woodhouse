package main

import (
	"log"

	"../woodhouse/input"
)

func main() {
	inputChannel, err := input.ListenForInput()
	if err != nil {
		log.Fatalf("could not listen for input: %v", err)
	}

	for i := range inputChannel {

		//listen only for new pressed key events
		if i.IsKeyEvent() && i.IsPressedEvent() {
			log.Print(i.KeyString())
		}
	}
}
