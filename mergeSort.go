package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&arr[i])
	}
	sort(&arr, 0, len(arr)-1)
}

func sort(arr *[]int, low int, high int) {
	if low < high {
		mid := (low + high) / 2
		sort(arr, low, mid)
		sort(arr, mid+1, high)
		merge(arr, low, mid, high)
	}
}

func merge(arr *[]int, low int, mid int, high int) {
	pt := mid + 1
	count := low
	for low <= mid && pt <= high {
		if arr[low] > arr[pt] {
			temp := arr[count]
			arr[count] = arr[pt]
			arr[pt] = temp
			pt++
		}

	}
}
