package main

import "fmt"

func Hello(input string)string{
	if input == ""{
		return "Hello, World"
	}
	return "Hello, " + input
}

func main() {
	fmt.Println(Hello("hello"))
}

