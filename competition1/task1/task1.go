package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	grid := make([][]rune, n)

	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		grid[i] = []rune(s)
	}

	ans := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '.' {
				if i+1 < n && grid[i+1][j] == '.' {
					ans++
				}
				if j+1 < m && grid[i][j+1] == '.' {
					ans++
				}
			}
		}
	}

	fmt.Println(ans)
}
