package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReaderSize(os.Stdin, 1<<20)

	var n int
	fmt.Fscan(in, &n)

	words := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &words[i])
	}

	m := len(words[0])

	cnt := make([]map[uint64]int, m+1)
	for i := 1; i <= m; i++ {
		cnt[i] = make(map[uint64]int)
	}

	const base uint64 = 911382323

	for _, s := range words {
		var h uint64 = 0
		for i := 0; i < m; i++ {
			h = h*base + uint64(s[i]-'a'+1)
			cnt[i+1][h]++
		}
	}

	ans := 0
	for k := 1; k <= m; k++ {
		ok := true
		for _, v := range cnt[k] {
			if v%2 != 0 {
				ok = false
				break
			}
		}
		if ok {
			ans = k
		} else {
			break
		}
	}

	fmt.Println(ans)
}
