package main

import (
	"fmt"
)

func max_heapify(A []int, i int, n int) {
	l := 2*i + 1
	r := 2*i + 2
	var lar int
	if l <= n && A[l] > A[i] {
		lar = l
	} else {
		lar = i
	}
	if r <= n && A[r] > A[lar] {
		lar = r
	}
	if lar != i {
		temp := A[i]
		A[i] = A[lar]
		A[lar] = temp
		max_heapify(A, lar, n)
	}
}
func build_heapify(A []int) {
	n := len(A) - 1
	for i := n / 2; i >= 0; i-- {
		max_heapify(A, i, n)
	}
}
func HeapSort(A []int) {
	//build_heapify(A)
	n := len(A) - 1
	var i int

	for i = n; i >= 0; i-- {
		temp := A[0]
		A[0] = A[i]
		A[i] = temp
		max_heapify(A, 0, i-1)

	}
}
func extract_max(a []int) int {
	if len(a) < 1 {
		fmt.Println("error")
	}
	build_heapify(a)
	fmt.Println(a)
	mx := a[0]
	a[0] = a[len(a)-1]
	a = a[:len(a)-1]
	max_heapify(a, 0, len(a)-1)
	return mx
}

func hepify_increase_key(a []int, key int, i int) {
	if a[i] > key {
		fmt.Println("error")
		return
	}
	a[i] = key
	for i >= 0 && a[i] > a[i/2] {
		temp := a[i]
		a[i] = a[i/2]
		a[i/2] = temp
		i = i / 2
	}
}
func main() {
	a := []int{1, 5, 3, 9, 8, 7, 6, 5}
	build_heapify(a)
	fmt.Println(a)
	HeapSort(a)
	fmt.Println(a)
	fmt.Println("extract max ", extract_max(a))
	a = a[:len(a)-1]
	fmt.Println(a)
	hepify_increase_key(a, 100, 4)
	fmt.Println("hepify increase key: ", a)
	//fmt.Println(a)
}
