package training

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
	"fmt"
	"os"

	"github.com/bhojpur/suga/pkg/analysis"
	"github.com/bhojpur/suga/pkg/network"
	"github.com/bhojpur/suga/pkg/util"
	"github.com/gookit/color"
)

// TrainData returns the inputs and outputs for the neural network
func TrainData(locale string) (inputs, outputs [][]float64) {
	words, classes, documents := analysis.Organize(locale)

	for _, document := range documents {
		outputRow := make([]float64, len(classes))
		bag := document.Sentence.WordsBag(words)

		// Change value to 1 where there is the document Tag
		outputRow[util.Index(classes, document.Tag)] = 1

		// Append data to inputs and outputs
		inputs = append(inputs, bag)
		outputs = append(outputs, outputRow)
	}

	return inputs, outputs
}

// CreateNeuralNetwork returns a new neural network which is loaded from res/training.json or
// trained from TrainData() inputs and targets.
func CreateNeuralNetwork(locale string, ignoreTrainingFile bool) (neuralNetwork network.Network) {
	// Decide if the network is created by the save or is a new one
	saveFile := "res/locales/" + locale + "/training.json"

	_, err := os.Open(saveFile)
	// Train the model if there is no training file
	if err != nil || ignoreTrainingFile {
		inputs, outputs := TrainData(locale)

		neuralNetwork = network.CreateNetwork(locale, 0.1, inputs, outputs, 50)
		neuralNetwork.Train(200)

		// Save the neural network in res/training.json
		neuralNetwork.Save(saveFile)
	} else {
		fmt.Printf(
			"%s %s\n",
			color.FgBlue.Render("Loading the Neural Network from"),
			color.FgRed.Render(saveFile),
		)
		// Initialize the intents
		analysis.SerializeIntents(locale)
		neuralNetwork = *network.LoadNetwork(saveFile)
	}

	return
}
