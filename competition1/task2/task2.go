package main

import (
	"fmt"
	"sort"
)

type Jump struct {
	left, right int
	val         int
}

func NewJump(l, r, val int) Jump {
	return Jump{left: l, right: r, val: val}
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	jumps := make([]Jump, n)

	var l, r, x int
	for i := 0; i < n; i++ {
		fmt.Scan(&l, &r, &x)
		jumps[i] = NewJump(l, r, x)
	}

	sort.Slice(jumps, func(i, j int) bool {
		if jumps[i].left == jumps[j].left {
			return jumps[i].right < jumps[j].right
		}
		return jumps[i].left < jumps[j].left
	})

	for i := 0; i < m; i++ {
		var q int
		fmt.Scan(&q)

		ans := 0

		for j := 0; j < n; j++ {
			if jumps[j].left <= q && q <= jumps[j].right {
				if (q-jumps[j].left)%2 == 0 {
					ans += jumps[j].val
				} else {
					ans -= jumps[j].val
				}
			}
		}

		fmt.Println(ans)
	}
}
