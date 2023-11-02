package main

import "fmt"

func main() {

	//a := []int{1, 2, 3, 4, 5}
	//fmt.Println(a)

	var c []string = make([]string, 3)
	fmt.Println(c, len(c))
	c = append(c, "zz", "xx", "cc")
	fmt.Println(c, len(c))

	var i []int = make([]int, 6, 10)
	fmt.Println(i, len(i))
	i = []int{1, 2}
	fmt.Println(i, len(i))
	i = append(i, 4, 6, 7, 8)
	fmt.Println(i, len(i))

	/*
		b := []int{7, 7, 7, 7}

		a = append(a, b...)
		fmt.Println(a)

		d := c[1:]
		fmt.Println(d)

		q := make([][]int, 0)
		q = append(q, []int{1, 2}, []int{5, 6})
		fmt.Println(q)

		w := make([]int, 0)
		w = append(w, 1, 2, 4, 5)
		fmt.Println(w)
	*/
}
