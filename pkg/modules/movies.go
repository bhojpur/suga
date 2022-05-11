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
	"math/rand"
	"strings"

	"github.com/bhojpur/suga/pkg/language"
	"github.com/bhojpur/suga/pkg/user"
	"github.com/bhojpur/suga/pkg/util"
)

var (
	// GenresTag is the intent tag for its module
	GenresTag = "movies genres"
	// MoviesTag is the intent tag for its module
	MoviesTag = "movies search"
	// MoviesAlreadyTag is the intent tag for its module
	MoviesAlreadyTag = "already seen movie"
	// MoviesDataTag is the intent tag for its module
	MoviesDataTag = "movies search from data"
)

// GenresReplacer gets the genre specified in the message and adds it to the user information.
// See modules/modules.go#Module.Replacer() for more details.
func GenresReplacer(locale, entry, response, token string) (string, string) {
	genres := language.FindMoviesGenres(locale, entry)

	// If there is no genres then reply with a message from res/datasets/messages.json
	if len(genres) == 0 {
		responseTag := "no genres"
		return responseTag, util.GetMessage(locale, responseTag)
	}

	// Change the user information to add the new genres
	user.ChangeUserInformation(token, func(information user.Information) user.Information {
		for _, genre := range genres {
			// Append the genre only is it isn't already in the information
			if util.Contains(information.MovieGenres, genre) {
				continue
			}

			information.MovieGenres = append(information.MovieGenres, genre)
		}
		return information
	})

	return GenresTag, response
}

// MovieSearchReplacer replaces the patterns contained inside the response by the movie's name
// and rating from the genre specified in the message.
// See modules/modules.go#Module.Replacer() for more details.
func MovieSearchReplacer(locale, entry, response, token string) (string, string) {
	genres := language.FindMoviesGenres(locale, entry)

	// If there is no genres then reply with a message from res/datasets/messages.json
	if len(genres) == 0 {
		responseTag := "no genres"
		return responseTag, util.GetMessage(locale, responseTag)
	}

	movie := language.SearchMovie(genres[0], token)

	return MoviesTag, fmt.Sprintf(response, movie.Name, movie.Rating)
}

// MovieSearchFromInformationReplacer replaces the patterns contained inside the response by the movie's name
// and rating from the genre in the user's information.
// See modules/modules.go#Module.Replacer() for more details.
func MovieSearchFromInformationReplacer(locale, _, response, token string) (string, string) {
	// If there is no genres then reply with a message from res/datasets/messages.json
	genres := user.GetUserInformation(token).MovieGenres
	if len(genres) == 0 {
		responseTag := "no genres saved"
		return responseTag, util.GetMessage(locale, responseTag)
	}

	movie := language.SearchMovie(genres[rand.Intn(len(genres))], token)
	genresJoined := strings.Join(genres, ", ")
	return MoviesDataTag, fmt.Sprintf(response, genresJoined, movie.Name, movie.Rating)
}
