package day01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const input = `
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

func TestPart1ShortInput(t *testing.T) {
	res := sumTopNthCaloriesCarried(input, 1)
	assert.Equal(t, 24000, res)
}

func TestPart2ShortInput(t *testing.T) {
	res := sumTopNthCaloriesCarried(input, 3)
	assert.Equal(t, 45000, res)
}
