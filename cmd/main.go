package main

import (
	"log"

	"github.com/MarinX/keylogger"

	"../woodhouse"
)

func main() {
	inputChannel, err := woodhouse.ListenForInput()
	if err != nil {
		log.Fatalf("could not listen for input: %v", err)
	}

	for i := range inputChannel {

		//listen only for key events
		if i.Type == keylogger.EV_KEY {

			// Value is 0 if released
			// Value is 1 if pushed
			// Value is 2 if held
			if i.Value == 1 {
				log.Print(i.KeyString())
			}
		}
	}
}
