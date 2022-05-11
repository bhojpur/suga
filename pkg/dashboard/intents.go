package dashboard

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
	"net/http"
	"os"

	"github.com/bhojpur/suga/pkg/analysis"
	"github.com/bhojpur/suga/pkg/util"
	"github.com/gookit/color"
	"github.com/gorilla/mux"
)

// An Error is what the api replies when an error occurs
type Error struct {
	Message string `json:"message"`
}

// DeleteRequest is for the parameters required to delete an intent via the REST Api
type DeleteRequest struct {
	Tag string `json:"tag"`
}

// WriteIntents writes the given intents to the intents file
func WriteIntents(locale string, intents []analysis.Intent) {
	analysis.CacheIntents(locale, intents)

	// Encode the json
	bytes, _ := json.MarshalIndent(intents, "", "  ")

	// Write it to the file
	file, err := os.Create("res/locales/" + locale + "/intents.json")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.Write(bytes)
}

// AddIntent adds the given intent to the intents file
func AddIntent(locale string, intent analysis.Intent) {
	intents := append(analysis.SerializeIntents(locale), intent)

	WriteIntents(locale, intents)

	fmt.Printf("Added %s intent.\n", color.FgMagenta.Render(intent.Tag))
}

// RemoveIntent removes the intent with the given tag from the intents file
func RemoveIntent(locale, tag string) {
	intents := analysis.SerializeIntents(locale)

	// Iterate through the intents to remove the right one
	for i, intent := range intents {
		if intent.Tag != tag {
			continue
		}

		intents[i] = intents[len(intents)-1]
		intents = intents[:len(intents)-1]
		fmt.Printf("The intent %s was deleted.\n", color.FgMagenta.Render(intent.Tag))
	}

	WriteIntents(locale, intents)
}

// GetIntents is the route to get the intents
func GetIntents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	data := mux.Vars(r)

	// Encode the intents for the given locale
	json.NewEncoder(w).Encode(analysis.GetIntents(data["locale"]))
}

// CreateIntent is the route to create a new intent
func CreateIntent(w http.ResponseWriter, r *http.Request) {
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,Bhojpur-Token"
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")

	data := mux.Vars(r)

	// Checks if the token present in the headers is the right one
	token := r.Header.Get("Bhojpur-Token")
	if !ChecksToken(token) {
		json.NewEncoder(w).Encode(Error{
			Message: util.GetMessage(data["locale"], "no permission"),
		})
		return
	}

	// Decode request json body
	var intent analysis.Intent
	json.NewDecoder(r.Body).Decode(&intent)

	if intent.Responses == nil || intent.Patterns == nil {
		json.NewEncoder(w).Encode(Error{
			Message: util.GetMessage(data["locale"], "patterns same"),
		})
		return
	}

	// Returns an error if the tags are the same
	for _, _intent := range analysis.GetIntents(data["locale"]) {
		if _intent.Tag == intent.Tag {
			json.NewEncoder(w).Encode(Error{
				Message: util.GetMessage(data["locale"], "tags same"),
			})
			return
		}
	}

	// Adds the intent
	AddIntent(data["locale"], intent)

	json.NewEncoder(w).Encode(intent)
}

// DeleteIntent is the route used to delete an intent
func DeleteIntent(w http.ResponseWriter, r *http.Request) {
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,Bhojpur-Token"
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")

	data := mux.Vars(r)

	// Checks if the token present in the headers is the right one
	token := r.Header.Get("Bhojpur-Token")
	if !ChecksToken(token) {
		json.NewEncoder(w).Encode(Error{
			Message: util.GetMessage(data["locale"], "no permission"),
		})
		return
	}

	var deleteRequest DeleteRequest
	json.NewDecoder(r.Body).Decode(&deleteRequest)

	RemoveIntent(data["locale"], deleteRequest.Tag)

	json.NewEncoder(w).Encode(analysis.GetIntents(data["locale"]))
}
