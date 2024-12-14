// datastructures for the API responses
package hackernews

type Story struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	URL    string `json:"url"`
	Score  int    `json:"score"`
	Author string `json:"by"`
	Time   int64  `json:"time"`
}
