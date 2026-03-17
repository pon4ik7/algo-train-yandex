package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Pair struct {
	x, y int
}

func checkFigure(points ...Pair) string {

	sort.Slice(points, func(i, j int) bool {
		if points[i].x == points[j].x {
			return points[i].y < points[j].y
		}
		return points[i].x < points[j].x
	})

	if points[0].x-points[1].x == points[2].x-points[3].x && points[0].y-points[1].y == points[2].y-points[3].y {
		return "YES"
	}

	return "NO"
}

func main() {

	in := bufio.NewReader(os.Stdin)

	var n int

	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		var point1, point2, point3, point4 Pair

		fmt.Fscan(in, &point1.x, &point1.y, &point2.x, &point2.y, &point3.x, &point3.y, &point4.x, &point4.y)

		fmt.Println(checkFigure(point1, point2, point3, point4))

	}

}
