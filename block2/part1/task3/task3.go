package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Set struct {
	table map[int]bool
}

func NewSet() *Set {
	return &Set{make(map[int]bool)}
}

func (s *Set) add(x ...int) {
	for _, i := range x {
		s.table[i] = true
	}
}

func (s *Set) pic(x int) bool {
	return s.table[x]
}

func mesh(first, second Set) []int {

	ans := make([]int, 0)

	for x := range first.table {
		if second.pic(x) {
			ans = append(ans, x)
		}
	}

	return ans

}

func parseLineToInts(line string) []int {
	fields := strings.Fields(line)
	arr := make([]int, len(fields))

	for i, s := range fields {
		arr[i], _ = strconv.Atoi(s)
	}
	return arr
}

func main() {
	in := bufio.NewReader(os.Stdin)

	line1, _ := in.ReadString('\n')
	line2, _ := in.ReadString('\n')

	a := parseLineToInts(line1)
	b := parseLineToInts(line2)

	setA := NewSet()
	setA.add(a...)
	setB := NewSet()
	setB.add(b...)

	ans := mesh(*setA, *setB)

	sort.Ints(ans)

	for _, x := range ans {
		fmt.Print(x, " ")
	}

}
