package main

import (
	"bufio"
	"fmt"
	"os"
)

type Set struct {
	table map[int]bool
}

func NewSet() *Set {
	return &Set{make(map[int]bool)}
}

func (s *Set) add(x int) {
	s.table[x] = true
}

func (s *Set) pic(x int) bool {
	return s.table[x]
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var arr []int
	for {
		var x int
		_, err := fmt.Fscan(in, &x)
		if err != nil {
			break
		}
		arr = append(arr, x)
	}

	set := NewSet()

	for _, x := range arr {
		if set.pic(x) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		set.add(x)
	}

}
