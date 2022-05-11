package date

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

import "testing"

func TestDeleteDates(t *testing.T) {
	sentences := map[string]string{
		"Remind me to call mom tomorrow":                 "Remind me to call mom",
		"Remind me to cook eggs after tomorrow":          "Remind me to cook eggs",
		"Remind me that I have an exam saturday":         "Remind me that I have an exam",
		"Remind me to wash the dishes the 28th of march": "Remind me to wash the dishes",
		"Remind me the conference call of the 04/12":     "Remind me the conference call",
	}

	for sentence, excepted := range sentences {
		deleteDatesSentence := DeleteDates("en", sentence)

		if excepted != deleteDatesSentence {
			t.Errorf("DeleteDates() failed, excepted %s got %s.", excepted, deleteDatesSentence)
		}
	}
}

func TestDeleteTimes(t *testing.T) {
	sentences := map[string]string{
		"Remind me to call mom tomorrow at 9:30pm":       "Remind me to call mom tomorrow",
		"Remind me to cook eggs after tomorrow at 12 am": "Remind me to cook eggs after tomorrow",
	}

	for sentence, excepted := range sentences {
		deleteTimesSentence := DeleteTimes("en", sentence)

		if excepted != deleteTimesSentence {
			t.Errorf("DeleteTimes() failed, excepted %s got %s.", excepted, deleteTimesSentence)
		}
	}
}
