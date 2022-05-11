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
	"fmt"
	"math/rand"
	"strconv"

	"github.com/bhojpur/suga/pkg/language"
	"github.com/bhojpur/suga/pkg/util"
)

// RandomTag is the intent tag for its module
var RandomTag = "random number"

// RandomNumberReplacer replaces the pattern contained inside the response by a random number.
// See modules/modules.go#Module.Replacer() for more details.
func RandomNumberReplacer(locale, entry, response, _ string) (string, string) {
	limitArr, err := language.FindRangeLimits(locale, entry)
	if err != nil {
		if limitArr != nil {
			return RandomTag, fmt.Sprintf(response, strconv.Itoa(rand.Intn(100)))
		}

		responseTag := "no random range"
		return responseTag, util.GetMessage(locale, responseTag)
	}

	min := limitArr[0]
	max := limitArr[1]
	randNum := rand.Intn((max - min)) + min
	return RandomTag, fmt.Sprintf(response, strconv.Itoa(randNum))
}
