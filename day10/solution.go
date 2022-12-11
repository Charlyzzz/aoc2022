package day10

import (
	"aoc2022/challenge"
	"aoc2022/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type c struct{}

func (c) Part1() interface{} {
	input := challenge.Input()
	return sumSignalStrength(input)
}

func (c) Part2() interface{} {
	input := challenge.Input()
	out := crtOutput(input)
	return fmt.Sprintf("\n%s\n", out)
}

func (c) Day() int {
	return 10
}

func Challenge() challenge.Challenge {
	return c{}
}

var sampleCycles = []int{20, 60, 100, 140, 180, 220}

func sumSignalStrength(input string) int {
	states := run(input)
	signalSum := 0
	for _, cycle := range sampleCycles {
		x := states[cycle-1]
		signalSum += strength(x, cycle)
	}
	return signalSum
}

func crtOutput(input string) string {
	states := run(input)
	output := make([]string, 240)
	for pixel := 0; pixel < 240; pixel++ {
		x := states[pixel]
		if isLit(pixel, x) {
			output[pixel] = "#"
		} else {
			output[pixel] = "."
		}
	}
	var crt strings.Builder
	lines := utils.Every(40, output)
	for l, line := range lines {
		for _, pixel := range line {
			crt.WriteString(pixel)
		}
		if l != len(lines)-1 {
			crt.WriteString("\n")
		}
	}
	return crt.String()
}

func isLit(pixel int, x int) bool {
	p := pixel % 40
	return p >= x-1 && p <= x+1
}

func run(program string) []int {
	states := []int{1}
	for _, instruction := range strings.Split(program, "\n") {
		if instruction == "noop" {
			states = append(states, fetchPrevState(states))
		} else {
			states = append(states, fetchPrevState(states))
			newState := addValue(instruction)
			states = append(states, fetchPrevState(states)+newState)
		}
	}
	return states
}

func fetchPrevState(states []int) int {
	if len(states) == 0 {
		return 1
	}
	return states[len(states)-1]
}

func addValue(instruction string) int {
	addValue, err := strconv.Atoi(instruction[5:])
	if err != nil {
		log.Fatalf("parsing instruction [%s] failed: %v", instruction, err)
	}
	return addValue
}

func strength(x int, cycle int) int {
	return x * cycle
}
