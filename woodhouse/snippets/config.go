package snippets

import (
	"encoding/json"
	"io/ioutil"
)

var Cfg Config

type Config struct {
	// This key means we should start seeing whether
	// what's typed matches a snippet.
	Leader string `json:"leader"`

	Snippets []Snippet `json:"snippets"`
}

func init() {
	// TODO: obviously this isn't cross-platform.
	// There have to be packages that do this for
	// you in a cross-platform way.
	//
	// We shold really be able to specify the file
	// anyway.

	byt, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(byt, &Cfg)
	if err != nil {
		panic(err)
	}
}
