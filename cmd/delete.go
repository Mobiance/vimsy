package cmd

import (
	"github.com/mobiance/vimsy/utils"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <name>",
	Short: "Delete a Neovim configuration",
	Long:  `Delete a Neovim configuration by specifying the name of the configuration you want to delete. This command will remove the specified configuration directory and its contents.`,
	Args:  cobra.ExactArgs(1), // Ensure exactly one argument is provided
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if err := utils.DeleteConfig(name); err != nil {
			cmd.PrintErrf("Error deleting configuration: %v\n", err)
			return
		}
		cmd.Printf("Deleted config '%s'! 🎉\n", name)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you can define flags and configuration settings for the delete command.
	// For example, you might want to add a flag for specifying the config directory.
	// deleteCmd.Flags().StringP("config-dir", "d", "", "Directory to store the configuration")
}
