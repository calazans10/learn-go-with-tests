package integers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdder(t *testing.T) {
	expected := 4
	actual := Add(2, 2)

	assert.Equal(t, expected, actual, "they should be equal")
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
