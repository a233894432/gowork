package main

import "fmt"

func buildHeap(array []int, length int) {
	var i, j int
	for i = 1; i < length; i = i + 1 {
		for j = i; j > 0 && array[j] > array[(j-1)/2]; j = (j - 1) / 2 {
			array[j], array[(j-1)/2] = array[(j-1)/2], array[j]
		}
	}
}

func heapSort(array []int, length int) {
	array[0], array[length-1] = array[length-1], array[0]
	if length <= 2 {
		return
	}
	i, j := 0, 0
	for {
		j = 2*i + 1
		if j+1 < length-1 {
			if array[j] < array[j+1] {
				j = j + 1
			}
		} else if j >= length-1 {
			break
		}
		array[i], array[j] = array[j], array[i]
		i = j
	}
	heapSort(array, length-1)
}
func main() {
	primes := [6]int{3, 11, 5, 2, 13, 7}
	fmt.Println("orginal", primes)
	buildHeap(primes[:], len(primes))
	fmt.Println("Max heap", primes)
	heapSort(primes[:], len(primes))
	fmt.Println("after sorting", primes)
}
