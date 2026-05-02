package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pair struct {
	color int
	cnt   int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	groups := make([]Pair, 0)
	for i := 0; i < n; i++ {
		if len(groups) > 0 && groups[len(groups)-1].color == a[i] {
			groups[len(groups)-1].cnt++
		} else {
			groups = append(groups, Pair{color: a[i], cnt: 1})
		}
	}

	stack := make([]Pair, 0)
	ans := 0

	for _, cur := range groups {
		if len(stack) > 0 && stack[len(stack)-1].color == cur.color {
			stack[len(stack)-1].cnt += cur.cnt
		} else {
			stack = append(stack, cur)
		}

		for len(stack) > 0 && stack[len(stack)-1].cnt >= 3 {
			ans += stack[len(stack)-1].cnt
			stack = stack[:len(stack)-1]

			if len(stack) >= 2 && stack[len(stack)-1].color == stack[len(stack)-2].color {
				stack[len(stack)-2].cnt += stack[len(stack)-1].cnt
				stack = stack[:len(stack)-1]
			}
		}
	}

	fmt.Println(ans)
}
