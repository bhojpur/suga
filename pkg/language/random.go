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
	"errors"
	"regexp"
	"sort"
	"strconv"
)

var decimal = "\\b\\d+([\\.,]\\d+)?"

// FindRangeLimits finds the range for random numbers and returns a sorted integer array
func FindRangeLimits(local, entry string) ([]int, error) {
	decimalsRegex := regexp.MustCompile(decimal)
	limitStrArr := decimalsRegex.FindAllString(entry, 2)
	limitArr := make([]int, 0)

	if limitStrArr == nil {
		return make([]int, 0), errors.New("No range")
	}

	if len(limitStrArr) != 2 {
		return nil, errors.New("Need 2 numbers, a lower and upper limit")
	}

	for _, v := range limitStrArr {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, errors.New("Non integer range")
		}
		limitArr = append(limitArr, num)
	}

	sort.Ints(limitArr)
	return limitArr, nil
}
