package main

import "fmt"

// exemplo de função privada
func Add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("O resultado de 10 + 20 é", Add(10, 20))
}
