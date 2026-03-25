package main

import (
	"fmt"
	"math"
)

func check(m int64, ans *int64) {
	for d := int64(1); d*d <= m; d++ {
		if m%d == 0 {
			x := d
			y := m / d

			if y%2 == 1 {
				r := x
				k := (y + 1) / 2
				*ans = int64(math.Min(float64(*ans), math.Abs(float64(r-k))))
			}

			if x != y && x%2 == 1 {
				r := y
				k := (x + 1) / 2
				*ans = int64(math.Min(float64(*ans), math.Abs(float64(r-k))))
			}
		}
	}
}

func solve(n int64) int64 {
	ans := int64(1000000000000000000)

	for d := int64(1); d*d <= n; d++ {
		if n%d == 0 {
			ans = int64(math.Min(float64(ans), math.Abs(float64(d-n/d))))
		}
	}

	check(2*n-1, &ans)
	check(2*n, &ans)
	check(2*n+1, &ans)

	return ans
}

func main() {
	var n int64
	fmt.Scan(&n)
	fmt.Println(solve(n))
}
