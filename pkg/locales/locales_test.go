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

import "testing"

func TestGetNameByTag(t *testing.T) {
	name := "english"
	excepted := "en"
	tag := GetTagByName(name)

	if tag != excepted {
		t.Errorf("GetNameByTag() failed, excepted %s got %s.", excepted, tag)
	}
}

func TestGetTagByName(t *testing.T) {
	tag := "en"
	excepted := "english"
	name := GetNameByTag(tag)

	if name != excepted {
		t.Errorf("GetTagByName() failed, excepted %s got %s.", excepted, name)
	}
}

func TestExists(t *testing.T) {
	tag := "en"
	excepted := true
	exists := Exists(tag)

	if exists != excepted {
		t.Errorf("Exists() failed, excepted %t got %t.", excepted, exists)
	}
}
