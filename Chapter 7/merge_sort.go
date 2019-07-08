package main

import "fmt"

func MergeSorter(array []int) []int {
	if len(array) < 2 {
		return array
	}
	var middle int
	middle = len(array)/2
	return JoinArrays(MergeSorter(array[:middle]), MergeSorter(array[middle:]))
}

func JoinArrays(leftArr []int, rightArr []int)[]int {

	var num int
	var i int
	var j int
	num, i, j = len(leftArr)+len(rightArr), 0, 0
	var array []int
	array = make([]int, num, num)

	var k int
	for k = 0; k < num; k++ {
		if i > len(leftArr)-1 && j <= len(rightArr)-1 {
			array[k] = rightArr[j]
			j++
		} else if j > len(rightArr)-1 && i <= len(leftArr)-1 {
			array[k] = leftArr[i]
			i++
		} else if leftArr[i] < rightArr[j] {
			array[k] = leftArr[i]
			i++
		} else {
			array[k] = rightArr[j]
			j++
		}
	}
	return array
}

func main()  {
	var elements []int
	elements = []int{34, 202, 13, 19, 6, 5, 1, 43, 506, 12, 20, 28, 17, 100, 25, 4, 5, 97, 1000, 27}
	fmt.Println("\n Before Sorting \n\n", elements)
	fmt.Println("\n-After Sorting\n\n", MergeSorter(elements), "\n")
}
