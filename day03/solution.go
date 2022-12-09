package day03

import (
	"aoc2022/challenge"
	"aoc2022/utils"
	"log"
	"strings"
)

type c struct{}

func (c) Part1() interface{} {
	input := challenge.Input()
	return sumDuplicatePriorities(input)
}

func (c) Part2() interface{} {
	input := challenge.Input()
	return sumBadgesPriorities(input)
}

func (c) Day() int {
	return 3
}

func Challenge() challenge.Challenge {
	return c{}
}

func sumBadgesPriorities(input string) int {
	rucksacks := parse(input)
	sum := 0
	for _, group := range utils.Every(3, rucksacks) {
		badge := findBadge(group)
		sum += itemPriority(badge)
	}
	return sum
}

func findBadge(group []string) string {
	occurrences := make(map[string]int)
	for i, rucksack := range group {
		for _, char := range rucksack {
			key := string(char)
			if occurrences[key] == i {
				occurrences[key] += 1
			}
		}
	}
	for badge, count := range occurrences {
		if count == len(group) {
			return badge
		}
	}
	log.Fatal("badge not found")
	return ""
}

func sumDuplicatePriorities(input string) int {
	dups := findDuplicates(input)
	sum := 0
	for _, dup := range dups {
		sum += itemPriority(dup)
	}
	return sum
}

func findDuplicates(input string) []string {
	var duplicates []string
	rucksack := parse(input)
	for _, items := range rucksack {
		compartmentLen := len(items) / 2
		firstCompartment := items[:compartmentLen]
		secondCompartment := items[compartmentLen:]
		duplicates = append(duplicates, compartmentDuplicates(firstCompartment, secondCompartment))
	}
	return duplicates
}

func parse(input string) []string {
	var rucksacks []string
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		rucksacks = append(rucksacks, line)
	}
	return rucksacks
}

func compartmentDuplicates(c1 string, c2 string) string {
	var duplicates []string
	occurences := map[string]int{}
	for _, char := range c1 {
		key := string(char)
		occurences[key] = 1
	}
	for _, char := range c2 {
		key := string(char)
		_, exists := occurences[key]
		if exists {
			occurences[key] = 0
		}
	}
	for char, count := range occurences {
		if count == 0 {
			duplicates = append(duplicates, char)
		}
	}
	return strings.Join(duplicates, "")
}

func itemPriority(item string) int {
	char := item[0]
	if isUpper(char) {
		return int(char - 38)
	}
	return int(char - 96)

}

func isUpper(char uint8) bool {
	return 65 <= char && char <= 90
}
