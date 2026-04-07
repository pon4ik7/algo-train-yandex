package main

import (
	"bufio"
	"fmt"
	"os"
)

func clean(grid [][]rune, i, j int) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[i]) {
		return
	}
	if grid[i][j] == '.' {
		return
	}
	grid[i][j] = '.'
	clean(grid, i, j+1)
	clean(grid, i, j-1)
	clean(grid, i+1, j)
	clean(grid, i-1, j)

}

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

	count := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '#' {
				count++
				clean(grid, i, j)
			}
		}
	}
	fmt.Println(count)

}
