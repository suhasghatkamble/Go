package main

import "fmt"

func main() {
	var student1 string
	var student2 = "Acts" // type is inferred
	x := 20 // type is inferred
	student1 = "CDAC" // corrected assignment

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	var a, b, c, d int = 1, 3, 5, 7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	var (
		a1 int
		b1 int = 1
		c1 string = "hello"
	)

	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)
}