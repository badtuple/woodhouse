package input

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

type InputDevice struct {
	Id   int
	Name string
}

func NewDevices() ([]*InputDevice, error) {
	var ret []*InputDevice

	if err := checkRoot(); err != nil {
		return ret, err
	}

	for i := 0; i < MAX_FILES; i++ {
		buff, err := ioutil.ReadFile(fmt.Sprintf(INPUTS, i))
		if err != nil {
			continue
		}
		ret = append(ret, newInputDeviceReader(buff, i))
	}

	return ret, nil
}

func newInputDeviceReader(buff []byte, id int) *InputDevice {
	rd := bufio.NewReader(bytes.NewReader(buff))
	rd.ReadLine()
	dev, _, _ := rd.ReadLine()
	splt := strings.Split(string(dev), "=")

	return &InputDevice{
		Id:   id,
		Name: splt[1],
	}
}
