package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hello(t *testing.T) {
	h1 := hello("Andrey")
	assert.Equal(t, "Hello, Andrey!", h1)
	h2 := hello("World")
	assert.Equal(t, "Hello, World!", h2)
}
