package day04

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const input = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func TestPart1ShortInput(t *testing.T) {
	sum := countAssignmentOverlaps(input, true)
	assert.Equal(t, 2, sum)
}

func TestPart2ShortInput(t *testing.T) {
	sum := countAssignmentOverlaps(input, false)
	assert.Equal(t, 4, sum)
}

func TestParseAssignment(t *testing.T) {
	expected := assignment{
		elf1: areas{from: 1, to: 3},
		elf2: areas{from: 6, to: 10},
	}
	assert.Equal(t, expected, parseAssignment("1-3,6-10"))
}
