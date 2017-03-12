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
	"io"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type SNSTopicDefinitionMap map[string]SNSTopicDefinition

type SNSTopicDefinition struct {
	Region string `yaml:"region"`
	ARN    string `yaml:"arn"`
}

func topicDefinitionMapFromYaml(r io.Reader) (SNSTopicDefinitionMap, error) {
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	topicsMap := make(SNSTopicDefinitionMap)

	err = yaml.Unmarshal(bytes, &topicsMap)
	if err != nil {
		return nil, err
	}

	return topicsMap, nil
}

func GetSNSTopicDefinitionMap(topicsFile string) (SNSTopicDefinitionMap, error) {
	file, err := os.Open(topicsFile)
	if err != nil {
		return nil, err
	}

	topicDefinitionMap, err := topicDefinitionMapFromYaml(file)
	if err != nil {
		return nil, err
	}

	return topicDefinitionMap, nil
}
