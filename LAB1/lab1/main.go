package main

import (
	"fmt"
	"math"
)

func Triangle(a, b, c float64) bool {
	return a+b > c && a+c > b && b+c > a
}

func Area(a, b, c float64) float64 {
	p := (a + b + c) / 2
	return math.Sqrt(p * (p - a) * (p - b) * (p - c))
}

func main() {
	var a, b, c float64

	fmt.Print("Введите длину первой стороны: ")
	fmt.Scan(&a)
	fmt.Print("Введите длину второй стороны: ")
	fmt.Scan(&b)
	fmt.Print("Введите длину третьей стороны: ")
	fmt.Scan(&c)

	if Triangle(a, b, c) {
		fmt.Println("Треугольник существует.")
		area := Area(a, b, c)
		fmt.Printf("Площадь треугольника: %.2f\n", area)
	} else {
		fmt.Println("Треугольник с такими сторонами не существует.")
	}
}
