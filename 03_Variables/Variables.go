//Youtube Ref : https://www.youtube.com/watch?v=ijX6BbGdiKI&list=PLXQpH_kZIxTWUe-Ee-DZEX5gfeoo4tHV6&index=6

package main

import "fmt"

func main() {

	var name string = "Jaideep Singh Rathore" 
	
	// This is the first way of defining a variable in which you give name , type and define value all in one time and one line

	var name2 ="Aditi jeena"

	// This is another way of defining string in which you no need to define the type , GoLang infers the type and assigns it as required

	name3 :="Chunu munnu chinni"

	//This is the most direct way of defining a variable in which you can directly define a variable with lambda notation

	var name4 string

	name4="test"

	//Note that if you want to define a variable first and define it later then , string defination has to be given while declaring

	fmt.Println("The input string is :"+name)
	fmt.Println("The input string2 is :"+name2)
	fmt.Println("The input string3 is :"+name3)
	fmt.Println("The input string4 is :"+name4)

	//Variable needs to be used one defined otherwise there will be an error

	var age int = 40
	var isAdult bool = true
	var persecIncome float32 = 1.334

	fmt.Println("The input int is :",age)
	fmt.Println("The input bool is :",isAdult)
	fmt.Println("The input float is :",persecIncome)

}