package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvery(t *testing.T) {
	actual := Every(3, []int{1, 2, 3, 4, 5, 6, 7})
	expected := [][]int{{1, 2, 3}, {4, 5, 6}, {7}}
	assert.ElementsMatch(t, expected, actual)
}
