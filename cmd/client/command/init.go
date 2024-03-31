package command

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "Chat client",
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to chat",
	Run:   connect,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(connectCmd)

	connectCmd.Flags().StringP("email", "e", "", "email")
	err := connectCmd.MarkFlagRequired("email")
	if err != nil {
		log.Fatalf("failed to mark username flag as required: %s\n", err.Error())
	}
	connectCmd.Flags().StringP("password", "p", "", "password")
	err = connectCmd.MarkFlagRequired("password")
	if err != nil {
		log.Fatalf("failed to mark password flag as required: %s\n", err.Error())
	}
	connectCmd.Flags().StringP("chat", "c", "", "chat id")
	err = connectCmd.MarkFlagRequired("chat")
	if err != nil {
		log.Fatalf("failed to mark chat flag as required: %s\n", err.Error())
	}
}
