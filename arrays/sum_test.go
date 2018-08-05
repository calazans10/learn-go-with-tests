package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	expected := 15
	actual := Sum(numbers)

	assert.Equal(t, expected, actual, "they should be equal")
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, expected, actual []int) {
		assert.Equal(t, expected, actual, "they should be equal")
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		expected := []int{2, 9}
		actual := SumAllTails([]int{1, 2}, []int{0, 9})
		checkSums(t, expected, actual)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		expected := []int{0, 9}
		actual := SumAllTails([]int{}, []int{3, 4, 5})
		checkSums(t, expected, actual)
	})
}
