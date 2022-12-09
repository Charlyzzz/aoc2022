package day08

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const input = `30373
25512
65332
33549
35390`

func TestPart1ShortInput(t *testing.T) {
	res := visibleTrees(input)
	assert.Equal(t, 21, res)
}

func TestPart2ShortInput(t *testing.T) {
	res := highestScenicScore(input)
	assert.Equal(t, 8, res)
}

func TestScenicScore(t *testing.T) {
	m := makeHeightMap(input)
	cases := []struct {
		row      int
		col      int
		expected int
	}{
		{1, 2, 4},
		{3, 2, 8},
		{0, 1, 0},
		{4, 1, 0},
		{1, 0, 0},
		{1, 4, 0},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("%d-%d scores %d", tc.row, tc.col, tc.expected), func(t *testing.T) {
			h := m[tc.row][tc.col]
			score := scenicScore(m, tc.row, tc.col, h)
			assert.Equal(t, tc.expected, score)
		})
	}
}
