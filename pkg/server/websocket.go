package server

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
	"net/http"
	"reflect"

	"github.com/bhojpur/suga/pkg/analysis"
	"github.com/bhojpur/suga/pkg/locales"
	"github.com/bhojpur/suga/pkg/modules/start"
	"github.com/bhojpur/suga/pkg/user"
	"github.com/bhojpur/suga/pkg/util"
	"github.com/gookit/color"
	"github.com/gorilla/websocket"
)

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// RequestMessage is the structure that uses entry connections to chat with the websocket
type RequestMessage struct {
	Type        int              `json:"type"` // 0 for handshakes and 1 for messages
	Content     string           `json:"content"`
	Token       string           `json:"user_token"`
	Locale      string           `json:"locale"`
	Information user.Information `json:"information"`
}

// ResponseMessage is the structure used to reply to the user through the websocket
type ResponseMessage struct {
	Content     string           `json:"content"`
	Tag         string           `json:"tag"`
	Information user.Information `json:"information"`
}

// SocketHandle manages the entry connections and reply with the neural network
func SocketHandle(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	fmt.Println(color.FgGreen.Render("A new connection has been opened"))

	for {
		// Read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		// Unmarshal the json content of the message
		var request RequestMessage
		if err = json.Unmarshal(msg, &request); err != nil {
			continue
		}

		// Set the information from the client into the cache
		if reflect.DeepEqual(user.GetUserInformation(request.Token), user.Information{}) {
			user.SetUserInformation(request.Token, request.Information)
		}

		// If the type of requests is a handshake then execute the start modules
		if request.Type == 0 {
			start.ExecuteModules(request.Token, request.Locale)

			message := start.GetMessage()
			if message != "" {
				// Generate the response to send to the user
				response := ResponseMessage{
					Content:     message,
					Tag:         "start module",
					Information: user.GetUserInformation(request.Token),
				}

				bytes, err := json.Marshal(response)
				if err != nil {
					panic(err)
				}

				if err = conn.WriteMessage(msgType, bytes); err != nil {
					continue
				}
			}

			continue
		}

		// Write message back to browser
		response := Reply(request)
		if err = conn.WriteMessage(msgType, response); err != nil {
			continue
		}
	}
}

// Reply takes the entry message and returns an array of bytes for the answer
func Reply(request RequestMessage) []byte {
	var responseSentence, responseTag string

	// Send a message from res/datasets/messages.json if it is too long
	if len(request.Content) > 500 {
		responseTag = "too long"
		responseSentence = util.GetMessage(request.Locale, responseTag)
	} else {
		// If the given locale is not supported yet, set english
		locale := request.Locale
		if !locales.Exists(locale) {
			locale = "en"
		}

		responseTag, responseSentence = analysis.NewSentence(
			locale, request.Content,
		).Calculate(*cache, neuralNetworks[locale], request.Token)
	}

	// Marshall the response in json
	response := ResponseMessage{
		Content:     responseSentence,
		Tag:         responseTag,
		Information: user.GetUserInformation(request.Token),
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	return bytes
}
