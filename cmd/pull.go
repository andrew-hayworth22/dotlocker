package cmd

import (
	"errors"
	"fmt"

	"github.com/andrew-hayworth22/dotlocker/internal/config"
	"github.com/andrew-hayworth22/dotlocker/internal/git"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pulls your dotfile repo down to your machine",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load(viper.GetViper())
		if err != nil {
			return err
		}

		if cfg.RepoURL == "" || cfg.RepoPath == "" {
			return errors.New("run dotlocker init to set up your dofile repository configuration")
		}
		
		created, err := git.PullRepo(cfg.RepoURL, cfg.RepoPath)
		if err != nil {
			return fmt.Errorf("pulling GitHub repo: %w", err)
		}

		if created {
			fmt.Println("successfully cloned your dotfiles")
			return nil
		}

		fmt.Println("successfully pulled your dotfiles")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(pullCmd)
}
