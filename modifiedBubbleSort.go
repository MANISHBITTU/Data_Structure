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
	swapped := 1
	for i := 0; i < n-1 && swapped == 1; i++ {
		swapped = 0
		for j := 0; j < n-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
				swapped = 1
			}
		}
	}
}
