package cmd

import (
	"github.com/mobiance/vimsy/utils"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "Create a new Neovim configuration",
	Long:  `Create a new Neovim configuration by specifying the name of the configuration you want to create. This command will set up a new directory for your Neovim config.`,
	Args:  cobra.ExactArgs(1), // Ensure exactly one argument is provided
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if err := utils.CreateConfig(name); err != nil {
			cmd.PrintErrf("Error creating configuration: %v\n", err)
			return
		}
		cmd.Printf("Created new config '%s'! 🎉\n", name)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you can define flags and configuration settings for the create command.
	// For example, you might want to add a flag for specifying the config directory.
	// createCmd.Flags().StringP("config-dir", "d", "", "Directory to store the configuration")
}
