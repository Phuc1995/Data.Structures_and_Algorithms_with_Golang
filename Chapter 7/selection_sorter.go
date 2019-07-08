package main

import "fmt"

func SelectionSorter(element []int)   {
	var i int
	for i=0; i < len(element)-1; i++{
		var min int
		min = i
		for j := i + 1 ; j <= len(element) -1; j++{
			if element[j] < element[i]{
				min = j
			}
		}
		swap(element, i, min)
	}
}

func swap(element []int, i,j int)  {
	var temp int
	temp = element[j]

	element[j] = element[i]
	element[i] = temp
}

func main()  {
	var elements []int
	elements = []int{2,3,4,5,1}
	fmt.Println("Before Sorting ", elements)
	SelectionSorter(elements)
	fmt.Println("After Sorting", elements)
}
