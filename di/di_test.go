package di

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	expected := "Hello, Chris"
	actual := buffer.String()

	assert.Equal(t, expected, actual, "they should be equal")
}
