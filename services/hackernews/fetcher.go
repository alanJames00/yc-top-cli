// high level api orchestrators

package hackernews


func FetchTopStories(limit int) ([]Story, error) {
	ids, err := FetchTopStoryIDs()
	if err != nil {
		return nil, err
	}

	// create an array to store the stories
	stories := make([]Story, 0, limit)
	for i, id := range ids {
		if i >= limit {
			break
		}

		story, err := FetchStory(id)
		if err != nil {
			continue // skip that id
		}

		stories = append(stories, story)
	}

	return stories, nil
}
