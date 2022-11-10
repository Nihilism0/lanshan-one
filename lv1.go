package main

import "fmt"

func main() {
	var n int
	var r int
	var l int
	fmt.Scanf("%d %d %d\n", &n, &l, &r)
	var arr [1000]int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	for i := l; i < n-1; i++ {
		for j := l + 1; j < n; j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
}
