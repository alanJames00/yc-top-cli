package ui

import (
	"yc-top/services/hackernews"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit // quit the program
		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down":
			if m.Cursor < len(m.Stories)-1 {
				m.Cursor++
			}
		case "r":
			// return the cmd for fetchStories
			return NewModel(), FetchStoriesCmd
		case "enter":
			// same model, nil cmd
			// get the url of current story
			url := m.Stories[m.Cursor].URL
			if err := hackernews.SilentOpenBrowser(url); err != nil {
				// m.Error = err
				return m, nil
			}
			return m, nil
		}

	case StoriesMsg:
		m.IsLoading = false
		m.Stories = msg.Stories
		m.Error = nil
		return m, nil

	case error:
		m.IsLoading = false
		m.Error = msg
		return m, nil
	}

	return m, nil
}
