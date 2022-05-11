package main

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
	"bufio"
	"fmt"
	"os"
	"strings"

	engine "github.com/bhojpur/suga/pkg/client"
	"github.com/bhojpur/suga/pkg/util/files"
	stamping "github.com/bhojpur/suga/pkg/version"
	"github.com/gookit/color"
)

const (
	logFileName    = "logfile.log"
	configFileName = "config.json"
)

// RequestMessage is the structure that uses entry connections to chat with the websocket
type RequestMessage struct {
	Type        int                    `json:"type"` // 0 for handshakes and 1 for messages
	Content     string                 `json:"content"`
	Token       string                 `json:"user_token"`
	Information map[string]interface{} `json:"information"`
	Locale      string                 `json:"locale"`
}

// ResponseMessage is the structure used to reply to the user through the websocket
type ResponseMessage struct {
	Content     string                 `json:"content"`
	Tag         string                 `json:"tag"`
	Information map[string]interface{} `json:"information"`
}

func main() {
	// Setup the logs and the config file
	files.SetupLog(logFileName)
	config := files.SetupConfig(configFileName)
	files.SetupLogLevel(*config)

	var information map[string]interface{}
	client, err := engine.NewClient(
		fmt.Sprintf("%s:%s", config.Host, config.Port),
		config.SSL,
		&information,
	)
	if err != nil {
		panic(err)
	}

	defer client.Close()

	fmt.Println(color.Magenta.Render("Enter message to Bhojpur " + config.BotName + " server " + stamping.FullVersion() + " or type:"))
	fmt.Printf("- %s to quit\n", color.Green.Render("/quit"))
	fmt.Printf("- %s to change the language\n", color.Green.Render("/lang <en|fr|es...>"))
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

Loop:
	for {
		fmt.Print(">")
		scanner.Scan()
		text := scanner.Text()

		switch {
		case strings.TrimSpace(text) == "":
			fmt.Println(
				color.Red.Render("Please enter a message"),
			)
			continue Loop
		case text == "/quit":
			os.Exit(0)
		case strings.HasPrefix(text, "/lang"):
			arguments := strings.Split(text, " ")[1:]

			if len(arguments) != 1 {
				fmt.Println(
					color.Red.Render("Wrong number of arguments, language command should contain only the locale"),
				)
				continue Loop
			}

			client.Locale = arguments[0]
			fmt.Printf("Language changed to %s.\n", color.Magenta.Render(arguments[0]))

			continue Loop
		}

		response, err := client.SendMessage(text)
		if err != nil {
			continue
		}

		fmt.Printf("%s> %s\n", color.Magenta.Render(config.BotName), response.Content)
	}
}
