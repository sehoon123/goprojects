package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare1(t *testing.T) {
	assert.Equal(t, 81, square(9), "should be 81")
	assert.Equal(t, 9, square(3), "should be 9")
	assert.Equal(t, 0, square(0), "should be 0")
	assert.Equal(t, 4, square(-2), "should be 4")
}
