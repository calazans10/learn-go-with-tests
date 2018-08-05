package iteration

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
	checkRepeat := func(t *testing.T, expected, actual string) {
		t.Helper()
		assert.Equal(t, expected, actual, "they should be equal")
	}

	t.Run("repeating the character two times", func(t *testing.T) {
		expected := "aa"
		actual := Repeat("a", 2)
		checkRepeat(t, actual, expected)
	})

	t.Run("repeating the character five times", func(t *testing.T) {
		expected := "aaaaa"
		actual := Repeat("a", 5)
		checkRepeat(t, actual, expected)
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	repeat := Repeat("a", 4)
	fmt.Println(repeat)
	// Output: aaaa
}
