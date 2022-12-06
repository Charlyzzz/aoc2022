package day05

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const input = `
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

func TestPart1ShortInput(t *testing.T) {
	crates := getTopStacks(input[1:], true)
	assert.Equal(t, "CMZ", crates)
}

func TestPart2ShortInput(t *testing.T) {
	crates := getTopStacks(input[1:], false)
	assert.Equal(t, "MCD", crates)
}
