// Copyright © 2017 David King <davidgwking@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"github.com/davidgwking/snsgun/snsgun"
	"github.com/spf13/cobra"
)

var (
	displayUsage bool
	topicsFile   string
	messagesFile string
)

var FireCmd = &cobra.Command{
	Use:   "fire",
	Short: "Send JSON messages to SNS Topics",
	Long:  `Send groups of JSON messages to one or more SNS Topics`,
	RunE: func(cmd *cobra.Command, args []string) error {
		topicDefinitionMap, err := snsgun.GetSNSTopicDefinitionMap(topicsFile)
		if err != nil {
			return err
		}

		messageDefinitions, err := snsgun.GetSNSMessageDefinitions(messagesFile)
		if err != nil {
			return err
		}

		inputs, err := messageDefinitions.ToSnsPublishInputs(topicDefinitionMap)
		if err != nil {
			return err
		}

		err = snsgun.SendMessages(inputs)
		return err
	},
}

func init() {
	RootCmd.AddCommand(FireCmd)

	FireCmd.Flags().StringVar(&topicsFile, "topics", "./sns_topics.yml", "a yaml file that contains a map from SNS topic name to SNS topic definition")
	FireCmd.Flags().StringVar(&messagesFile, "messages", "./sns_messages.yml", "a yaml file that contains an array of SNS message definitions")
}
