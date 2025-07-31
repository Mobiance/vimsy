package cmd

import (
	"github.com/mobiance/vimsy/utils"
	"github.com/spf13/cobra"
)

var renameCmd = &cobra.Command{
	Use:   "rename <old-name> <new-name>",
	Short: "Rename an existing Neovim configuration",
	Long:  `Rename an existing Neovim configuration by specifying the old name and the new name. This command will change the directory name of the specified configuration.`,
	Args:  cobra.ExactArgs(2), // Ensure exactly two arguments are provided
	Run: func(cmd *cobra.Command, args []string) {
		oldName := args[0]
		newName := args[1]

		if err := utils.RenameConfig(oldName, newName); err != nil {
			cmd.PrintErrf("Error renaming configuration: %v\n", err)
			return
		}
		cmd.Printf("Renamed config from '%s' to '%s'! 🎉\n", oldName, newName)
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)

	// Here you can define flags and configuration settings for the rename command.
	// For example, you might want to add a flag for specifying the config directory.
	// renameCmd.Flags().StringP("config-dir", "d", "", "Directory to store the configuration")
}
