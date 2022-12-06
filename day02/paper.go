package day02

type paper struct{}

func (p paper) play(opponent pick) int {
	return opponent.againstPaper() + p.value()
}

func (paper) againstRock() int {
	return lost
}

func (paper) againstPaper() int {
	return draw
}

func (paper) againstScissors() int {
	return won
}

func (p paper) value() int {
	return 2
}

func (paper) win() pick {
	return scissors{}
}

func (paper) draw() pick {
	return paper{}
}

func (paper) loose() pick {
	return rock{}
}
