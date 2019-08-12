package main

import (
	"fmt"
)

func main() {

	var n, c int = 15, 3

	var a = []int{1, 2, 4, 8, 9}

	var i, j, ans int64 = 0, 1000000000, 0

	for i <= j {

		mid := int((j + i) / 2)
		fi, temp := a[0], 1

		for k := 1; k < n; k++ {
			if a[k]-fi >= mid {
				temp++
				fi = a[k]
			}
		}
		if temp < c {
			j = int64(mid - 1)
		} else {
			ans = int64(mid)
			i = int64(mid + 1)
		}
	}

	fmt.Println(ans)
}
