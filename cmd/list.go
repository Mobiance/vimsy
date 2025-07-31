package cmd

import (
	"github.com/mobiance/vimsy/utils"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available Neovim configurations",
	Long:  `List all available Neovim configurations. This command will display the names of all configurations that are currently managed by Vimsy, allowing you to see which configurations are available for switching.`,
	Run: func(cmd *cobra.Command, args []string) {
		configs, err := utils.ListConfigs()
		if err != nil {
			cmd.PrintErrf("Error listing configurations: %v\n", err)
			return
		}
		if len(configs) == 0 {
			cmd.Println("No configurations found.")
			return
		}
		cmd.Println("Available configurations:")
		for _, config := range configs {
			cmd.Println("- " + config)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you can define flags and configuration settings for the list command.
	// For example, you might want to add a flag for specifying the config directory.
	// listCmd.Flags().StringP("config-dir", "d", "", "Directory to store the configuration")
}
