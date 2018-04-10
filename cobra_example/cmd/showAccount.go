package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// addUserCmd represents the addUser command
var showAccountCmd = &cobra.Command{
	Use:   "showAccount",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			account_type := strings.ToLower(args[0])
			switch account_type {
			case "gmail":
				fmt.Println("example@domain.com")
			case "instagram":
				fmt.Println("instagram")
			case "facebook":
				fmt.Println("example_user")
			case "youtube":
				fmt.Println("example_channel")
			default:
				fmt.Println("account not found")
			}
		} else {
			fmt.Println("He has Facebook, Gmail, Instagram and Youtube accounts")
		}
	},
}

func init() {
	rootCmd.AddCommand(showAccountCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showAccountCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showAccountCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
