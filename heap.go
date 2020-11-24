package main

import "fmt"

func main() {
	a := []int{1, 5, 3, 9, 8, 7, 6, 5}
	heap(a)
	fmt.Println(a)
}

func heap(arr []int) {
	n := len(arr) - 1
	for i := n / 2; i >= 0; i-- {
		add(arr, i)
	}
}
func add(arr []int, val int) {
	left := 2*val + 1
	right := left + 1
	var sm int = -1
	if left <= len(arr) {
		sm = left
	}
	if right <= len(arr) && arr[right] > arr[left] {
		sm = right
	}
	if sm != -1 && arr[val] > arr[sm] {
		temp := arr[val]
		arr[val] = arr[sm]
		arr[sm] = temp
		add(arr, sm)
	}

}
