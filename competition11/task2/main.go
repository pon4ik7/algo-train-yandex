package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pair struct {
	a int
	b int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	var right string
	fmt.Fscan(in, &right)

	var m int
	fmt.Fscan(in, &m)

	answers := make([]string, m)
	correctCount := make([]int, m)

	for i := 0; i < m; i++ {
		fmt.Fscan(in, &answers[i])

		cnt := 0
		for j := 0; j < n; j++ {
			if answers[i][j] == right[j] {
				cnt++
			}
		}
		correctCount[i] = cnt
	}

	result := make([]Pair, 0)

	for i := 0; i < m; i++ {
		for j := i + 1; j < m; j++ {
			bothCorrect := 0
			bothWrongSame := 0

			for k := 0; k < n; k++ {
				if answers[i][k] == right[k] && answers[j][k] == right[k] {
					bothCorrect++
				} else if answers[i][k] == answers[j][k] && answers[i][k] != right[k] {
					bothWrongSame++
				}
			}

			c1 := correctCount[i]
			c2 := correctCount[j]
			w1 := n - c1
			w2 := n - c2

			if bothCorrect*2 > c1 &&
				bothCorrect*2 > c2 &&
				bothWrongSame*2 > w1 &&
				bothWrongSame*2 > w2 {
				result = append(result, Pair{i + 1, j + 1})
			}
		}
	}

	fmt.Println(len(result))
	for _, p := range result {
		fmt.Println(p.a, p.b)
	}
}
