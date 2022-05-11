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
	"regexp"
	"strconv"
	"strings"
)

// MathDecimals is the map for having the regex on decimals in different languages
var MathDecimals = map[string]string{
	"en": `(\d+( |-)decimal(s)?)|(number (of )?decimal(s)? (is )?\d+)`,
	"de": `(\d+( |-)decimal(s)?)|(nummer (von )?decimal(s)? (ist )?\d+)`,
	"fr": `(\d+( |-)decimale(s)?)|(nombre (de )?decimale(s)? (est )?\d+)`,
	"es": `(\d+( |-)decimale(s)?)|(numero (de )?decimale(s)? (de )?\d+)`,
	"ca": `(\d+( |-)decimal(s)?)|(nombre (de )?decimal(s)? (de )?\d+)`,
	"it": `(\d+( |-)decimale(s)?)|(numero (di )?decimale(s)? (è )?\d+)`,
	"tr": `(\d+( |-)desimal(s)?)|(numara (dan )?desimal(s)? (mı )?\d+)`,
	"nl": `(\d+( |-)decimal(en)?)|(nummer (van )?decimal(en)? (is )?\d+)`,
	"el": `(\d+( |-)δεκαδικ(ό|ά)?)|(αριθμός (από )?δεκαδικ(ό|ά)? (είναι )?\d+)`,
}

// FindMathOperation finds a math operation in a string an returns it
func FindMathOperation(entry string) string {
	mathRegex := regexp.MustCompile(
		`((\()?(((\d+|pi)(\^\d+|!|.)?)|sqrt|cos|sin|tan|acos|asin|atan|log|ln|abs)( )?[+*\/\-x]?( )?(\))?[+*\/\-]?)+`,
	)

	operation := mathRegex.FindString(entry)
	// Replace "x" symbol by "*"
	operation = strings.Replace(operation, "x", "*", -1)
	return strings.TrimSpace(operation)
}

// FindNumberOfDecimals finds the number of decimals asked in the query
func FindNumberOfDecimals(locale, entry string) int {
	decimalsRegex := regexp.MustCompile(
		MathDecimals[locale],
	)
	numberRegex := regexp.MustCompile(`\d+`)

	decimals := numberRegex.FindString(decimalsRegex.FindString(entry))
	decimalsInt, _ := strconv.Atoi(decimals)

	return decimalsInt
}
