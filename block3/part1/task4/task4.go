package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var k int64
	fmt.Fscan(in, &n, &k)

	arr := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}

	var sum int64
	ans := 0
	l := 0

	for r := 0; r < n; r++ {
		sum += arr[r]

		for sum > k {
			sum -= arr[l]
			l++
		}

		if sum == k {
			ans++
		}
	}

	fmt.Println(ans)
}
