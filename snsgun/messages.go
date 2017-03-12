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

package snsgun

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/service/sns"

	"gopkg.in/yaml.v2"
)

type SNSMessageDefinitions []SNSMessageDefinition

func (messageDefinitions SNSMessageDefinitions) ToSnsPublishInputs(topicDefinitionMap SNSTopicDefinitionMap) ([]*sns.PublishInput, error) {
	inputs := make([]*sns.PublishInput, 0, 10)

	for _, messageDefinition := range messageDefinitions {
		input := &sns.PublishInput{}
		input.SetMessage(messageDefinition.Message)

		topicName := messageDefinition.SNSTopicName
		topic, ok := topicDefinitionMap[topicName]
		if !ok {
			err := fmt.Errorf("failed to find SNS Topic definition for %s", topicName)
			return nil, err
		}
		input.SetTopicArn(topic.ARN)

		inputs = append(inputs, input)
	}

	return inputs, nil
}

type SNSMessageDefinition struct {
	SNSTopicName string `yaml:"topicName"`
	Message      string `yaml:"message"`
}

func messageDefinitionsFromYaml(r io.Reader) (SNSMessageDefinitions, error) {
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	messages := make(SNSMessageDefinitions, 0, 10)

	err = yaml.Unmarshal(bytes, &messages)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func GetSNSMessageDefinitions(messagesFile string) (SNSMessageDefinitions, error) {
	file, err := os.Open(messagesFile)
	if err != nil {
		return nil, err
	}

	messageDefinitions, err := messageDefinitionsFromYaml(file)
	if err != nil {
		return nil, err
	}

	return messageDefinitions, nil
}
