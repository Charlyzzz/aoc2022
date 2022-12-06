package challenge

import (
	"log"
	"os"
	"path"
	"regexp"
	"runtime"
)

var packageRegexp = regexp.MustCompile(`aoc2022/day(\d{2})`)

func Input() string {
	day := challengeDay()
	inputFile, err := os.ReadFile(path.Join("day"+day, "input"))
	if err != nil {
		log.Fatal(err)
	}
	return string(inputFile)
}

func challengeDay() string {
	_, file, _, ok := runtime.Caller(2)
	if !ok {
		log.Fatal("could not infer challenge input")
	}
	matches := packageRegexp.FindStringSubmatch(file)
	day := matches[1]
	return day
}

func OutputPart1(res interface{}) {
	log.Printf("day %s - part 1: %v", challengeDay(), res)
}

func OutputPart2(res interface{}) {
	log.Printf("day %s - part 2: %v", challengeDay(), res)
	log.Println("---------------------")
}
