package main

import "fmt"

func printalphabet() {
	for i := 'a'; i <= 'z'; i++ {
		fmt.Printf("%c", i)
	}
}
func main() {
	printalphabet()
	fmt.Println()
}
