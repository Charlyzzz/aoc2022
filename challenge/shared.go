package challenge

import (
	"log"
	"os"
	"path"
	"regexp"
	"runtime"
)

func Input() string {
	day := challengeDay()
	inputFile, err := os.ReadFile(path.Join("day"+day, "input"))
	if err != nil {
		log.Fatal(err)
	}
	return string(inputFile)
}

type Challenge interface {
	Part1() interface{}
	Part2() interface{}
}

var packageRegexp = regexp.MustCompile(`aoc2022/day(\d{2})`)

func challengeDay() string {
	_, file, _, ok := runtime.Caller(2)
	if !ok {
		log.Fatal("could not infer challenge input")
	}
	matches := packageRegexp.FindStringSubmatch(file)
	day := matches[1]
	return day
}

func OutputPart(day int, part int, res interface{}) {
	log.Printf("day %d - part %d: %v", day, part, res)
}
