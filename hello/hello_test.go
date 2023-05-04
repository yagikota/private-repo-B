package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello1(t *testing.T) {
	expected := "func hello1: arg=world"
	result := Hello1("world")
	assert.Equal(t, expected, result)
}

func TestHello2(t *testing.T) {
	expected := "func hello2: arg=world"
	result := Hello2("world")
	assert.Equal(t, expected, result)
}
