//might be error if you build it
//because main and clearScreen func is being declare in clean.go and main.go
//use go run main.go or go run clean.go to run it individually
package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var option int
var num1,num2 float32

func main(){
	for{
		menu()
		fmt.Scanln(&option)
		if(option == 5){
			fmt.Printf("Exit Program")
			fmt.Scanln()
			os.Exit(0)
		}
		fmt.Printf("Number 1 = \n")
		fmt.Scanln(&num1)
		fmt.Printf("Number 2 = \n")
		fmt.Scanln(&num2)
		executeOption(option, num1, num2)
		clearScreen()
	}
}

func menu(){
	fmt.Println("Simple Calculator")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Println("5. Exit")
	fmt.Print("Choose Option: ")
}
func executeOption(option int, num1, num2 float32) {
	switch option {
	case 1:
		fmt.Printf("Result %.2f + %.2f = %.2f\n", num1, num2, Addition(num1,num2))
	case 2:
		fmt.Printf("Result %.2f - %.2f = %.2f\n", num1, num2, Subtraction(num1,num2))
	case 3:
		fmt.Printf("Result %.2f * %.2f = %.2f\n", num1, num2, Multiplication(num1,num2))
	case 4:
		if num2 != 0 {
			fmt.Printf("Result %.2f / %.2f = %.2f\n", num1, num2, Division(num1,num2))
		} else {
			fmt.Println("Cannot divide by zero")
		}
	}
}

func Addition(num1, num2 float32) float32{
	return num1+num2
}
func Subtraction(num1,num2 float32) float32{
	return num1-num2
}
func Multiplication(num1,num2 float32) float32{
	return num1*num2
}
func Division(num1,num2 float32) float32{
	return num1/num2	
}

func clearScreen(){
	waitForEnter()
	var cmd *exec.Cmd
	if runtime.GOOS == "windows"{
		cmd = exec.Command("cmd","/c", "cls")
	} else{
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}