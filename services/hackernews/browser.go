// functions for opening the story URL in browser for major OS
package hackernews

import (
	"os/exec"
	"runtime"
)

// silent browser opening without printing to stdout
func SilentOpenBrowser(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:  // linux and other unix-like systems
		cmd = exec.Command("xdg-open", url)
	}

	// surpress std outputs
	cmd.Stderr = nil
	cmd.Stdout = nil

	return cmd.Run()
}
