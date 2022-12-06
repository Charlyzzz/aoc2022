package day06

import (
	"aoc2022/challenge"
	"log"
)

func Run() {
	part1()
	part2()
}

func part1() {
	input := challenge.Input()
	overlapCount := signalMarker(input, 4)
	challenge.OutputPart1(overlapCount)
}

func part2() {
	input := challenge.Input()
	res := signalMarker(input, 14)
	challenge.OutputPart2(res)
}

func signalMarker(input string, preambleLen int) int {
	var markerBuffer []uint8
	for i := 0; i < len(input); i++ {
		char := input[i]
		markerBuffer = appendAndTrim(preambleLen, char, markerBuffer)
		if len(markerBuffer) == preambleLen && allDifferent(&markerBuffer) {
			return i + 1
		}
	}
	log.Fatal("marker not found")
	return 0
}

func allDifferent(markerBuffer *[]uint8) bool {
	ocurrences := make(map[uint8]int, 4)
	for _, u := range *markerBuffer {
		if ocurrences[u] != 0 {
			return false
		}
		ocurrences[u] += 1
	}
	return true
}

func appendAndTrim[T any](take int, elem T, elems []T) []T {
	elems = append(elems, elem)
	trimLen := trimLen(take, &elems)
	return elems[trimLen:]
}

func trimLen[T any](take int, markerBuffer *[]T) int {
	if len(*markerBuffer) <= take {
		return 0
	}
	return len(*markerBuffer) - take
}
