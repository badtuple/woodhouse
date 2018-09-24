package woodhouse

import (
	"log"

	"github.com/MarinX/keylogger"
)

func ListenForInput() (chan keylogger.InputEvent, error) {
	devs, err := keylogger.NewDevices()
	if err != nil {
		return nil, err
	}

	log.Printf("%v devices found", len(devs))
	for _, val := range devs {
		log.Printf("	ID %v) %v", val.Id, val.Name)
	}

	// TODO: We're choosing the keyboard on my laptop
	// by index. This won't be the same on every
	// computer so we'll need to figure out how to find
	// which is the one to use.
	dev := devs[0]
	log.Printf("using dev %v (%v)", dev.Name, dev.Id)
	rd := keylogger.NewKeyLogger(devs[0])

	in, err := rd.Read()
	return in, err
}
