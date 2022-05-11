package user

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
	"golang.org/x/oauth2"
)

// Information is the user's information retrieved from the client
type Information struct {
	Name           string        `json:"name"`
	MovieGenres    []string      `json:"movie_genres"`
	MovieBlacklist []string      `json:"movie_blacklist"`
	Reminders      []Reminder    `json:"reminders"`
	SpotifyToken   *oauth2.Token `json:"spotify_token"`
	SpotifyID      string        `json:"spotify_id"`
	SpotifySecret  string        `json:"spotify_secret"`
}

// A Reminder is something the user asked for Bhojpur Suga to remember
type Reminder struct {
	Reason string `json:"reason"`
	Date   string `json:"date"`
}

// userInformation is a map which is the cache for user information
var userInformation = map[string]Information{}

// ChangeUserInformation requires the token of the user and a function which gives the actual
// information and returns the new information.
func ChangeUserInformation(token string, changer func(Information) Information) {
	userInformation[token] = changer(userInformation[token])
}

// SetUserInformation sets the user's information by its token.
func SetUserInformation(token string, information Information) {
	userInformation[token] = information
}

// GetUserInformation returns the information of a user with his token
func GetUserInformation(token string) Information {
	return userInformation[token]
}
