// bubbletea custom commands
package ui

import (
	"yc-top/services/hackernews"

	tea "github.com/charmbracelet/bubbletea"
)

// command to fetch top stories
func FetchStoriesCmd() tea.Msg {

	// fetch top stories from hackernews
	stories, err := hackernews.FetchTopStories(10)
	if err != nil {
		return err // if there is error, return it as a message
	}

	// return the fetched stories as teaMsg
	return StoriesMsg{Stories: stories}
}
