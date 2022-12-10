package day09

import (
	"aoc2022/challenge"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type c struct{}

func (c) Part1() interface{} {
	input := challenge.Input()
	return countTailPositions(input, 2)
}

func (c) Part2() interface{} {
	input := challenge.Input()
	return countTailPositions(input, 10)
}

func (c) Day() int {
	return 9
}

func Challenge() challenge.Challenge {
	return c{}
}

var visited = struct{}{}

type visitMap = map[string]interface{}

func markVisited(m *visitMap, pos *position) {
	(*m)[pos.key()] = visited
}

type position struct {
	x int
	y int
}

func (p *position) key() string {
	return fmt.Sprintf("%d-%d", p.x, p.y)
}

type movement = func(pos *position, delta int)

func up(pos *position, delta int) {
	pos.y += delta
}

func down(pos *position, delta int) {
	pos.y -= delta
}

func right(pos *position, delta int) {
	pos.x += delta
}

func left(pos *position, delta int) {
	pos.x -= delta
}

func countTailPositions(input string, knots int) int {
	v := newVisitMap()
	var rope []*position
	for i := 0; i < knots; i++ {
		knot := new(position)
		rope = append(rope, knot)
	}
	head := rope[0]
	tail := rope[knots-1]
	for _, command := range strings.Split(input, "\n") {
		dir, delta := parseCommand(command)
		for i := 1; i <= delta; i++ {
			dir(head, 1)
			for n := 0; n < knots-1; n++ {
				k1 := rope[n]
				k2 := rope[n+1]
				follow(k1, k2)
			}
			markVisited(v, tail)
		}
	}
	return len(*v)
}

func knotName(k int) string {
	if k == 0 {
		return "H"
	} else {
		return strconv.Itoa(k)
	}
}

func newVisitMap() *visitMap {
	v := make(visitMap, 0)
	return &v
}

func follow(head *position, tail *position) {
	if samePosition(head, tail) || lessThanTwoAway(head, tail) {
		return
	}
	deltaX := normalize(head.x - tail.x)
	deltaY := normalize(head.y - tail.y)
	tail.x += deltaX
	tail.y += deltaY
}

func normalize(x int) int {
	if x == 0 {
		return 0
	}
	if x > 0 {
		return 1
	} else {
		return -1
	}
}

func lessThanTwoAway(head *position, tail *position) bool {
	return intAbs(head.x-tail.x) < 2 && intAbs(head.y-tail.y) < 2
}

func samePosition(head *position, tail *position) bool {
	return head.x == tail.x && head.y == tail.y
}

func parseCommand(command string) (movement, int) {
	chunks := strings.Split(command, " ")
	delta, err := strconv.Atoi(chunks[1])
	if err != nil {
		log.Fatalf("parsing failed [%s]: %v", command, err)
	}
	return direction(chunks[0]), delta
}

func direction(s string) movement {
	switch s {
	case "U":
		return up
	case "D":
		return down
	case "L":
		return left
	case "R":
		return right
	default:
		log.Fatalf("unknown direction %s", s)
		return nil
	}
}

func intAbs(n int) int {
	return int(math.Abs(float64(n)))
}
