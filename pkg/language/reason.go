package language

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
	"strings"
)

// ReasonKeywords is for having the keywords in different languages
var ReasonKeywords = map[string]ReasonKeyword{
	"en": {
		That: "that",
		To:   "to",
	},
	"de": {
		That: "das",
		To:   "zu",
	},
	"fr": {
		That: "que",
		To:   "de",
	},
	"es": {
		That: "que",
		To:   "para",
	},
	"ca": {
		That: "que",
		To:   "a",
	},
	"it": {
		That: "quel",
		To:   "per",
	},
	"tr": {
		That: "için",
		To:   "sebebiyle",
	},
	"nl": {
		That: "dat",
		To:   "naar",
	},
	"el": {
		That: "το οποίο",
		To:   "στο",
	},
}

// ReasonKeyword are used to find reason for different languages
type ReasonKeyword struct {
	That string
	To   string
}

// SearchReason returns the reason found in the given sentence for the reminders,
// here is an example: "Remind me that I need to **call mom** tomorrow".
func SearchReason(locale, sentence string) string {
	var response []string

	// Split the given sentence into words
	words := strings.Split(sentence, " ")

	// Initialize the appeared boolean for the keywords
	appeared := false
	// Iterate through the words
	for _, word := range words {
		// Add the word to the response array if the keyword already appeared
		if appeared {
			response = append(response, word)
		}

		// If the keyword didn't appeared and one of the keywords match set the appeared condition
		// to true
		if !appeared && (LevenshteinDistance(word, ReasonKeywords[locale].That) <= 2 ||
			LevenshteinDistance(word, ReasonKeywords[locale].To) < 2) {
			appeared = true
		}
	}

	// Join the words and return the sentence
	return strings.Join(response, " ")
}
