package ui

import "fmt"

// bubbletea VIEW
func (m model) View() string {
	// header section
	s := "news.ycombinator.com\n\n"

	nav_footer := "\n\n[up/down: navigate, enter: open, q:quit]"
	// check if IsLoading
	if m.IsLoading {
		s += "Loading Stories...\n"
		s += nav_footer
		return s
	}

	// checl if error
	if m.Error != nil {
		s += fmt.Sprintf("Error Loading Stories: %v", m.Error.Error())
		s += nav_footer
		return s
	}

	// now iterate thru the stories and build the string
	for i, story := range m.Stories {
		cursor := " "

		// check if cursor pointing to the item
		if m.Cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s (%d points)\n", cursor, story.Title, story.Score)
	}

	s += nav_footer
	return s

}
