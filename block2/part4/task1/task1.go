package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	var s string
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &s)

	cnt := make([]int, 26)

	for _, ch := range s {
		cnt[ch-'A']++
	}

	var left strings.Builder
	middle := ""

	for i := 0; i < 26; i++ {
		for j := 0; j < cnt[i]/2; j++ {
			left.WriteByte(byte('A' + i))
		}
		if cnt[i]%2 == 1 && middle == "" {
			middle = string(byte('A' + i))
		}
	}

	leftPart := left.String()

	var right strings.Builder
	for i := len(leftPart) - 1; i >= 0; i-- {
		right.WriteByte(leftPart[i])
	}

	fmt.Println(leftPart + middle + right.String())
}
