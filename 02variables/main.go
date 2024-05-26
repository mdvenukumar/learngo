package main

import "fmt"

const LoginToken string = "Iamlovely" //public -> first letter capital

func main() {

	// String Type
	var username string = "venu" //if you create a variable you must use in go
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)
	fmt.Println("---------------------------------------------------")

	// Boolean Type
	var isLoggedIn bool = true //if you create a variable you must use in go
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)
	fmt.Println("---------------------------------------------------")

	// Integer Type
	var number int = 8 //if you create a variable you must use in go
	fmt.Println(number)
	fmt.Printf("variable is of type: %T \n", number)
	fmt.Println("---------------------------------------------------")

	// float32 / float64 Type
	var decimal float64 = 2.2351345345 //if you create a variable you must use in go
	fmt.Println(decimal)
	fmt.Printf("variable is of type: %T \n", decimal)
	fmt.Println("---------------------------------------------------")

	// Default values and some aliases
	var other int
	fmt.Println(other)
	fmt.Printf("Variable is of type: %T \n", other) // the value is set 0
	fmt.Println("---------------------------------------------------")

	// implicit type
	var website = "venukumarmd.web.app"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)
	fmt.Println("---------------------------------------------------")

	// No Var style
	userNo := 10
	fmt.Println(userNo)
	fmt.Printf("Variable is of type: %T \n", userNo)
	fmt.Println("---------------------------------------------------")

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
	fmt.Println("---------------------------------------------------")
}
