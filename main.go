package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Dylan-Rinker/gh-ghas-viewer/ui"
	"github.com/Dylan-Rinker/gh-ghas-viewer/ui/markdown"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
)

func createModel(debug bool) (ui.Model, *os.File) {
	var loggerFile *os.File
	var err error

	if debug {
		loggerFile, err = tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("Error setting up logger")
		}
	}

	return ui.NewModel(), loggerFile
}

func main() {
	debug := flag.Bool(
		"debug",
		false,
		"passing this flag will allow writing debug output to debug.log",
	)
	flag.Parse()

	// see https://github.com/charmbracelet/lipgloss/issues/73
	markdown.InitializeMarkdownStyle(termenv.HasDarkBackground())

	model, logger := createModel(*debug)
	if logger != nil {
		defer logger.Close()
	}

	p := tea.NewProgram(
		model,
		tea.WithAltScreen(),
	)
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
