package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Append(t *testing.T) {
	s := Stack[int]{}
	s.Append(4)
	s.Append(5)
	assert.ElementsMatch(t, []int{4, 5}, s.elems)
}

func TestStack_Pop(t *testing.T) {
	s := Stack[int]{}
	s.Append(4)
	s.Append(5)
	assert.Equal(t, 4, s.Pop())
	assert.ElementsMatch(t, []int{5}, s.elems)
}

func TestStack_PopN(t *testing.T) {
	s := Stack[int]{}
	s.Append(4)
	s.Append(5)
	s.Append(6)
	assert.ElementsMatch(t, []int{4, 5}, s.PopN(2))
	assert.ElementsMatch(t, []int{6}, s.elems)
}

func TestStack_Peek(t *testing.T) {
	s := Stack[int]{}
	s.Append(4)
	s.Append(5)
	assert.Equal(t, 4, s.Peek())
	assert.ElementsMatch(t, []int{4, 5}, s.elems)
}

func TestStack_Push(t *testing.T) {
	s := Stack[int]{}
	s.Push(4)
	s.Push(5)
	assert.ElementsMatch(t, []int{5, 4}, s.elems)
}
