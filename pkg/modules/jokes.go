package modules

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
	"io/ioutil"
	"net/http"

	"github.com/bhojpur/suga/pkg/util"
)

const jokeURL = "https://official-joke-api.appspot.com/random_joke"

// JokesTag is the intent tag for its module
var JokesTag = "jokes"

// Joke represents the response from the joke api
type Joke struct {
	ID        int64  `json:"id"`
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

// JokesReplacer replaces the pattern contained inside the response by a random joke from the api
// specified in jokeURL.
// See modules/modules.go#Module.Replacer() for more details.
func JokesReplacer(locale, entry, response, _ string) (string, string) {

	resp, err := http.Get(jokeURL)
	if err != nil {
		responseTag := "no jokes"
		return responseTag, util.GetMessage(locale, responseTag)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		responseTag := "no jokes"
		return responseTag, util.GetMessage(locale, responseTag)
	}

	joke := &Joke{}

	err = json.Unmarshal(body, joke)
	if err != nil {
		responseTag := "no jokes"
		return responseTag, util.GetMessage(locale, responseTag)
	}

	jokeStr := joke.Setup + " " + joke.Punchline

	return JokesTag, fmt.Sprintf(response, jokeStr)
}
