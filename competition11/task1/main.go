package main

import (
	"bufio"
	"fmt"
	"os"
)

type Bracket struct {
	open  byte
	close byte
}

func NewBracket(open, close byte) Bracket {
	return Bracket{
		open:  open,
		close: close,
	}
}

func (b Bracket) Match(openCh, closeCh byte) bool {
	return b.open == openCh && b.close == closeCh
}

type Stack struct {
	data []byte
}

func NewStack() *Stack {
	return &Stack{
		data: make([]byte, 0),
	}
}

func (s *Stack) Push(x byte) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() byte {
	last := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return last
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func isCorrect(t string) bool {
	if len(t) == 0 {
		return true
	}

	stack := NewStack()

	round := NewBracket('(', ')')
	square := NewBracket('[', ']')
	curly := NewBracket('{', '}')

	for i := 0; i < len(t); i++ {
		ch := t[i]

		if ch == '(' || ch == '[' || ch == '{' {
			stack.Push(ch)
		} else {
			if stack.Empty() {
				return false
			}

			top := stack.Pop()

			ok := false
			if round.Match(top, ch) {
				ok = true
			}
			if square.Match(top, ch) {
				ok = true
			}
			if curly.Match(top, ch) {
				ok = true
			}

			if !ok {
				return false
			}
		}
	}

	return stack.Empty()
}

func shiftString(s string, k int) string {
	return s[k:] + s[:k]
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscan(in, &s)

	if len(s)%2 == 1 {
		fmt.Println("NO")
		return
	}

	if isCorrect(s) {
		fmt.Println("YES")
		return
	}

	for k := 1; k < len(s); k++ {
		t := shiftString(s, k)
		if isCorrect(t) {
			fmt.Println("YES")
			return
		}
	}

	fmt.Println("NO")
}
