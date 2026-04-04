package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for ; t > 0; t-- {
		var n int
		var d int64
		fmt.Fscan(in, &n, &d)

		time := make([]int64, n)
		k := make([]int64, n)

		for i := 0; i < n; i++ {
			fmt.Fscan(in, &time[i], &k[i])
		}

		wait := make([]int64, n)
		var sum int64 = 0
		for i := 0; i < n; i++ {
			wait[i] = sum
			sum += k[i]
		}

		ans := n + 1

		ok := true
		for i := n - 1; i >= 0; i-- {
			free := time[i] - wait[i]
			if ok && free >= d {
				ans = i + 1
			} else {
				ok = false
			}
		}

		fmt.Fprintln(out, ans)
	}
}
