package network

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
	"math/rand"
	"reflect"
	"testing"
)

func TestRandomMatrix(t *testing.T) {
	rand.Seed(0)
	random := RandomMatrix(2, 2)
	excepted := Matrix{
		{0.8903922985882329, -0.5100698294124405},
		{0.31191253039081035, -0.8913123208005992},
	}

	if !reflect.DeepEqual(random, excepted) {
		t.Errorf("RandomMatrix() failed, excepted %f got %f.", excepted, random)
	}
}

func TestCreateMatrix(t *testing.T) {
	matrix := CreateMatrix(2, 2)
	excepted := Matrix{
		{0, 0},
		{0, 0},
	}

	if !reflect.DeepEqual(matrix, excepted) {
		t.Errorf("CreateMatrix() failed, excepted %f got %f.", excepted, matrix)
	}
}

func TestApplyFunction(t *testing.T) {
	a := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}

	ApplyFunction(a, func(x float64) float64 {
		return x + 1
	})

	// Excepted value
	r := Matrix{
		{2, 3, 4},
		{5, 6, 7},
	}

	if !reflect.DeepEqual(a, r) {
		t.Errorf("ApplyFunction(fn(x)=x+1) failed, excepted %v, got %v", r, a)
	}
}

func TestApplyRate(t *testing.T) {
	a := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}

	ApplyRate(a, 0.5)

	excepted := Matrix{
		{0.5, 1, 1.5},
		{2, 2.5, 3},
	}

	if !reflect.DeepEqual(a, excepted) {
		t.Errorf("ApplyRate() failed, excepted %f got %f.", excepted, a)
	}
}

func TestSum(t *testing.T) {
	a := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}

	b := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}

	Sum(a, b)

	// Excepted value
	r := Matrix{
		{2, 4, 6},
		{8, 10, 12},
	}

	if !reflect.DeepEqual(a, r) {
		t.Errorf("Sum(%v) failed, excepted %v, got %v", b, r, a)
	}
}

func TestDifference(t *testing.T) {
	a := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}

	b := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}

	a = Difference(a, b)

	// Excepted value
	r := Matrix{
		{0, 0, 0},
		{0, 0, 0},
	}

	if !reflect.DeepEqual(a, r) {
		t.Errorf("Difference(%v) failed, excepted %v, got %v", b, r, a)
	}
}

func TestMultiplication(t *testing.T) {
	a := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}

	b := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}

	a = Multiplication(a, b)

	excepted := Matrix{
		{1, 4, 9},
		{16, 25, 36},
	}

	if !reflect.DeepEqual(a, excepted) {
		t.Errorf("Multiplication() failed, excepted %f got %f.", excepted, a)
	}
}

func TestDotProduct(t *testing.T) {
	a := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}

	b := Matrix{
		{7, 8},
		{9, 10},
		{11, 12},
	}

	// Actual value
	p := DotProduct(a, b)

	// Excepted value
	r := Matrix{
		{58, 64},
		{139, 154},
	}

	if !reflect.DeepEqual(p, r) {
		t.Errorf("DotProduct(%v) failed, excepted %v, got %v", b, r, p)
	}
}

func TestTranspose(t *testing.T) {
	a := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}

	a = Transpose(a)

	r := Matrix{
		{1, 4},
		{2, 5},
		{3, 6},
	}

	if !reflect.DeepEqual(a, r) {
		t.Errorf("Transpose failed, excepted %v, got %v", r, a)
	}
}
