package main

import (
	"data-struct/sort/baseSort"
	"fmt"
)

func main() {
	arr := []int{3, 40, 5, 30, 3, 40, 3, 50, 3, 0, 4, 3, 2, 45, 56, 78, 2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 55}
	fmt.Println(baseSort.RadixSort(arr))
}
