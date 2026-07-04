package prompt_test

import (
	"strings"
	"testing"

	"github.com/andrew-hayworth22/dotlocker/internal/prompt"
)

func TestRepoURL(t *testing.T) {
	str := "git@github.com:me/dotfiles.git"
	input := strings.NewReader(str + "\n")

	got, err := prompt.RepoURL(input)
	if err != nil {
		t.Fatal(err)
	}

	if got != str {
		t.Errorf("got %q", got)
	}
}
