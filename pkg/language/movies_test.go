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
	"reflect"
	"testing"
)

func TestSerializeMovies(t *testing.T) {
	movies := SerializeMovies()
	excepted := "Toy Story (1995)"

	if movies[0].Name != excepted {
		t.Errorf("SerializeMovies() failed, excepted %s got %s.", excepted, movies[0].Name)
	}
}

func TestSearchMovie(t *testing.T) {
	movie := SearchMovie("Adventure", "0")
	excepted := "2001: A Space Odyssey (1968)"

	if movie.Name != excepted {
		t.Errorf("SearchMovie() failed, excepted %s got %s.", excepted, movie.Name)
	}
}

func TestFindMoviesGenres(t *testing.T) {
	sentence := "I like movies of adventure, sci-fi"
	excepted := []string{"Adventure", "Sci-Fi"}
	genres := FindMoviesGenres("en", sentence)

	if !reflect.DeepEqual(excepted, genres) {
		t.Errorf("FindMoviesGenres() failed, excepted %s got %s.", excepted, genres)
	}
}
