package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod int64 = 1000000007

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	var s1, s2, s3 int64

	for i := 0; i < n; i++ {
		var x int64
		fmt.Fscan(in, &x)
		x %= mod

		s3 = (s3 + s2*x) % mod
		s2 = (s2 + s1*x) % mod
		s1 = (s1 + x) % mod
	}

	fmt.Println(s3)
}
