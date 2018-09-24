package snippets

import "strings"

type Snippet struct {
	// The text that should be replaced by the snippet
	Symbol string `json:"symbol"`
	// The template of the snippet itself
	Tmpl string `json:"template"`
}

func (s Snippet) SymbolBytes() []byte {
	return []byte(strings.ToUpper(s.Symbol))
}

// Temporary placeholder until a db is connected
func allSnippets() []Snippet {
	return Cfg.Snippets
}
