package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Pair struct {
	val int64
	rem int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, k int
	fmt.Fscan(in, &n, &k)

	a := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	minPref := make([]int64, k)
	for i := 0; i < k; i++ {
		minPref[i] = math.MaxInt64
	}

	minPref[0] = 0

	best1 := Pair{0, 0}
	best2 := Pair{math.MaxInt64, -1}

	var pref int64 = 0
	var ans int64 = 0

	for i := 0; i < n; i++ {
		pref += a[i]

		rem := int(pref % int64(k))
		if rem < 0 {
			rem += k
		}

		var candidate int64 = -math.MaxInt64
		if best1.rem != rem {
			candidate = pref - best1.val
		} else if best2.rem != -1 {
			candidate = pref - best2.val
		}

		if candidate > ans {
			ans = candidate
		}

		if pref < minPref[rem] {
			minPref[rem] = pref

			if best1.rem == rem {
				best1.val = pref
			} else if best2.rem == rem {
				best2.val = pref
				if best2.val < best1.val {
					best1, best2 = best2, best1
				}
			} else {
				if pref < best1.val {
					best2 = best1
					best1 = Pair{pref, rem}
				} else if pref < best2.val {
					best2 = Pair{pref, rem}
				}
			}
		}
	}

	fmt.Println(ans)
}
