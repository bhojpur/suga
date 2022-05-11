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
	"encoding/json"
	"fmt"
	"strings"

	"github.com/bhojpur/suga/pkg/util"
)

// Country is the serializer of the countries.json file in the res folder
type Country struct {
	Name     map[string]string `json:"name"`
	Capital  string            `json:"capital"`
	Code     string            `json:"code"`
	Area     float64           `json:"area"`
	Currency string            `json:"currency"`
}

var countries = SerializeCountries()

// SerializeCountries returns a list of countries, serialized from `res/datasets/countries.json`
func SerializeCountries() (countries []Country) {
	err := json.Unmarshal(util.ReadFile("res/datasets/countries.json"), &countries)
	if err != nil {
		fmt.Println(err)
	}

	return countries
}

// FindCountry returns the country found in the sentence and if no country is found, returns an empty Country struct
func FindCountry(locale, sentence string) Country {
	for _, country := range countries {
		name, exists := country.Name[locale]

		if !exists {
			continue
		}

		// If the actual country isn't contained in the sentence, continue
		if !strings.Contains(strings.ToLower(sentence), strings.ToLower(name)) {
			continue
		}

		// Returns the right country
		return country
	}

	// Returns an empty country if none has been found
	return Country{}
}
