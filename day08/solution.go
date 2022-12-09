package day08

import (
	"aoc2022/challenge"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type c struct{}

func (c) Part1() interface{} {
	input := challenge.Input()
	return visibleTrees(input)
}

func (c) Part2() interface{} {
	input := challenge.Input()
	return highestScenicScore(input)
}

func (c) Day() int {
	return 8
}

func Challenge() challenge.Challenge {
	return c{}
}

type visitMap = map[string]interface{}

var visited = struct{}{}

func makeKey(row int, col int) string {
	return fmt.Sprintf("%d-%d", row, col)
}

func markVisited(m *visitMap, row int, col int) {
	key := makeKey(row, col)
	(*m)[key] = visited
}

func visibleTrees(input string) int {
	visitedMap := make(visitMap, 0)
	m := makeHeightMap(input)
	rows := len(m)
	cols := len(m[0])

	// north-south
	for col := 0; col < cols; col++ {
		trackVisibility(m, down, 0, col, visitedMap)
	}

	// west-east
	for col := 0; col < cols; col++ {
		trackVisibility(m, right, 0, col, visitedMap)
	}

	// east-west
	for row := 0; row < rows; row++ {
		trackVisibility(m, left, row, cols, visitedMap)
	}

	// south-north
	for col := 0; col < cols; col++ {
		trackVisibility(m, up, rows, col, visitedMap)
	}

	return len(visitedMap)
}

func trackVisibility(m heightMap, dir direction, row int, col int, visitedMap visitMap) {
	tallestTree := math.MinInt
	walk(m, dir, row, col, func(r int, c int) bool {
		height := m[r][c]
		if height > tallestTree {
			tallestTree = height
			markVisited(&visitedMap, r, c)
		}
		return true
	})
}

func highestScenicScore(input string) int {
	heightMap := makeHeightMap(input)
	highestScore := math.MinInt
	for rowIdx, row := range heightMap {
		for colIdx, height := range row {
			score := scenicScore(heightMap, rowIdx, colIdx, height)
			if score > highestScore {
				highestScore = score
			}
		}
	}
	return highestScore
}

type heightMap = [][]int

var directions = []direction{up, down, left, right}

func scenicScore(m heightMap, row int, col int, height int) int {
	score := 1
	for _, dir := range directions {
		dirScore := 0
		walk(m, dir, row, col, func(r int, c int) bool {
			tree := m[r][c]
			dirScore += 1
			return tree < height
		})
		score *= dirScore
	}
	return score
}

type direction = func(row int, col int) (int, int)

func up(row int, col int) (int, int) {
	return row - 1, col
}

func down(row int, col int) (int, int) {
	return row + 1, col
}

func right(row int, col int) (int, int) {
	return row, col + 1
}

func left(row int, col int) (int, int) {
	return row, col - 1
}

func walk(m heightMap, dir direction, row int, col int, onEach func(row int, col int) bool) {
	keepGoing := true
	rowBoundary := len(m) - 1
	colBoundary := len(m[0]) - 1
	for keepGoing {
		row, col = dir(row, col)
		if outOfBounds(row, rowBoundary) || outOfBounds(col, colBoundary) {
			return
		}
		keepGoing = onEach(row, col)
	}
}

func outOfBounds(position int, boundary int) bool {
	return position == -1 || position > boundary
}

func makeHeightMap(input string) heightMap {
	heightMap := make(heightMap, 0)
	for _, rowHeight := range strings.Split(input, "\n") {
		var row []int
		for _, heightReading := range rowHeight {
			height, _ := strconv.Atoi(string(heightReading))
			row = append(row, height)
		}
		heightMap = append(heightMap, row)
	}
	return heightMap
}
