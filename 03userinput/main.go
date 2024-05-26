package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome user!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your Name: ")

	//Comma OK Syntax (or) Error OK Syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for writing!")
	fmt.Printf("your name is " + input)
	fmt.Printf("The type is: %T \n", input)

}
