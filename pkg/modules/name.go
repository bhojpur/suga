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
	"github.com/bhojpur/suga/pkg/user"
	"github.com/bhojpur/suga/pkg/util"
)

var (
	// NameGetterTag is the intent tag for its module
	NameGetterTag = "name getter"
	// NameSetterTag is the intent tag for its module
	NameSetterTag = "name setter"
)

// NameGetterReplacer replaces the pattern contained inside the response by the user's name.
// See modules/modules.go#Module.Replacer() for more details.
func NameGetterReplacer(locale, _, response, token string) (string, string) {
	name := user.GetUserInformation(token).Name

	if strings.TrimSpace(name) == "" {
		responseTag := "don't know name"
		return responseTag, util.GetMessage(locale, responseTag)
	}

	return NameGetterTag, fmt.Sprintf(response, name)
}

// NameSetterReplacer gets the name specified in the message and save it in the user's information.
// See modules/modules.go#Module.Replacer() for more details.
func NameSetterReplacer(locale, entry, response, token string) (string, string) {
	name := language.FindName(entry)

	// If there is no name in the entry string
	if name == "" {
		responseTag := "no name"
		return responseTag, util.GetMessage(locale, responseTag)
	}

	// Capitalize the name
	name = strings.Title(name)

	// Change the name inside the user information
	user.ChangeUserInformation(token, func(information user.Information) user.Information {
		information.Name = name
		return information
	})

	return NameSetterTag, fmt.Sprintf(response, name)
}
