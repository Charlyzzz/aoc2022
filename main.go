package main

import (
	"aoc2022/challenge"
	"aoc2022/day01"
	"aoc2022/day02"
	"aoc2022/day03"
	"aoc2022/day04"
	"aoc2022/day05"
	"aoc2022/day06"
	"aoc2022/day07"
	"aoc2022/day08"
	"aoc2022/day09"
	"aoc2022/day10"
)

func main() {
	challenges := []challenge.Challenge{
		day01.Challenge(),
		day02.Challenge(),
		day03.Challenge(),
		day04.Challenge(),
		day05.Challenge(),
		day06.Challenge(),
		day07.Challenge(),
		day08.Challenge(),
		day09.Challenge(),
		day10.Challenge(),
	}

	for _, c := range challenges {
		output := c.Part1()
		day := c.Day()
		challenge.OutputPart(day, 1, output)
		output = c.Part2()
		challenge.OutputPart(day, 2, output)
	}
}
