package main

import "Calculator/Arithmatic"
import "Calculator/Geometric"

import "fmt"

func main(){

	// Arithmatic
	fmt.Println("Addition : ",Arithmatic.Addition(2,5))
	fmt.Println("Division: ", Arithmatic.Division(10,5))

	// Geometric
	fmt.Println("Area : ", Geometric.AreaSquare(10))


}