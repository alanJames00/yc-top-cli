// high level api orchestrators

package hackernews

import "fmt"

func FetchTopStories(limit int) ([]Story, error) {
	ids, err := FetchTopStoryIDs()
	if err != nil {
		return nil, err
	}

	// limit the number of stories to fetch
	if (len(ids)) > limit {
		ids = ids[:limit]
	}

	// channel to collect fetched stories and errors
	storiesCh := make(chan Story, limit)
	errCh := make(chan error, limit)

	// send off goroutines to fetch stories concurrently
	for _, id := range ids {
		go func(id int) {
			story, err := FetchStory(id)
			if err != nil {
				errCh <- err
			}
			storiesCh <- story
		}(id)
	}

	// collect the results
	stories := make([]Story, 0, limit)
	for i := 0; i < limit; i++ {
		select {
		case story := <- storiesCh:
			stories = append(stories, story)
		case err := <- errCh:
			// log or handle errors, but continue collecting other stories
			fmt.Printf("error fetching stories: %v", err)
		}
	}

	// close both channels
	close(storiesCh)
	close(errCh)

	return stories, nil
}
