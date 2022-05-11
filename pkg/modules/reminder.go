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
	"strings"

	"github.com/bhojpur/suga/pkg/language"
	"github.com/bhojpur/suga/pkg/language/date"
	"github.com/bhojpur/suga/pkg/user"
	"github.com/bhojpur/suga/pkg/util"
)

var (
	// ReminderSetterTag is the intent tag for its module
	ReminderSetterTag = "reminder setter"
	// ReminderGetterTag is the intent tag for its module
	ReminderGetterTag = "reminder getter"
)

// ReminderSetterReplacer replaces the pattern contained inside the response by the date of the reminder
// and its reason.
// See modules/modules.go#Module.Replacer() for more details.
func ReminderSetterReplacer(locale, entry, response, token string) (string, string) {
	// Search the time and
	sentence, date := date.SearchTime(locale, entry)
	reason := language.SearchReason(locale, sentence)

	// Format the date
	formattedDate := date.Format("01/02/2006 03:04")

	// Add the reminder inside the user's information
	user.ChangeUserInformation(token, func(information user.Information) user.Information {
		information.Reminders = append(information.Reminders, user.Reminder{
			Reason: reason,
			Date:   formattedDate,
		})

		return information
	})

	return ReminderSetterTag, fmt.Sprintf(response, reason, formattedDate)
}

// ReminderGetterReplacer gets the reminders in the user's information and replaces the pattern in the
// response patterns by the current reminders
// See modules/modules.go#Module.Replacer() for more details.
func ReminderGetterReplacer(locale, _, response, token string) (string, string) {
	reminders := user.GetUserInformation(token).Reminders
	var formattedReminders []string

	// Iterate through the reminders and parse them
	for _, reminder := range reminders {
		formattedReminder := fmt.Sprintf(
			util.GetMessage(locale, "reminder"),
			reminder.Reason,
			reminder.Date,
		)
		formattedReminders = append(formattedReminders, formattedReminder)
	}

	// If no reminder has been found
	if len(formattedReminders) == 0 {
		return ReminderGetterTag, util.GetMessage(locale, "no reminders")
	}

	return ReminderGetterTag, fmt.Sprintf(response, strings.Join(formattedReminders, " "))
}
