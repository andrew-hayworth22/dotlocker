package cmd

import (
	"fmt"
	"os"

	"github.com/andrew-hayworth22/dotlocker/internal/config"
	"github.com/andrew-hayworth22/dotlocker/internal/prompt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().String("repo", "", "The github repository URL that will be synced")
	initCmd.Flags().String("path", "", "The local path where the files reside or should be pulled into")
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes the dotfile CLI tool",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, err := cmd.Flags().GetString("repo")
		if err != nil {
			return err
		}
		if repo == "" {
			repo, err = prompt.RepoURL(os.Stdin)
			if err != nil {
				return fmt.Errorf("reading repo URL: %w", err)
			}
		}

		path, err := cmd.Flags().GetString("path")
		if err != nil {
			return err
		}
		if path == "" {
			path, err = prompt.RepoPath(os.Stdin)
			if err != nil {
				return fmt.Errorf("reading repo path: %w", err)
			}
		}

		cfg := config.Config{
			RepoURL:  repo,
			RepoPath: path,
		}

		if err := config.SaveConfig(viper.GetViper(), cfg); err != nil {
			return fmt.Errorf("saving config: %w", err)
		}

		fmt.Println("Initialized dotlocker. Run `dotlocker pull` to fetch your dotfiles.")
		return nil
	},
}
