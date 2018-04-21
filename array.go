package main

import "fmt"

func main() {
	a := [3]int{12, 78, 50}
	fmt.Println(a)

	b := [...]int{1, 2, 3, 4}
	fmt.Println(len(b))

	c := [...]int{235, 63, 2452, 2352, 262}

	for i, v := range c {
		fmt.Printf("#%d: %d\n", i+1, v)
	}

	d := [5]int{1, 2, 3, 4, 5}
	var e []int = d[1:4]
	fmt.Println(e)
}
