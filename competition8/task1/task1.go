package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscan(in, &s)

	grid := make(map[Point]int)

	x, y := 0, 0
	start := Point{x, y}
	grid[start] = 1

	for _, c := range s {

		switch c {
		case 'U':
			y++
		case 'D':
			y--
		case 'R':
			x++
		case 'L':
			x--
		}

		p := Point{x, y}
		grid[p]++
	}

	ans := 0
	for _, cnt := range grid {
		if cnt > 1 {
			ans++
		}
	}

	fmt.Println(ans)
}
