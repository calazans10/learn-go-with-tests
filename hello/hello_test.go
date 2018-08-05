package hello

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	checkMessage := func(t *testing.T, expected, actual string) {
		t.Helper()
		assert.Equal(t, expected, actual, "they should be equal")
	}

	t.Run("saying hello to people", func(t *testing.T) {
		expected := "Hello, Jason"
		actual := Hello("Jason", "")
		checkMessage(t, expected, actual)
	})

	t.Run("saying hello world when an empty string is supplied", func(t *testing.T) {
		expected := "Hello, World"
		actual := Hello("", "")
		checkMessage(t, expected, actual)
	})

	t.Run("saying hello to people in Spanish", func(t *testing.T) {
		expected := "Hola, Julieta"
		actual := Hello("Julieta", "Spanish")
		checkMessage(t, expected, actual)
	})

	t.Run("saying hello to people in French", func(t *testing.T) {
		expected := "Bonjour, Brigitte"
		actual := Hello("Brigitte", "French")
		checkMessage(t, expected, actual)
	})

	t.Run("saying hello to people in Portuguese", func(t *testing.T) {
		expected := "Olá, Jeferson"
		actual := Hello("Jeferson", "Portuguese")
		checkMessage(t, expected, actual)
	})
}

func ExampleHello() {
	hello := Hello("Joane", "")
	fmt.Println(hello)
	// Output: Hello, Joane
}
