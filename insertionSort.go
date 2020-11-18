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
	var i, j, temp int
	for i = 1; i < len(*arr); i++ {
		temp = (*arr)[i]
		for j = i; j > 0 && (*arr)[j-1] > temp; j-- {
			(*arr)[j] = (*arr)[j-1]
		}
		(*arr)[j] = temp
	}
}
