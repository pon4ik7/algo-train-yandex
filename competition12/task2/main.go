package main

import (
	"bufio"
	"fmt"
	"os"
)

const n = 10

type Point struct {
	x, y int
}

var field [n][]byte
var used [n][n]bool
var ship []Point

var dx = []int{-1, 1, 0, 0}
var dy = []int{0, 0, -1, 1}

func dfs(x, y int) {
	used[x][y] = true
	ship = append(ship, Point{x, y})

	for k := 0; k < 4; k++ {
		nx := x + dx[k]
		ny := y + dy[k]

		if nx >= 0 && nx < n && ny >= 0 && ny < n {
			if field[nx][ny] == '#' && !used[nx][ny] {
				dfs(nx, ny)
			}
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(in, &s)
		field[i] = []byte(s)
	}

	cnt := make(map[int]int)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if field[i][j] != '#' || used[i][j] {
				continue
			}

			ship = nil
			dfs(i, j)

			minX, maxX := ship[0].x, ship[0].x
			minY, maxY := ship[0].y, ship[0].y

			for _, p := range ship {
				if p.x < minX {
					minX = p.x
				}
				if p.x > maxX {
					maxX = p.x
				}
				if p.y < minY {
					minY = p.y
				}
				if p.y > maxY {
					maxY = p.y
				}
			}

			if minX != maxX && minY != maxY {
				fmt.Println("NO")
				return
			}

			need := (maxX - minX + 1) * (maxY - minY + 1)
			if need != len(ship) {
				fmt.Println("NO")
				return
			}

			for x := minX - 1; x <= maxX+1; x++ {
				for y := minY - 1; y <= maxY+1; y++ {
					if x < 0 || x >= n || y < 0 || y >= n {
						continue
					}
					if field[x][y] == '#' {
						inside := x >= minX && x <= maxX && y >= minY && y <= maxY
						if !inside {
							fmt.Println("NO")
							return
						}
					}
				}
			}

			size := len(ship)
			if size < 1 || size > 4 {
				fmt.Println("NO")
				return
			}

			cnt[size]++
		}
	}

	if cnt[1] == 4 && cnt[2] == 3 && cnt[3] == 2 && cnt[4] == 1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
