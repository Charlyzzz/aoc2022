package day01

import (
	"aoc2022/challenge"
	"sort"
	"strconv"
	"strings"
)

func Run() {
	part1()
	part2()
}

func part1() {
	res := caloriesForTop(1)
	challenge.OutputPart1(res)
}

func part2() {
	res := caloriesForTop(3)
	challenge.OutputPart2(res)
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
