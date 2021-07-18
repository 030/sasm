/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/030/sasm/pkg/slack"
	"github.com/spf13/cobra"
)

// plainCmd represents the plain command
var plainCmd = &cobra.Command{
	Use:   "plain",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("plain called")

		t := slack.Text{Type: "mrkdwn", Text: txt}
		b := []slack.Blocks{{Type: "section", Text: &t}}
		data := slack.Data{Blocks: b, Channel: "#" + slackChannel, Icon: ":robot_face:", Username: "sasm"}
		data.PostMessage(slackToken)
	},
}

var txt string

func init() {
	rootCmd.AddCommand(plainCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	plainCmd.PersistentFlags().StringVarP(&txt, "txt", "t", "", "text to send to slack")
	if err := plainCmd.MarkPersistentFlagRequired("txt"); err != nil {
		log.Fatal(err)
	}

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// plainCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
