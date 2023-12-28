package main

import "fmt"

func main(){
	var a, b int
	fmt.Print("Please input first number:")
	fmt.Scan(&a)
	fmt.Print("Please input second number:")
	fmt.Scan(&b)
	fmt.Printf("%d + %d = %d\n",a, b, a+b)
}