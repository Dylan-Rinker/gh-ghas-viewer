package keys

import "github.com/charmbracelet/bubbles/key"

type ResultKeyMap struct {
	Comment key.Binding
	Close   key.Binding
	Reopen  key.Binding
}

var ResultKeys = ResultKeyMap{
	Comment: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "comment"),
	),
	Close: key.NewBinding(
		key.WithKeys("x"),
		key.WithHelp("x", "close"),
	),
	Reopen: key.NewBinding(
		key.WithKeys("X"),
		key.WithHelp("X", "reopen"),
	),
}

func ResultFullHelp() []key.Binding {
	return []key.Binding{
		ResultKeys.Comment,
		ResultKeys.Close,
		ResultKeys.Reopen,
	}
}
