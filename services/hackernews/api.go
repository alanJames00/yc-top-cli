// official hackernews apis
package hackernews

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const baseUrl = "https://hacker-news.firebaseio.com/v0"

// fetch top 500 story ids
func FetchTopStoryIDs() ([]int, error) {
	resp, err := http.Get(baseUrl + "/topstories.json?")
	if err != nil {
		fmt.Println("Error Occured while fetching")
		return nil, err
	}
	defer resp.Body.Close()

	var ids []int

	err = json.NewDecoder(resp.Body).Decode(&ids)
	if err != nil {
		fmt.Println("Error Occured while decoding the json resp")
		return nil, err
	}

	return ids, nil
}

// fetch story from id
func FetchStory(id int) (Story, error) {
	resp, err := http.Get(baseUrl + "/item/" + fmt.Sprintf("%d.json", id))
	if err != nil {
		fmt.Println("Error Occured while fetching")
		return Story{}, err
	}
	defer resp.Body.Close()

	var story Story
	err = json.NewDecoder(resp.Body).Decode(&story)
	if err != nil {
		fmt.Println("Error Occured while decoding the json resp")
		return Story{}, err
	}

	return story, nil
}
