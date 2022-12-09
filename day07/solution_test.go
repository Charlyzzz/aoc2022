package day07

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const input = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`

func TestPart1ShortInput(t *testing.T) {
	res := sumDirectoriesWithSizeUnder(input)
	assert.Equal(t, 95437, res)
}

func TestPart2ShortInput(t *testing.T) {
	res := smallestDirectorySizeToDelete(input)
	assert.Equal(t, 24933642, res)
}
