package main

import (
	"bufio"
	"fmt"
	"os"
)

type Record struct {
	a, b int
	name string
}

func solve(set map[string]int, recs []Record) (string, int) {
	prevA, prevB := 0, 0

	for _, r := range recs {
		points := (r.a - prevA) + (r.b - prevB)
		set[r.name] += points
		prevA = r.a
		prevB = r.b
	}

	maxVal := -1
	name := ""
	for k, v := range set {
		if v > maxVal {
			maxVal = v
			name = k
		}
	}
	return name, maxVal
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	names := make([]string, n)
	set := make(map[string]int)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &names[i])
		set[names[i]] = 0
	}

	var m int
	fmt.Fscan(in, &m)

	records := make([]Record, m)
	for i := 0; i < m; i++ {
		var score string
		var name string
		fmt.Fscan(in, &score, &name)

		var a, b int
		fmt.Sscanf(score, "%d:%d", &a, &b)

		records[i] = Record{a: a, b: b, name: name}
	}

	resName, resScore := solve(set, records)
	fmt.Println(resName, resScore)
}
