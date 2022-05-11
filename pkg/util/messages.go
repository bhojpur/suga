package util

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

// Message contains the message's tag and its contained matched sentences
type Message struct {
	Tag      string   `json:"tag"`
	Messages []string `json:"messages"`
}

var messages = map[string][]Message{}

// SerializeMessages serializes the content of `res/datasets/messages.json` in JSON
func SerializeMessages(locale string) []Message {
	var currentMessages []Message
	err := json.Unmarshal(ReadFile("res/locales/"+locale+"/messages.json"), &currentMessages)
	if err != nil {
		fmt.Println(err)
	}

	messages[locale] = currentMessages

	return currentMessages
}

// GetMessages returns the cached messages for the given locale
func GetMessages(locale string) []Message {
	return messages[locale]
}

// GetMessageByTag returns a message found by the given tag and locale
func GetMessageByTag(tag, locale string) Message {
	for _, message := range messages[locale] {
		if tag != message.Tag {
			continue
		}

		return message
	}

	return Message{}
}

// GetMessage retrieves a message tag and returns a random message chose from res/datasets/messages.json
func GetMessage(locale, tag string) string {
	for _, message := range messages[locale] {
		// Find the message with the right tag
		if message.Tag != tag {
			continue
		}

		// Returns the only element if there aren't more
		if len(message.Messages) == 1 {
			return message.Messages[0]
		}

		// Returns a random sentence
		rand.Seed(time.Now().UnixNano())
		return message.Messages[rand.Intn(len(message.Messages))]
	}

	return ""
}
