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

	prefix := make([]int, n+1)
	prefix[0] = 0

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
		prefix[i+1] = prefix[i] + arr[i]
	}
	for i := 1; i < n+1; i++ {
		fmt.Printf("%d ", prefix[i])
	}
}
