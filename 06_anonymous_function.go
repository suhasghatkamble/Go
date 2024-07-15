package main
import "fmt"

func findSquare(num int)int{
	square:=num * num
 	return square
}

func main(){
	func(){
		fmt.Println("Welcome to CDAC ACTS !")
	}()

	//Assigning an anonnymous
	//function to a variable
	value :=func(){
		fmt.Println("Welcome to CDAC ACTS !")
	} 
	value()

	//anonymous function with arguments
	var sum = func(n1,n2 int){
		sum:=n1+n2
		fmt.Printf("Sum is : %d\n",sum)
	}
	sum(5,3)

	//return value from an anonymous function 
	var sum1 = func(n1,n2 int)int{
		sum1:=n1+n2
		return sum1
	}
	
	result:= sum1(5,3)
	fmt.Printf("Sum is : %d\n",result)

	area:=func(length, breadth int)int{
		return length * breadth
	}

	fmt.Printf("The are of reactangle is %d\n",area(3,4))

	// Anonymous fucntion as an arguement to another function 
	sum2:=func(number1 int, number2 int)int {
		return number1 + number2
	}

	result =findSquare(sum2(6,9))
	fmt.Printf("The result is :%d\n",result)
}