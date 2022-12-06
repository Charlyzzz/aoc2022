package day03

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const input = `
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`

func TestItemPriority(t *testing.T) {
	testCases := []struct {
		item     string
		expected int
	}{
		{"p", 16},
		{"L", 38},
		{"P", 42},
		{"v", 22},
		{"t", 20},
		{"s", 19},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("%s priority is %d", test.item, test.expected), func(t *testing.T) {
			assert.Equal(t, test.expected, itemPriority(test.item))
		})
	}
}

func TestPart1ShortInput(t *testing.T) {
	sum := sumDuplicatePriorities(input)
	assert.Equal(t, 157, sum)
}

func TestPart2ShortInput(t *testing.T) {
	sum := sumBadgesPriorities(input)
	assert.Equal(t, 70, sum)
}
