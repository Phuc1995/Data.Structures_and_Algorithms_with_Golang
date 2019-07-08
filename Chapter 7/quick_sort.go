package main

import "fmt"

//Quick Sorter method
func QuickSorter(elements []int, below int, upper int) {
	if below < upper {
		var part int
		part = divideParts(elements, below, upper)
		QuickSorter(elements, below, part-1)
		QuickSorter(elements, part+1, upper)
	}
}

// divideParts method
func divideParts(elements []int, below int, upper int) int {
	var center int
	center = elements[upper]
	var i int
	i = below
	var j int
	for j = below; j < upper; j++ {
		if elements[j] <= center {
			swap1(&elements[i], &elements[j])
			i += 1
		}
	}
	swap1(&elements[i], &elements[upper])
	return i
}

//swap method
func swap1(element1 *int, element2 *int) {
	var val int
	val = *element1
	*element1 = *element2
	*element2 = val
}

// main method
func main() {
	var elements []int
	elements = []int{34, 202, 13, 19, 6, 5, 1, 43, 506, 12, 20, 28, 17, 100, 25, 4, 5, 97, 1000, 27}

	fmt.Print("Elements: ", elements, "\n")
	QuickSorter(elements, 0, len(elements)-1)
	fmt.Print("Sorted Elements: ", elements, "\n")
}
