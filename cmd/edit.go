package cmd

import (
	"github.com/mobiance/vimsy/utils"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [name]",
	Short: "Edit an existing Neovim configuration",
	Long:  `Edit an existing Neovim configuration by specifying the name, or open the currently active configuration if no name is provided.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var name string
		if len(args) == 1 {
			name = args[0]
		}

		if err := utils.EditConfig(name); err != nil {
			cmd.PrintErrf("Error editing configuration: %v\n", err)
			return
		}

		if name == "" {
			cmd.Println("Opened current config for editing! 🎉")
		} else {
			cmd.Printf("Opened '%s' config for editing! 🎉\n", name)
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you can define flags and configuration settings for the edit command.
	// For example, you might want to add a flag for specifying the editor to use.
	// editCmd.Flags().StringP("editor", "e", "nvim", "Editor to use for editing the configuration")
}
