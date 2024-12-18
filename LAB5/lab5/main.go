package main

import (
	"fmt"
	"math"
)

func CalculateRoots(a, b, c float64) (float64, float64, error) {
	D := b*b - 4*a*c
	if D < 0 {
		return 0, 0, fmt.Errorf("уравнение не имеет действительных корней")
	}
	x1 := (-b + math.Sqrt(D)) / (2 * a)
	x2 := (-b - math.Sqrt(D)) / (2 * a)
	return x1, x2, nil
}

func main() {
	var a, b, c float64
	fmt.Println("Введите коэффициенты квадратного уравнения ax^2 + bx + c = 0:")
	fmt.Print("a = ")
	fmt.Scanln(&a)
	fmt.Print("b = ")
	fmt.Scanln(&b)
	fmt.Print("c = ")
	fmt.Scanln(&c)

	x1, x2, err := CalculateRoots(a, b, c)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Корни уравнения: x1 = %.2f, x2 = %.2f\n", x1, x2)
	}
}
