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
	"testing"
)

func TestSearchReason(t *testing.T) {
	sentences := map[string]string{
		"Remind me to call mom":         "call mom",
		"Remind me to cook eggs":        "cook eggs",
		"Remind me that I have an exam": "I have an exam",
		"Remind me to wash the dishes":  "wash the dishes",
		"Remind me the conference call": "conference call",
	}

	for sentence, excepted := range sentences {
		reason := SearchReason("en", sentence)
		if reason != excepted {
			t.Errorf("SearchReason() failed, excepted %s got %s.", excepted, reason)
		}
	}
}

func BenchmarkSearchReason(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SearchReason("en", "Remind me to wash the dishes the 28th of march")
	}
}
