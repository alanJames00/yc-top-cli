package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"yc-top/services/hackernews"
)

// bubbletea model
type model struct {
	IsLoading bool  // is the content being loaded
	Error     error // error occured during the data fetching
	Stories   []hackernews.Story
	Cursor    int
}

// initial mode function
func (m model) Init() tea.Cmd {
	return FetchStoriesCmd
}

type StoriesMsg struct {
	Stories []hackernews.Story
}

func NewModel() model {
	return model{
		IsLoading: true,
		Cursor:    0,
		Error:     nil,
		Stories:   []hackernews.Story{},
	}
}
