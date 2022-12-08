package day02

import (
	"aoc2022/challenge"
	"log"
	"strings"
)

const (
	won  = 6
	draw = 3
	lost = 0
)

type pick interface {
	play
	againstRock() int
	againstPaper() int
	againstScissors() int
	win() pick
	draw() pick
	loose() pick
	value() int
}

type play interface {
	play(opponent pick) int
}

type playWin struct{}

func (p playWin) play(opponent pick) int {
	return opponent.win().play(opponent)
}

type playDraw struct{}

func (p playDraw) play(opponent pick) int {
	return opponent.draw().play(opponent)
}

type playLoose struct{}

func (p playLoose) play(opponent pick) int {
	return opponent.loose().play(opponent)
}

type c struct{}

func (c) Part1() interface{} {
	return computeScore(true)
}

func (c) Part2() interface{} {
	return computeScore(false)
}

func Challenge() challenge.Challenge {
	return c{}
}

func computeScore(isPlay bool) int {
	input := challenge.Input()
	return strategyScore(input, isPlay)
}

func strategyScore(input string, isPlay bool) int {
	score := 0
	strategy := parseTournament(input, isPlay)
	for _, round := range strategy {
		score += round.me.play(round.opponent)
	}
	return score
}

func parseTournament(input string, isPlay bool) []round {
	var tournament []round
	rounds := strings.Split(input, "\n")
	for _, line := range rounds {
		if line == "" {
			continue
		}
		round := parseRound(line, isPlay)
		tournament = append(tournament, round)
	}
	return tournament
}

type round struct {
	opponent pick
	me       play
}

func parseRound(line string, isPlay bool) round {
	var round round
	for _, move := range strings.Split(line, " ") {
		switch move {
		case "A":
			round.opponent = rock{}
		case "B":
			round.opponent = paper{}
		case "C":
			round.opponent = scissors{}
		case "X":
			if isPlay {
				round.me = rock{}
			} else {
				round.me = playLoose{}
			}
		case "Y":
			if isPlay {
				round.me = paper{}
			} else {
				round.me = playDraw{}
			}
		case "Z":
			if isPlay {
				round.me = scissors{}
			} else {
				round.me = playWin{}
			}
		default:
			log.Fatalf("unexpected move: [%s]", move)
		}
	}
	return round
}
