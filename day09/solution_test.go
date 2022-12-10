package day09

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const input = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

func TestPart1ShortInput(t *testing.T) {
	res := countTailPositions(input, 2)
	assert.Equal(t, 13, res)
}

func TestPart2ShortInput(t *testing.T) {
	res := countTailPositions(input, 10)
	assert.Equal(t, 1, res)
}

func TestPart2MidInput(t *testing.T) {
	midInput := `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`
	res := countTailPositions(midInput, 10)
	assert.Equal(t, 36, res)
}

func TestFollow(t *testing.T) {
	testCases := []struct {
		head      position
		tail      position
		expected  position
		direction string
	}{
		{position{0, 0}, position{0, 0}, position{0, 0}, "same place"},
		{position{2, 0}, position{0, 0}, position{1, 0}, "right"},
		{position{-2, 0}, position{0, 0}, position{-1, 0}, "left"},
		{position{0, 2}, position{0, 0}, position{0, 1}, "top"},
		{position{0, -2}, position{0, 0}, position{0, -1}, "bottom"},
		//    H
		//   (T)
		// T
		{position{1, 2}, position{0, 0}, position{1, 1}, "top right"},
		// H
		//(T)
		//     T
		{position{0, 2}, position{1, 0}, position{0, 1}, "top left"},
		// T
		//   (T)
		//    H
		{position{1, 0}, position{0, 2}, position{1, 1}, "bottom right"},
		//    T
		//(T)
		// H
		{position{0, 0}, position{1, 2}, position{0, 1}, "bottom left"},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("tail moves to %s", testCase.direction), func(t *testing.T) {
			head := testCase.head
			tail := testCase.tail
			follow(&head, &tail)
			assert.Equal(t, testCase.expected, tail)
		})
	}
}
