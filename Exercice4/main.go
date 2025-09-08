package main

import "fmt"

func printPGCB(nb, nb2 int) int {
	if nb2 == 0 {
		return nb
	}
	return printPGCB(nb2, nb%nb2)
}
func PGCB(nb int) int {
	return printPGCB(nb, 18)
}
func main() {
	fmt.Println(PGCB(48))
}
