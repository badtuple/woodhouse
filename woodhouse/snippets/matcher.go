package snippets

import (
	"log"

	"../input"
	"github.com/atotto/clipboard"
	"github.com/go-vgo/robotgo"
)

const (
	// maximum symbol length. This is used when reading
	// into the symbolBuffer.  24 is likely way too
	// small, but it'll need to be something.
	maxSymbolSize int = 24
)

var (
	// This key means we should start seeing whether
	// what's typed matches a snippet.
	//
	// TODO: make it a config option, and handle more
	// than one char
	leader string = ","

	// The actual buffer we use to keep track of key
	// strokes to see if it matches a symbol.  Does
	// not keep track of the leader.
	symbolBuffer = make([]byte, 0, maxSymbolSize)
)

func MatchInputToSnippet(inputChannel chan input.InputEvent) {
	var leaderTriggered bool
	for in := range inputChannel {
		if in.KeyString() == leader {
			leaderTriggered = true
			symbolBuffer = symbolBuffer[:0]
			continue
		}

		//listen only for new pressed key events
		if leaderTriggered && in.IsKeyEvent() && in.IsPressedEvent() {
			ks := in.KeyString()

			// Only accept single characters, and not special keys
			// like ESC or L_CTRL. This check depends on us using
			// strings in the key map lookup.
			if len(ks) != 1 {
				continue
			}

			symbolBuffer = append(symbolBuffer, ks[0])

			snippet := checkBufferForMatch()
			if len(symbolBuffer) > maxSymbolSize {
				symbolBuffer = symbolBuffer[:0]
				leaderTriggered = false
			}

			if snippet == nil {
				continue
			}

			log.Printf("found snippet %+v", snippet)
			symbolBuffer = symbolBuffer[:0]
			leaderTriggered = false

			err := pasteSnippet(*snippet)
			if err != nil {
				log.Printf("could not read or write to clipboard: %v", err.Error())
			}
		}
	}
}

// Right now we brute force our way through the snippets to keep it simple.
// If this doesn't scale we may need to build a prefix tree out of the Snippets.
//
// TODO: Right now the match is NOT case sensitive, but it should be. This is
// due to the fact that how we're getting keyboard input doesn't distinguish
// between "a" and "A".  I'm not sure if we need to go a different route to
// get keyboard input or if we need to model the full keyboard to know the
// case/state.
func checkBufferForMatch() *Snippet {
SnippetLoop:
	for _, s := range allSnippets() {
		log.Printf("checking %v against %v", string(s.Symbol), string(symbolBuffer))
		if len(symbolBuffer) != len(s.Symbol) {
			continue
		}

		for i, r := range s.Symbol {
			if r != symbolBuffer[i] {
				break SnippetLoop
			}

			// found it!
			return &s
		}
	}
	return nil
}

func pasteSnippet(s Snippet) error {
	old, err := clipboard.ReadAll()
	if err != nil {
		// this may mean that there was nothing in the clipboard
		log.Println("could not read from clipboard")
	}

	log.Printf("removed old string from keyboard: %v", old)

	err = clipboard.WriteAll(s.Tmpl)
	if err != nil {
		return err
	}

	// delete typed symbol
	for i := 0; i < len(s.Symbol)+1; i++ {
		keyTap("backspace")
	}

	// paste snippet in
	keyTap("v", "control")

	// replace old clipboard
	err = clipboard.WriteAll(old)
	return err
}

// We're offloading this to the robotgo lib because it's
// crossplatform and all ways I've seen to do this so far
// are crazy gnarly.  We'll see if there's a way to bring
// it inhouse later so that there aren't any blackbox
// deps. Since we'd likely have to interface with C it may
// be worth leaving as is.
func keyTap(tapKey string, args ...interface{}) {
	robotgo.KeyTap(tapKey, args...)
}
