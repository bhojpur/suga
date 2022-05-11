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
	"regexp"

	"github.com/bhojpur/suga/pkg/language"
	"github.com/bhojpur/suga/pkg/util"
	"github.com/soudy/mathcat"
)

// MathTag is the intent tag for its module
var MathTag = "math"

// MathReplacer replaces the pattern contained inside the response by the answer of the math
// expression specified in the message.
// See modules/modules.go#Module.Replacer() for more details.
func MathReplacer(locale, entry, response, _ string) (string, string) {
	operation := language.FindMathOperation(entry)

	// If there is no operation in the entry message reply with a "don't understand" message
	if operation == "" {
		responseTag := "don't understand"
		return responseTag, util.GetMessage(locale, responseTag)
	}

	res, err := mathcat.Eval(operation)
	// If the expression isn't valid reply with a message from res/datasets/messages.json
	if err != nil {
		responseTag := "math not valid"
		return responseTag, util.GetMessage(locale, responseTag)
	}
	// Use number of decimals from the query
	decimals := language.FindNumberOfDecimals(locale, entry)
	if decimals == 0 {
		decimals = 6
	}

	result := res.FloatString(decimals)

	// Remove trailing zeros of the result with a Regex
	trailingZerosRegex := regexp.MustCompile(`\.?0+$`)
	result = trailingZerosRegex.ReplaceAllString(result, "")

	return MathTag, fmt.Sprintf(response, result)
}
