package main

import (
	"fmt"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)

	temp := 0
	ans := 0

	for i := 0; i < n; i++ {
		if s[i] != 'a' && s[i] != 'h' {
			temp = 0
		} else {
			if i > 0 && ((s[i-1] == 'a' && s[i] == 'h') || (s[i-1] == 'h' && s[i] == 'a')) {
				temp++
			} else {
				temp = 1
			}
		}

		if temp > ans {
			ans = temp
		}
	}

	fmt.Println(ans)
}
