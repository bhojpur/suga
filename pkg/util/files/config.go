package files

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
	"io/ioutil"
	"math/rand"
	"os"

	log "github.com/sirupsen/logrus"
)

// Configuration is the data required to start the tool
type Configuration struct {
	Port       string `json:"port"`
	Host       string `json:"host"`
	SSL        bool   `json:"ssl"`
	DebugLevel string `json:"debug_level"`
	BotName    string `json:"bot_name"`
	UserToken  string `json:"user_token"`
}

// FileExists checks if a file exists at a given path, it returns the condition
func FileExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}

// GenerateToken returns a random token of 50 characters
func GenerateToken() string {
	b := make([]byte, 50)
	rand.Read(b)

	return fmt.Sprintf("%x", b)
}

// SetupConfig initializes the config file if it does not exists and returns the config itself
func SetupConfig(fileName string) *Configuration {
	config := Configuration{
		Port:       "8080",
		SSL:        false,
		Host:       "localhost",
		DebugLevel: "error",
		BotName:    "Suga",
		UserToken:  GenerateToken(),
	}

	if FileExists(fileName) {
		// Read and parse the json file
		file, _ := ioutil.ReadFile(fileName)
		err := json.Unmarshal(file, &config)
		if err != nil {
			log.Fatal(err)
		}

		return &config
	}

	file, _ := json.MarshalIndent(config, "", " ")
	_ = ioutil.WriteFile(fileName, file, 0644)

	return &config
}
