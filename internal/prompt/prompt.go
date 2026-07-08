// Package prompt provides prompts for the CLI
package prompt

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// RepoURL prompts the user for the remote dotfile Github repo URL
func RepoURL(r io.Reader) (string, error) {
	fmt.Print("Enter your dotfiles Github repo URL: ")

	reader := bufio.NewReader(r)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}

// RepoPath prompts the user for the local path of the dotfile repo
func RepoPath(r io.Reader) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	defaultPath := filepath.Join(homeDir, ".dotlocker", "dotfiles")
	fmt.Printf("Enter the path where your dotfiles should reside (%s): ", defaultPath)

	reader := bufio.NewReader(r)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)
	if input == "" {
		return defaultPath, nil
	}

	if input == "~" || strings.HasPrefix(input, "~/") {
		input = filepath.Join(homeDir, strings.TrimPrefix(input, "~"))
	}

	return input, nil
}
