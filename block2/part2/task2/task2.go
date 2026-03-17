package main

import (
	"bufio"
	"fmt"
	"os"
)

type Turtle struct {
	forward, backward int
}

func NewTurtle(a, b int) Turtle {
	return Turtle{
		forward:  a,
		backward: b,
	}
}

func main() {

	in := bufio.NewReader(os.Stdin)

	var n int

	fmt.Fscan(in, &n)

	setA := make(map[Turtle]bool)

	var a, b int

	for i := 0; i < n; i++ {

		fmt.Fscan(in, &a, &b)

		express := NewTurtle(a, b)

		if express.forward >= 0 && express.backward >= 0 {
			if !setA[express] {
				if express.forward+express.backward+1 == n {
					setA[express] = true
				}
			}
		}

	}

	fmt.Println(len(setA))

}
