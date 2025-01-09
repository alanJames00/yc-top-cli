// functions for opening the story URL in browser for major OS
package hackernews

import "github.com/pkg/browser"

func OpenBrowser(url string) error {
	err := browser.OpenURL(url)
	if err != nil {
		return err
	}
	
	// opened succesfully
	return nil
}
