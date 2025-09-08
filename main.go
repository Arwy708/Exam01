package main

import "fmt"

func AlphaCount(s string) int {
	count := 0
	for _, r := range s {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
			count++
		}
	}
	return count
}
func main() {
	fmt.Println(AlphaCount("Hello, World!"))
	fmt.Println(AlphaCount("123 456"))
	fmt.Println(AlphaCount("GoLang1.21"))
}
