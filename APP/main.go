package main

import "APP/Database"

func main(){
	_,err :=  Database.DatabaseConnection()

	if err != nil{
		fmt.Println(err)
	}
}