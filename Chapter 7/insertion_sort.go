package main

import (
	"fmt"
)

func InsertionSorter(elements []int)  {
	var n = len(elements)
	var i int

	for i =1; i < n; i++{
		var j =i

		for j > 0{
			if elements[j-1] > elements[j]{
				elements[j-1], elements[j] = elements[j], elements[j-1]
			}
			j = j -1
		}
	}
}

func main()  {
	var sequence []int
	sequence = []int{5,1,4,7}
	fmt.Println("\n^^^^^^ Before Sorting ^^^ \n\n", sequence)
	InsertionSorter(sequence)
	fmt.Println("\n--- After Sorting ---\n\n", sequence, "\n")
}
