package input

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

type inputDevice struct {
	id   int
	name string
}

func newDevices() ([]*inputDevice, error) {
	var ret []*inputDevice

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

func newInputDeviceReader(buff []byte, id int) *inputDevice {
	rd := bufio.NewReader(bytes.NewReader(buff))
	rd.ReadLine()
	dev, _, _ := rd.ReadLine()
	splt := strings.Split(string(dev), "=")

	return &inputDevice{
		id:   id,
		name: splt[1],
	}
}
