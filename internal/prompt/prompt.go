// Package prompt provides prompts for the CLI
package prompt

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func RepoURL(r io.Reader) (string, error) {
	fmt.Print("Enter your dotfiles Github repo URL: ")

	reader := bufio.NewReader(r)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}
