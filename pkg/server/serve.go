package server

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
	"time"

	"github.com/bhojpur/suga/pkg/analysis"
	"github.com/bhojpur/suga/pkg/dashboard"
	"github.com/bhojpur/suga/pkg/modules/spotify"
	"github.com/bhojpur/suga/pkg/network"
	"github.com/bhojpur/suga/pkg/training"
	stamping "github.com/bhojpur/suga/pkg/version"
	"github.com/gookit/color"
	"github.com/gorilla/mux"
	gocache "github.com/patrickmn/go-cache"
)

var (
	// Create the neural network variable to use it everywhere
	neuralNetworks map[string]network.Network
	// Initializes the cache with a 5 minute lifetime
	cache = gocache.New(5*time.Minute, 5*time.Minute)
)

// Serve serves the server in the given port
func Serve(_neuralNetworks map[string]network.Network, port string) {
	// Set the current global network as a global variable
	neuralNetworks = _neuralNetworks

	// Initializes the router
	router := mux.NewRouter()
	router.HandleFunc("/callback", spotify.CompleteAuth)
	// Serve the websocket
	router.HandleFunc("/websocket", SocketHandle)
	// Serve the API
	router.HandleFunc("/api/{locale}/dashboard", GetDashboardData).Methods("GET")
	router.HandleFunc("/api/{locale}/intent", dashboard.CreateIntent).Methods("POST")
	router.HandleFunc("/api/{locale}/intent", dashboard.DeleteIntent).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/{locale}/train", Train).Methods("POST")
	router.HandleFunc("/api/{locale}/intents", dashboard.GetIntents).Methods("GET")
	router.HandleFunc("/api/coverage", analysis.GetCoverage).Methods("GET")

	magenta := color.FgMagenta.Render
	fmt.Printf("\nBhojpur Suga server %s listening on port %s...\n", stamping.FullVersion(), magenta(port))

	// Serves the chat
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}

// Train is the route to re-train the neural network
func Train(w http.ResponseWriter, r *http.Request) {
	// Checks if the token present in the headers is the right one
	token := r.Header.Get("Bhojpur-Token")
	if !dashboard.ChecksToken(token) {
		json.NewEncoder(w).Encode(dashboard.Error{
			Message: "You don't have the permission to do this.",
		})
		return
	}

	magenta := color.FgMagenta.Render
	fmt.Printf("\nRe-training the %s..\n", magenta("neural network"))

	for locale := range neuralNetworks {
		neuralNetworks[locale] = training.CreateNeuralNetwork(locale, true)
	}
}
