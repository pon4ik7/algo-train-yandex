package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscan(in, &s)

	var mean float64
	var lowest float64 = 100

	for _, char := range s {
		x := 25 - int(char-'A')
		if float64(x) < lowest {
			lowest = float64(x)
		}
		mean += float64(x)
	}

	mean = mean / float64(len(s))

	ans := math.Round(mean)

	if ans > lowest+1 {
		ans = lowest + 1
	}

	fmt.Println(string('A' + rune(25-int(ans))))
}
