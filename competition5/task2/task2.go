package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscan(in, &s)
	n := len(s)

	positions := make([][]int, 26)
	for i := 0; i < n; i++ {
		positions[s[i]-'a'] = append(positions[s[i]-'a'], i)
	}

	maxCount := 0
	for i := 0; i < 26; i++ {
		if len(positions[i]) > maxCount {
			maxCount = len(positions[i])
		}
	}

	if maxCount == 1 {
		fmt.Println(n)
		return
	}

	const hashMul uint64 = 911382323
	pow := make([]uint64, n+1)
	pref := make([]uint64, n+1)

	pow[0] = 1
	for i := 0; i < n; i++ {
		pow[i+1] = pow[i] * hashMul
		pref[i+1] = pref[i]*hashMul + uint64(s[i])
	}

	getHash := func(l, r int) uint64 {
		return pref[r] - pref[l]*pow[r-l]
	}

	can := func(pos []int, length int) bool {
		if pos[len(pos)-1]+length > n {
			return false
		}

		for i := 0; i+1 < len(pos); i++ {
			if pos[i]+length > pos[i+1] {
				return false
			}
		}

		firstHash := getHash(pos[0], pos[0]+length)
		for i := 1; i < len(pos); i++ {
			if getHash(pos[i], pos[i]+length) != firstHash {
				return false
			}
		}

		return true
	}

	answer := 1

	for c := 0; c < 26; c++ {
		if len(positions[c]) != maxCount {
			continue
		}

		pos := positions[c]

		left, right := 1, n
		best := 1

		for left <= right {
			mid := (left + right) / 2

			if can(pos, mid) {
				best = mid
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

		if best > answer {
			answer = best
		}
	}

	fmt.Println(answer)
}
