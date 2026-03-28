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

	var n int
	fmt.Fscan(in, &n)

	var left int64 = 0
	var right int64 = 2000000000

	for i := 0; i < n; i++ {
		var x, d int64
		fmt.Fscan(in, &x, &d)

		l := x - d
		r := x + d

		if l > left {
			left = l
		}
		if r < right {
			right = r
		}
	}

	if left > right {
		fmt.Fprintln(out, -1)
	} else {
		fmt.Fprintln(out, right)
	}
}
