package main

import "fmt"

func main() {
	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}

	for x := 0; x < len(a); x++ {
		for y := 0; y < len(a[x]); y++ {
			fmt.Printf("a[%d][%d] = %d\n", x, y, a[x][y])
		}
	}
}
