package day04

import (
	"aoc2022/challenge"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type c struct{}

func (c) Part1() interface{} {
	input := challenge.Input()
	return countAssignmentOverlaps(input, true)
}

func (c) Part2() interface{} {
	input := challenge.Input()
	return countAssignmentOverlaps(input, false)
}

func Challenge() challenge.Challenge {
	return c{}
}

func countAssignmentOverlaps(input string, total bool) int {
	count := 0
	for _, line := range strings.Split(input, "\n") {
		elfAssignments := parseAssignment(line)
		if elfAssignments.overlaps(total) {
			count += 1
		}
	}
	return count
}

type areas struct {
	from int
	to   int
}

func (a areas) contains(otherArea areas) bool {
	return a.from <= otherArea.from && a.to >= otherArea.to
}

func (a areas) intersects(otherArea areas) bool {
	return a.withinBoundaries(otherArea.from) || a.withinBoundaries(otherArea.to)
}

func (a areas) withinBoundaries(areaId int) bool {
	return areaId >= a.from && areaId <= a.to
}

type assignment struct {
	elf1 areas
	elf2 areas
}

func (a assignment) overlaps(total bool) bool {
	if total {
		return a.elf1.contains(a.elf2) || a.elf2.contains(a.elf1)
	} else {
		return a.elf1.intersects(a.elf2) || a.elf2.intersects(a.elf1)
	}
}

var assignmentRegexp = regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)

func parseAssignment(line string) assignment {
	matches := assignmentRegexp.FindStringSubmatch(line)
	var areaIds []int
	for _, match := range matches[1:5] {
		areaId, err := strconv.Atoi(match)
		if err != nil {
			log.Fatalf("parsing area failed %v", err)
		}
		areaIds = append(areaIds, areaId)
	}
	return assignment{
		elf1: areas{from: areaIds[0], to: areaIds[1]},
		elf2: areas{from: areaIds[2], to: areaIds[3]},
	}
}
