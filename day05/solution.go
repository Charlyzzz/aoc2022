package day05

import (
	"aoc2022/challenge"
	"aoc2022/utils"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type c struct{}

func (c) Part1() interface{} {
	input := challenge.Input()
	return getTopStacks(input, true)
}

func (c) Part2() interface{} {
	input := challenge.Input()
	return getTopStacks(input, false)
}

func (c) Day() int {
	return 5
}

func Challenge() challenge.Challenge {
	return c{}
}

var containerRegexp = regexp.MustCompile(`[(\w+)]`)

func getTopStacks(input string, moveOneAtATime bool) string {
	var instructions []string
	var rows []string
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if isCrates(line) {
			rows = append(rows, line)
			continue
		}
		if isProcedure(line) {
			instructions = append(instructions, line)
		}
	}
	crates := stackCrates(lines, rows)
	for _, instruction := range instructions {
		n, from, to := parseInstruction(instruction)
		from -= 1
		to -= 1
		if moveOneAtATime {
			for i := 0; i < n; i++ {
				toMove := crates[from].Pop()
				crates[to].Push(toMove)
			}
		} else {
			var toMove []string
			for i := 0; i < n; i++ {
				toMove = append([]string{crates[from].Pop()}, toMove...)
			}
			for _, move := range toMove {
				crates[to].Push(move)
			}
		}
	}

	topCrates := strings.Builder{}
	for _, crate := range crates {
		topCrates.WriteString(crate.Peek())
	}
	return topCrates.String()
}

var instructionRegexp = regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

func parseInstruction(instruction string) (int, int, int) {
	match := instructionRegexp.FindStringSubmatch(instruction)
	var columns []int
	for _, position := range match[1:] {
		pos, _ := strconv.Atoi(position)
		columns = append(columns, pos)
	}
	return columns[0], columns[1], columns[2]
}

func stackCrates(lines []string, rows []string) []utils.Stack[string] {
	nCols := int(math.Ceil(float64(len(lines[0])) / 4))
	stacks := make([]utils.Stack[string], nCols)
	for _, crate := range rows {
		columns := utils.Every(4, []rune(crate))
		for colNumber, col := range columns {
			container := string(col)
			match := containerRegexp.FindStringSubmatch(container)
			if match != nil {
				crate := match[0]
				stacks[colNumber].Append(crate)
			}
		}
	}
	return stacks
}

func isProcedure(line string) bool {
	return strings.HasPrefix(line, "move")
}

func isCrates(line string) bool {
	return strings.Contains(line, "[")
}
