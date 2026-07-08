package git

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// PullRepo refreshes the local dotfile directory by cloning or pulling the remote PullRepo
// Returns (true, nil) if the repo was freshly pulled and created and (false, nil) if the repo
// was refreshed via a git pull
func PullRepo(url, path string) (bool, error) {
	exists := true

	_, err := os.Stat(path)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return false, fmt.Errorf("reading existing directory: %w", err) 
		}
		exists = false
	}

	if !exists {
		_, err = exec.Command("git", "clone", "--depth", "1", url, path).Output()
		if err != nil {
			if exitErr, ok := errors.AsType[*exec.ExitError](err); ok {
				return false, fmt.Errorf("git clone failed: %s", strings.TrimSpace(string(exitErr.Stderr)))
			}
			return false, fmt.Errorf("running git command: %w", err)
		}
		return true, nil
	}

	cmd := exec.Command("git", "pull")
	cmd.Dir = path

	_, err = cmd.Output()
	if err != nil {
		if exitErr, ok := errors.AsType[*exec.ExitError](err); ok {
			return false, fmt.Errorf("git pull failed: %s", strings.TrimSpace(string(exitErr.Stderr)))
		}
		return false, fmt.Errorf("running git command: %w: ", err)
	}

	return false, nil
}
