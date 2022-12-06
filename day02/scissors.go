package day02

type scissors struct{}

func (s scissors) play(opponent pick) int {
	return opponent.againstScissors() + s.value()
}

func (scissors) againstRock() int {
	return won
}

func (scissors) againstPaper() int {
	return lost
}

func (scissors) againstScissors() int {
	return draw
}

func (scissors) value() int {
	return 3
}

func (scissors) win() pick {
	return rock{}
}

func (scissors) draw() pick {
	return scissors{}
}

func (scissors) loose() pick {
	return paper{}
}
