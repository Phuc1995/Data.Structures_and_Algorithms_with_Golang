package main

import "fmt"

func main() {
	var string1 = "abc"
	r := []rune(string1)

	fmt.Println(string1)
	fmt.Println(string(r))
	for i,v := range string1 {
		fmt.Printf("%d-%v\n",i,v)
	}
fmt.Println()
	for i,v := range []rune(r) {
		fmt.Printf("%d-%v\n",i,v)
	}
}
