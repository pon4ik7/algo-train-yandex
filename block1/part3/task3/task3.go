package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func countMin(A, B, C, D int) (int, int) {
	bestM, bestN := -1, -1

	update := func(m, n int) {
		if bestM == -1 || m+n < bestM+bestN {
			bestM, bestN = m, n
		}
	}

	if A > 0 && C > 0 {
		update(B+1, D+1)
	}

	if B > 0 && D > 0 {
		update(A+1, C+1)
	}

	if A > 0 && B > 0 {
		update(max(A, B)+1, 1)
	}

	if C > 0 && D > 0 {
		update(1, max(C, D)+1)
	}

	return bestM, bestN
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var A, B, C, D int
	fmt.Fscan(in, &A, &B, &C, &D)

	m, n := countMin(A, B, C, D)
	fmt.Println(m, n)
}
