package main

import "fmt"

func printdigits() {
	for i := '0'; i <= '9'; i++ {
		fmt.Print(i)
		fmt.Print(" ")
	}
	fmt.Println()
}
func main() {
	printdigits()
}
