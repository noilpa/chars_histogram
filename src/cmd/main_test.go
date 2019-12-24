package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calculateCharsFrequency(t *testing.T) {
	file := []byte("aa\n\n\n    3PP")
	freqMap := calculateCharsFrequency(file)
	assert.Equal(t, 2, freqMap["a"])
	assert.Equal(t, 3, freqMap["\n"])
	assert.Equal(t, 4, freqMap[" "])
	assert.Equal(t, 1, freqMap["3"])
	assert.Equal(t, 2, freqMap["P"])
}