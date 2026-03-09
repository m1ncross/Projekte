package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Println("Alle reäle Zahlen mit dem Punkt antippenen, z.B. 1.5 oder 0.25")
	//time.Sleep(2 * time.Second)

	fmt.Print("Die erste Variable eintippen:")
	fmt.Scanln(&a)

	fmt.Print("Die zweite Variable eintippen:")
	fmt.Scanln(&b)

	fmt.Print("Die dritte Variable eintippen:")
	fmt.Scanln(&c)

	fmt.Println(a, b, c)
	calculate(a, b, c)
}

func calculate(a, b, c float64) (bool, float64, float64) {
	D := b*b - 4*a*c

	if D < 0 {
		fmt.Println("Die Gleichung hat keine reäle Lösung")
		return false, 0, 0
	} else if D == 0 {
		x := -b / (2 * a)
		fmt.Printf("Die Gleichung hat eine doppelte reäle Lösung: x = %.2f\n", x)
		return true, x, x
	} else {
		x1 := (-b + math.Sqrt(D)) / (2 * a)
		x2 := (-b - math.Sqrt(D)) / (2 * a)
		fmt.Printf("Die Gleichung hat zwei reäle Lösungen: x1 = %.2f, x2 = %.2f\n", x1, x2)
		return true, x1, x2
	}
}
