package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Println("Введите коэффициенты квадратного уравнения ax^2 + bx + c = 0:")
	fmt.Print("a = ")
	fmt.Scanln(&a)
	fmt.Print("b = ")
	fmt.Scanln(&b)
	fmt.Print("c = ")
	fmt.Scanln(&c)

	D := b*b - 4*a*c

	if D < 0 {
		fmt.Println("Уравнение не имеет действительных корней")
		return
	}

	x1 := (-b + math.Sqrt(D)) / (2 * a)
	x2 := (-b - math.Sqrt(D)) / (2 * a)

	fmt.Printf("Корни уравнения: x1 = %.2f, x2 = %.2f\n", x1, x2)
}
