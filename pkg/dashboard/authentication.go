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
	"crypto/rand"
	"fmt"
	"os"

	"github.com/bhojpur/suga/pkg/util"

	"github.com/gookit/color"
	"golang.org/x/crypto/bcrypt"
)

var fileName = "res/authentication.txt"

var authenticationHash []byte

// GenerateToken generates a random token of 30 characters and returns it
func GenerateToken() string {
	b := make([]byte, 30)
	rand.Read(b)

	fmt.Println("hey")
	return fmt.Sprintf("%x", b)
}

// HashToken gets the given tokens and returns its hash using bcrypt
func HashToken(token string) []byte {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(token), 14)
	return bytes
}

// ChecksToken checks if the given token is the good one from the authentication file
func ChecksToken(token string) bool {
	err := bcrypt.CompareHashAndPassword(authenticationHash, []byte(token))
	return err == nil
}

// AuthenticationFileExists checks if the authentication file exists and return the condition
func AuthenticationFileExists() bool {
	_, err := os.Open(fileName)
	return err == nil
}

// SaveHash saves the given hash to the authentication file
func SaveHash(hash string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.WriteString(hash)
}

// Authenticate checks if the authentication file exists and if not it generates the file with a new token
func Authenticate() {
	// Do nothing if the authentication file exists
	if AuthenticationFileExists() {
		authenticationHash = util.ReadFile(fileName)
		return
	}

	// Generates the token and gives it to the user
	token := GenerateToken()
	fmt.Printf("Your authentication token is: %s\n", color.FgLightGreen.Render(token))
	fmt.Println("Save it, you won't be able to get it again unless you generate a new one.")
	fmt.Println()

	// Hash the token and save it
	hash := HashToken(token)
	SaveHash(string(hash))

	authenticationHash = hash
}
