package start

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
	"github.com/bhojpur/suga/pkg/user"
	"github.com/bhojpur/suga/pkg/util"
	"strings"
	"time"
)

func init() {
	RegisterModule(Module{
		Action: CheckReminders,
	})
}

// CheckReminders will check the dates of the user's reminder and if they are outdated, remove them
func CheckReminders(token, locale string) {
	reminders := user.GetUserInformation(token).Reminders
	var messages []string

	// Iterate through the reminders to check if they are outdated
	for i, reminder := range reminders {
		date, _ := time.Parse("01/02/2006 03:04", reminder.Date)

		now := time.Now()
		// If the date is today
		if date.Year() == now.Year() && date.Day() == now.Day() && date.Month() == now.Month() {
			messages = append(messages, fmt.Sprintf("“%s”", reminder.Reason))

			// Removes the current reminder
			RemoveUserReminder(token, i)
		}
	}

	// Send the startup message
	if len(messages) != 0 {
		// If the message is already filled in return.
		if GetMessage() != "" {
			return
		}

		// Set the message with the current reminders
		listRemindersMessage := util.GetMessage(locale, "list reminders")
		if listRemindersMessage == "" {
			return
		}

		message := fmt.Sprintf(
			listRemindersMessage,
			user.GetUserInformation(token).Name,
			strings.Join(messages, ", "),
		)
		SetMessage(message)
	}
}

// RemoveUserReminder removes the reminder at a specific index in the user's information
func RemoveUserReminder(token string, index int) {
	user.ChangeUserInformation(token, func(information user.Information) user.Information {
		reminders := information.Reminders

		// Removes the element from the reminders slice
		if len(reminders) == 1 {
			reminders = []user.Reminder{}
		} else {
			reminders[index] = reminders[len(reminders)-1]
			reminders = reminders[:len(reminders)-1]
		}

		// Set the updated slice
		information.Reminders = reminders

		return information
	})
}
