package day02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const input = `
A Y
B X
C Z
`

func TestPart1ShortInput(t *testing.T) {
	score := strategyScore(input, true)
	assert.Equal(t, 15, score)
}

func TestPart2ShortInput(t *testing.T) {
	score := strategyScore(input, false)
	assert.Equal(t, 12, score)
}
