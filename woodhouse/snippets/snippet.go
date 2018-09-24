package snippets

type Snippet struct {
	// The text that should be replaced by the snippet
	Symbol []byte
	// The template of the snippet itself
	Tmpl string
}

// Temporary placeholder until a db is connected
func allSnippets() []Snippet {
	return []Snippet{
		{[]byte("A"), "single letter test"},
		{[]byte("TT"), "double letter test"},
	}
}
