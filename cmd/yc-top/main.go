package main

import (
	"fmt"
	"yc-top/services/hackernews"
)

func main() {
	fmt.Println("Top HackerNews from YC")
	hackernews.FetchTopStoryIDs()
}
