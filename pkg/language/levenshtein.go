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

// LevenshteinDistance calculates the Levenshtein Distance between two given words and returns it.
func LevenshteinDistance(first, second string) int {
	// Returns the length if it's empty
	if first == "" {
		return len(second)
	}
	if second == "" {
		return len(first)
	}

	if first[0] == second[0] {
		return LevenshteinDistance(first[1:], second[1:])
	}

	a := LevenshteinDistance(first[1:], second[1:])
	if b := LevenshteinDistance(first, second[1:]); a > b {
		a = b
	}

	if c := LevenshteinDistance(first[1:], second); a > c {
		a = c
	}

	return a + 1
}

// LevenshteinContains checks for a given matching string in a given sentence with a minimum rate for Levenshtein.
func LevenshteinContains(sentence, matching string, rate int) bool {
	words := strings.Split(sentence, " ")
	for _, word := range words {
		// Returns true if the distance is below the rate
		if LevenshteinDistance(word, matching) <= rate {
			return true
		}
	}

	return false
}
