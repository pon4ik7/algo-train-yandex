package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	in := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscan(in, &s)

	a := make([]int, len(s))

	sum := 0

	for i := 0; i < len(s); i++ {
		a[i] = int(s[i] - '0')
		sum += a[i]
	}

	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })

	if sum%3 == 0 {
		for i := 0; i < len(a); i++ {
			fmt.Print(a[i])
		}
	} else {
		r := sum % 3

		for i := len(a) - 1; i >= 0; i-- {
			if a[i]%3 == r {
				a = append(a[:i], a[i+1:]...)
				for _, x := range a {
					fmt.Print(x)
				}
				return
			}
		}

		need := 3 - r
		cnt := 0
		res := []int{}

		for i := len(a) - 1; i >= 0; i-- {
			if a[i]%3 == need && cnt < 2 {
				cnt++
				continue
			}
			res = append(res, a[i])
		}

		for i := len(res) - 1; i >= 0; i-- {
			fmt.Print(res[i])
		}
	}
}
