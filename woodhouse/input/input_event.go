package input

import "syscall"

type InputEvent struct {
	// Time the event happened
	Time syscall.Timeval

	// Type of event. We care almost exclusively about
	// Type 0x01 which is a keyboard event.
	Type uint16

	// The key code of the key that was pressed. We use
	// this to lookup the ASCII key with keyCodeMap
	Code uint16

	// The value of the event.
	// 0: a key has been released.
	// 1: a key has been pressed.
	// 2: a key has been held.
	Value int32
}

func (e InputEvent) KeyString() string {
	return keyCodeMap[e.Code]
}

func (e InputEvent) IsPressedEvent() bool {
	return e.Value == 1
}

func (e InputEvent) IsKeyEvent() bool {
	return e.Type == EV_KEY
}
