package day06

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type tc struct {
	buffer   string
	expected int
}

func TestPart1ShortInput(t *testing.T) {
	var testCases = []tc{
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%s marker is %d", testCase.buffer, testCase.expected), func(t *testing.T) {
			marker := signalMarker(testCase.buffer, 4)
			assert.Equal(t, testCase.expected, marker)
		})
	}
}

func TestPart2ShortInput(t *testing.T) {
	var testCases = []tc{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%s marker is %d", testCase.buffer, testCase.expected), func(t *testing.T) {
			marker := signalMarker(testCase.buffer, 14)
			assert.Equal(t, testCase.expected, marker)
		})
	}
}

func TestAppendAndTrim(t *testing.T) {
	assert.ElementsMatch(t, []int{3, 4, 5}, appendAndTrim(3, 5, []int{1, 2, 3, 4}))
	assert.ElementsMatch(t, []int{1, 2, 3}, appendAndTrim(3, 3, []int{1, 2}))
}
