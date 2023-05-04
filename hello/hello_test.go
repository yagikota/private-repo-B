package hello

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHello(t *testing.T) {
	result := hello("world")
	require.Equal(t, "helloworld", result)
}
