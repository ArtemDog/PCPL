package main

import (
	"fmt"
)

func primeFactors(n int) []int {
	var factors []int

	for n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}

	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}

	if n > 2 {
		factors = append(factors, n)
	}

	return factors
}

func main() {
	var num int

	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	if num > 1 {
		factors := primeFactors(num)
		fmt.Printf("Простые множители числа %d: %v\n", num, factors)
	} else {
		fmt.Println("Введите число больше 1.")
	}
}
