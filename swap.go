package main

import "fmt"

//it makes a double return
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("Valle", "Lorena")
	fmt.Println(a, b)
}
