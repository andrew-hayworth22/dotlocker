/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/andrew-hayworth22/dotlocker/internal/config"
	"github.com/andrew-hayworth22/dotlocker/internal/prompt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().String("repo", "", "The github repository URL that will be synced")
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
				return fmt.Errorf("reading repo URL: %v", err)
			}
		}
		if err := config.Save(repo); err != nil {
			return fmt.Errorf("saving config: %v", err)
		}

		fmt.Println("Initialized dotlocker. Run `dotlocker pull` to fetch your dotfiles.")
		return nil
	},
}
