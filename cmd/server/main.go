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
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/bhojpur/suga/pkg/dashboard"
	"github.com/bhojpur/suga/pkg/locales"
	"github.com/bhojpur/suga/pkg/network"
	"github.com/bhojpur/suga/pkg/server"
	"github.com/bhojpur/suga/pkg/training"
	"github.com/bhojpur/suga/pkg/util"
	"github.com/gookit/color"
)

var neuralNetworks = map[string]network.Network{}

func main() {
	port := flag.String("port", "8080", "The port for Bhojpur API and WebSocket.")
	localesFlag := flag.String("re-train", "", "The locale(s) to re-train.")
	flag.Parse()

	// If the locales flag isn't empty then retrain the given models
	if *localesFlag != "" {
		reTrainModels(*localesFlag)
	}

	// Print the Bhojpur Suga ASCII text
	sugaASCII := string(util.ReadFile("res/suga-ascii.txt"))
	fmt.Println(color.FgLightGreen.Render(sugaASCII))

	// Create the authentication token
	dashboard.Authenticate()

	for _, locale := range locales.Locales {
		util.SerializeMessages(locale.Tag)

		neuralNetworks[locale.Tag] = training.CreateNeuralNetwork(
			locale.Tag,
			false,
		)
	}

	// Get port from environment variables if there is
	if os.Getenv("PORT") != "" {
		*port = os.Getenv("PORT")
	}

	// Serves the server
	server.Serve(neuralNetworks, *port)
}

// reTrainModels retrain the given locales
func reTrainModels(localesFlag string) {
	// Iterate locales by separating them by comma
	for _, localeFlag := range strings.Split(localesFlag, ",") {
		path := fmt.Sprintf("res/locales/%s/training.json", localeFlag)
		err := os.Remove(path)

		if err != nil {
			fmt.Printf("Cannot re-train %s model.", localeFlag)
			return
		}
	}
}
