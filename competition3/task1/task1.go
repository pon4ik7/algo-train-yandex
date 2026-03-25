package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func solve(line string) string {
	line = strings.TrimSpace(line)

	parts := strings.Split(line, " ")

	ans := ""

	for _, part := range parts {
		left := 0
		right := 0

		for left < len(part) && part[left] == '\'' {
			left++
		}

		for right < len(part) && part[len(part)-1-right] == '\'' {
			right++
		}

		word := part[left : len(part)-right]

		ans += word[left : len(word)-right]
	}
	return ans
}

func main() {
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')

	ans := solve(line)

	fmt.Println(ans)
}
