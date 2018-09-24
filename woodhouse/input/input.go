package input

import (
	"fmt"
	"log"
	"os/user"

	"unsafe"
)

const (
	INPUTS        = "/sys/class/input/event%d/device/uevent"
	DEVICE_FILE   = "/dev/input/event%d"
	MAX_FILES     = 255
	MAX_NAME_SIZE = 256
)

// event types
const (
	EV_SYN       = 0x00
	EV_KEY       = 0x01
	EV_REL       = 0x02
	EV_ABS       = 0x03
	EV_MSC       = 0x04
	EV_SW        = 0x05
	EV_LED       = 0x11
	EV_SND       = 0x12
	EV_REP       = 0x14
	EV_FF        = 0x15
	EV_PWR       = 0x16
	EV_FF_STATUS = 0x17
	EV_MAX       = 0x1f
)

var eventsize = int(unsafe.Sizeof(InputEvent{}))

func checkRoot() error {
	u, err := user.Current()
	if err != nil {
		return err
	}
	if u.Uid != "0" {
		return fmt.Errorf("Cannot read device files. Are you running as root?")
	}
	return nil
}

func ListenForInput() (chan InputEvent, error) {
	devs, err := newDevices()
	if err != nil {
		return nil, err
	}

	log.Printf("%v devices found", len(devs))
	for _, val := range devs {
		log.Printf("	ID %v) %v", val.id, val.name)
	}

	// TODO: We're choosing the keyboard on my laptop
	// by index. This won't be the same on every
	// computer so we'll need to figure out how to find
	// which is the one to use.
	dev := devs[0]
	log.Printf("using dev %v (%v)", dev.name, dev.id)
	rd := newKeyLogger(devs[0])

	in, err := rd.read()
	return in, err
}
