package cmd

import (
	"github.com/mobiance/vimsy/utils"
	"github.com/spf13/cobra"
)

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Show the current Neovim configuration",
	Long:  `Display the name of the currently active Neovim configuration. This command helps you identify which configuration is currently in use.`,
	Run: func(cmd *cobra.Command, args []string) {
		configName, err := utils.CurrentConfig()
		if err != nil {
			cmd.PrintErrf("Error retrieving current configuration: %v\n", err)
			return
		}
		cmd.Printf("Current configuration: %s\n", configName)
	},
}

func init() {
	rootCmd.AddCommand(currentCmd)

	// Here you can define flags and configuration settings for the current command.
	// For example, you might want to add a flag for specifying the config directory.
	// currentCmd.Flags().StringP("config-dir", "d", "", "Directory to store the configuration")
}
