package analysis

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
	"reflect"
	"testing"
)

func TestSentence_WordsBag(t *testing.T) {
	sentence := Sentence{"en", "Hi how are you"}
	words := Sentence{"en", "hi hello good morning are is were you seven"}.stem()

	wordsBag := sentence.WordsBag(words)
	excepted := []float64{0, 0, 0, 0, 1, 0}

	if !reflect.DeepEqual(excepted, wordsBag) {
		t.Errorf("sentence.WordsBag() failed, excepted %v, got %v", excepted, wordsBag)
	}
}

func TestSentence_Arrange(t *testing.T) {
	sentence := Sentence{"en", "Hello. how are you!   "}
	sentence.arrange()

	excepted := "Hello how are you"

	if sentence.Content != excepted {
		t.Errorf("sentence.Arrange() failed, excepted %v, got %v", excepted, sentence.Content)
	}
}

func TestSentence_Tokenize(t *testing.T) {
	sentence := Sentence{"en", "Hello How are you"}
	tokens := sentence.tokenize()

	excepted := []string{"hello", "how", "are", "you"}

	if !reflect.DeepEqual(tokens, excepted) {
		t.Errorf("sentence.Tokenize() failed, excepted %v, got %v", excepted, tokens)
	}
}
