package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wiliamsouza/warbot/telegram"
)

// telegramCmd represents the telegram command
var telegramCmd = &cobra.Command{
	Use:   "telegram",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		bot := telegram.NewBot("")
		bot.Start()
	},
}

func init() {
	startCmd.AddCommand(telegramCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// telegramCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// telegramCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
