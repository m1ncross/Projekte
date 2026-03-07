package main

import "fmt"

func main() {
	var a, b, c float64

	fmt.Print("Die erste Variable eintippen:")
	fmt.Scanln(&a)

	fmt.Print("Die zweite Variable eintippen:")
	fmt.Scanln(&b)

	fmt.Print("Die dritte Variable eintippen:")
	fmt.Scanln(&c)

	fmt.Print(a, b, c)
}
