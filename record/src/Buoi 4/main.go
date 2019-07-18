package main

import (
	"fmt"
)

func main()  {
	array := [2000]int{1,2,3,4,5}
	s1 := append([]int{}, array[:2]...)
	array[4] = 5
	fmt.Println(cap(s1))
}
