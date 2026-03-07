package main

import (
	"fmt"
	"time"
)

func main() {
	var a, b, c float64

	fmt.Println("Alle reäle Zahlen mit dem Punkt antippenen, z.B. 1.5 oder 0.25")
	time.Sleep(2 * time.Second)

	fmt.Print("Die erste Variable eintippen:")
	fmt.Scanln(&a)

	fmt.Print("Die zweite Variable eintippen:")
	fmt.Scanln(&b)

	fmt.Print("Die dritte Variable eintippen:")
	fmt.Scanln(&c)

	fmt.Print(a, b, c)
}
