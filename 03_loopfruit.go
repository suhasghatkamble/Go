package main
import ("fmt")

func main() {
	fruits := [3]string{"apple", "orange","banana"}

	for j:=0; j< len(fruits); j++ {
		fmt.Println(fruits[j])
	}

	for _, val := range fruits {
		fmt.Printf("%v\n", val)
	}
}



// package main

// import "fmt"

// func main() {
// 	fruits := [3]string{"apple", "orange", "banana"}

// 	for j := 0; j < len(fruits); j++ {
// 		fmt.Println(fruits[j])
// 	}

// 	for i, val := range fruits {
// 		fmt.Printf("%v: %v\n", i, val)
// 	}
// }