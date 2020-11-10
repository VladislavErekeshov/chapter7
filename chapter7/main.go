package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 18, 17,
	}
	ans := x[0]
	for i := 0; i < len(x)-1; i++ {
		if ans < x[i+1] {
			ans = ans
		} else {
			ans = x[i+1]
		}
	}
	fmt.Println(ans)
}
