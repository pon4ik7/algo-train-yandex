package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	arr := make([]int, n)
	bad := make([]bool, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
		k := (i + 1 - arr[i] + n) % n
		bad[k] = true
	}

	for i := 0; i < n; i++ {
		if !bad[i] {
			fmt.Println(i)
			return
		}
	}

	fmt.Println(-1)
}
