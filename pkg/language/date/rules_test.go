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

import (
	"testing"
	"time"
)

// CheckEquality checks if the two given dates are the same
func CheckEquality(a, b time.Time) bool {
	return a.Day() == b.Day() || a.Year() == b.Year() || a.Month() == b.Month()
}

func TestRuleToday(t *testing.T) {
	sentence := "Remind me that I have an exam today"
	date := RuleToday("en", sentence)

	if !CheckEquality(time.Now(), RuleToday("en", sentence)) {
		t.Errorf("RuleToday() failed, excepted %s got %s.", time.Now(), date)
	}
}

func TestRuleTomorrow(t *testing.T) {
	day := time.Hour * 24

	sentences := map[string]time.Time{
		"Remind me to call mom tomorrow":       time.Now().Add(day),
		"Remind me to call mom after tomorrow": time.Now().Add(day * 2),
	}

	for sentence, date := range sentences {
		foundDate := RuleTomorrow("en", sentence)
		if !CheckEquality(date, foundDate) {
			t.Errorf("SearchTime() failed, excepted %s got %s.", date, foundDate)
		}
	}
}

func TestRuleDayOfWeek(t *testing.T) {
	sentence := "Remind me that I have an exam saturday"
	excepted := 6
	weekday := int(RuleDayOfWeek("en", sentence).Weekday())

	if excepted != weekday {
		t.Errorf("RuleDayOfWeek() failed, excepted %d got %d.", excepted, weekday)
	}
}

func TestRuleNaturalDate(t *testing.T) {
	sentence := "Nothing here"
	date := RuleNaturalDate("en", sentence)
	excepted := time.Time{}

	if date != excepted {
		t.Errorf("RuleNaturalDate() failed, excepted %s got %s.", excepted, date)
	}

	sentence = "Remind me that I have an exam the 28th of march"
	date = RuleNaturalDate("en", sentence)

	if date.Month() != 3 || date.Day() != 28 {
		t.Errorf("RuleNaturalDate() failed, excepted 3/28 got %s.", date)
	}

	sentence = "Remind me that I have an exam in december"
	date = RuleNaturalDate("en", sentence)

	if date.Month() != 12 || date.Day() != 1 {
		t.Errorf("RuleNaturalDate() failed, excepted 1/12 got %s.", date)
	}
}

func TestRuleDate(t *testing.T) {
	sentence := "Remind me that I have an exam the 12/04"
	date := RuleDate("en", sentence)

	if date.Day() != 4 || date.Month() != 12 {
		t.Errorf("RuleData() failed, excepted 04/12 got %s.", date)
	}
}

func TestRuleTime(t *testing.T) {
	sentence := "Remind me that I have an exam the 12/04 at 9:30 pm"
	time := RuleTime(sentence)

	if time.Hour() != 21 || time.Minute() != 30 {
		t.Errorf("RuleTime() failed, excepted 21:30 got %s.", time)
	}
}
