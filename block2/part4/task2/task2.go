package main

import (
	"bufio"
	"fmt"
	"os"
)

func getIndex(c byte) int {
	if c >= 'a' && c <= 'z' {
		return int(c - 'a')
	}
	return int(c-'A') + 26
}

func equal(a, b [52]int) bool {
	for i := 0; i < 52; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var g, s int
	var w, text string

	fmt.Fscan(in, &g, &s)
	fmt.Fscan(in, &w)
	fmt.Fscan(in, &text)

	var need [52]int
	var window [52]int

	for i := 0; i < g; i++ {
		need[getIndex(w[i])]++
		window[getIndex(text[i])]++
	}

	ans := 0
	if equal(need, window) {
		ans++
	}

	for i := g; i < s; i++ {
		window[getIndex(text[i-g])]--
		window[getIndex(text[i])]++

		if equal(need, window) {
			ans++
		}
	}

	fmt.Println(ans)
}
