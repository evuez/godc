package vm

import (
	"errors"
	"fmt"
	"strconv"
)

type Stack struct {
	elems []string
	count int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(e interface{}) {
	s.elems = append(s.elems[:s.count], fmt.Sprintf("%v", e))
	s.count++
}

func (s *Stack) Pop() (string, error) {
	if s.count == 0 {
		return "", errors.New("Can't pop out of an empty stack.")
	}
	s.count--
	return s.elems[s.count], nil
}

func (s *Stack) PopInt() (int, error) {
	top, errPop := s.Pop()

	if errPop != nil {
		return 0, errPop
	}

	topInt, errConv := strconv.Atoi(top)

	if errConv != nil {
		return 0, errConv
	}

	return topInt, nil
}

func (s *Stack) Top() (string, error) {
	if s.count == 0 {
		return "", errors.New("Can't get the top of an empty stack.")
	}
	return s.elems[s.count-1], nil
}
