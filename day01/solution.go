package day01

import (
	"aoc2022/challenge"
	"sort"
	"strconv"
	"strings"
)

type c struct{}

func (c) Part1() interface{} {
	return caloriesForTop(1)
}

func (c) Part2() interface{} {
	return caloriesForTop(3)
}

func (c) Day() int {
	return 1
}

func Challenge() challenge.Challenge {
	return c{}
}

func caloriesForTop(elvesCount int) int {
	input := challenge.Input()
	return sumTopNthCaloriesCarried(input, elvesCount)
}

func sumTopNthCaloriesCarried(input string, take int) int {
	inventory := parseElvesInventory(input)
	sort.Sort(sort.Reverse(sort.IntSlice(inventory)))
	topCarryingElves := inventory[:take]
	totalCalories := 0
	for _, calories := range topCarryingElves {
		totalCalories += calories
	}
	return totalCalories
}

func parseElvesInventory(input string) []int {
	var elves []int
	inventoryByElf := strings.Split(input, "\n\n")
	for _, foodCarried := range inventoryByElf {
		calories := caloriesCarried(foodCarried)
		elves = append(elves, calories)
	}
	return elves
}

func caloriesCarried(carried string) int {
	caloriesCarried := 0
	for _, food := range strings.Split(carried, "\n") {
		caloricCount, _ := strconv.Atoi(food)
		caloriesCarried += caloricCount
	}
	return caloriesCarried
}
