package cmd

import (
	"github.com/mobiance/vimsy/utils"
	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch <name>",
	Short: "Switch to a different Neovim configuration",
	Long:  `Switch to a different Neovim configuration by specifying the name of the configuration you want to switch to. This command will update your Neovim setup to use the specified configuration.`,
	Args:  cobra.ExactArgs(1), // Ensure exactly one argument is provided
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if err := utils.SwitchConfig(name); err != nil {
			cmd.PrintErrf("Error adding configuration: %v\n", err)
			return
		}
		cmd.Printf("Switched to '%s' config! 🎉\n", name)
	},
}
func init() {
	rootCmd.AddCommand(switchCmd)

	// Here you can define flags and configuration settings for the add command.
	// For example, you might want to add a flag for specifying the config directory.
	// addCmd.Flags().StringP("config-dir", "d", "", "Directory to store the configuration")
}
