package main

import "fmt"

func main() {
	var a, b, c float64

	fmt.Print("Введите первое число:")
	fmt.Scanln(&a)

	fmt.Print("Введите второе число:")
	fmt.Scanln(&b)

	fmt.Print("Введите третье число:")
	fmt.Scanln(&c)

	fmt.Print(a, b, c)
}
