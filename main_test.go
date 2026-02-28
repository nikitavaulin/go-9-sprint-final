package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	testCases := []struct {
		testName     string
		inputSize    int
		expectedSize int
	}{
		{"negative size", -5, 0},
		{"one element in slice", 0, 0},
		{"normal input size", 5, 5},
	}
	for _, testCase := range testCases {
		elements := generateRandomElements(testCase.inputSize)
		assert.Len(t, elements, testCase.expectedSize)
	}
}

func TestMaximum(t *testing.T) {
	testCases := []struct {
		testName       string
		inputData      []int
		expectedResult int
	}{
		{"nil slice", nil, 0},
		{"empty slice", []int{}, 0},
		{"one element in slice", []int{5}, 5},
		{"ordinary elements in slice", []int{5, 5, 5}, 5},
		{"max in first position", []int{7, 6, 2, 3, 6}, 7},
		{"max in middle", []int{5, 6, 2, 7, 6}, 7},
		{"max in last position", []int{5, 6, 2, 3, 7}, 7},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.expectedResult, maximum(testCase.inputData), testCase.testName)
	}
}
