package main

import "fmt"

func BinaireDecimal(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			result = result*2 + 1
		} else if s[i] == '0' {
			result = result * 2
		} else {
			return -1
		}
	}
	return result
}

func main() {
	fmt.Println(BinaireDecimal("101"))
	fmt.Println(BinaireDecimal("1001"))
}
