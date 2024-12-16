package main

import (
	"log"
	"yc-top/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// init the bubble tea model
	m := ui.NewModel()

	// create a program
	p := tea.NewProgram(m)

	// run
	_, err := p.Run()
	if err != nil {
		log.Fatal("Error running bubbletea program")
	}
}
