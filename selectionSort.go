package main

import "fmt"

func main() {
	fmt.Println("enter the value of n")
	var n int
	fmt.Scanln(&n)
	fmt.Println("enter the value which want to sort")
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&arr[i])
	}
	sort(n, &arr)
	fmt.Println("sorted elements are:")
	fmt.Println(arr)
}
func sort(n int, arr *[]int) {
	var i, j, max int
	for i = 0; i < len(*arr)-1; i++ {
		max = 0
		for j = 1; j < len(*arr)-i; j++ {
			if (*arr)[j] > (*arr)[max] {
				max = j
			}
		}
		(*arr)[len(*arr)-1-i], (*arr)[max] = (*arr)[max], (*arr)[len(*arr)-1-i]
	}
}
