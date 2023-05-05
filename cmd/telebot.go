/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"

	bot "gopkg.in/telebot.v3"
)

var (
	TeleToken = os.Getenv("TELE_TOKEN")
)

// telebotCmd represents the telebot command
var telebotCmd = &cobra.Command{
	Use:   "telebot",
	Aliases: []string{"start"},
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("telebot %s started", appVersion)

		telebot, err := bot.NewBot(bot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &bot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatal("Please check TELE_TOKEN env variable. %s", err)

			return

		}

		telebot.Handle(bot.OnText, func(m bot.Context) error {

			log.Print(m.Message().Payload, m.Text())

			payload := m.Message().Payload

			switch payload {
			case "hello":
				err = m.Send(fmt.Sprintf("Hello, I'm Telebot %s!", appVersion))
			} 

			return err

		})

		telebot.Start()
	},
}

func init() {
	rootCmd.AddCommand(telebotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// telebotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// telebotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
