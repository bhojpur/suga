package locales

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
	// Import these packages to trigger the init() function
	_ "github.com/bhojpur/suga/res/locales/ca"
	_ "github.com/bhojpur/suga/res/locales/de"
	_ "github.com/bhojpur/suga/res/locales/el"
	_ "github.com/bhojpur/suga/res/locales/en"
	_ "github.com/bhojpur/suga/res/locales/es"
	_ "github.com/bhojpur/suga/res/locales/fr"
	_ "github.com/bhojpur/suga/res/locales/it"
	_ "github.com/bhojpur/suga/res/locales/nl"
	_ "github.com/bhojpur/suga/res/locales/tr"
)

// Locales is the list of locales's tags and names
// Please check if the language is supported in https://github.com/tebeka/snowball,
// if it is please add the correct language name.
var Locales = []Locale{
	{
		Tag:  "en",
		Name: "english",
	},
	{
		Tag:  "de",
		Name: "german",
	},
	{
		Tag:  "fr",
		Name: "french",
	},
	{
		Tag:  "es",
		Name: "spanish",
	},
	{
		Tag:  "ca",
		Name: "catalan",
	},
	{
		Tag:  "it",
		Name: "italian",
	},
	{
		Tag:  "tr",
		Name: "turkish",
	},
	{
		Tag:  "nl",
		Name: "dutch",
	},
	{
		Tag:  "el",
		Name: "greek",
	},
}

// A Locale is a registered locale in the file
type Locale struct {
	Tag  string
	Name string
}

// GetNameByTag returns the name of the given locale's tag
func GetNameByTag(tag string) string {
	for _, locale := range Locales {
		if locale.Tag != tag {
			continue
		}

		return locale.Name
	}

	return ""
}

// GetTagByName returns the tag of the given locale's name
func GetTagByName(name string) string {
	for _, locale := range Locales {
		if locale.Name != name {
			continue
		}

		return locale.Tag
	}

	return ""
}

// Exists checks if the given tag exists in the list of locales
func Exists(tag string) bool {
	for _, locale := range Locales {
		if locale.Tag == tag {
			return true
		}
	}

	return false
}
