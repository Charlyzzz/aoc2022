package day02

type rock struct{}

func (r rock) play(opponent pick) int {
	return opponent.againstRock() + r.value()
}

func (rock) againstRock() int {
	return draw
}

func (rock) againstPaper() int {
	return won
}

func (rock) againstScissors() int {
	return lost
}

func (rock) value() int {
	return 1
}

func (rock) win() pick {
	return paper{}
}

func (rock) draw() pick {
	return rock{}
}

func (rock) loose() pick {
	return scissors{}
}
