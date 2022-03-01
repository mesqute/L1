package main

import (
	"fmt"
	"math/big"
)

/*Задание:
Разработать программу, которая перемножает, делит,
складывает, вычитает две числовых переменных a, b, значение которых > 2^20.
*/

func main() {

	// вычисление с применением math/big
	Big()
}

func Big() {
	fmt.Println("Big")
	// объявляем переменные и их значение
	a := new(big.Int)
	a.SetString("180000000000000000000000000000000", 10) // 18 * 10^30
	b := new(big.Int)
	b.SetString("56000000000000000000000000000000000", 10) // 56 * 10^32
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	// сложение
	sum := new(big.Int)
	sum.Add(a, b)
	fmt.Println("a + b:", sum)

	// вычитание
	dif := new(big.Int)
	dif.Sub(b, a)
	fmt.Println("b - a:", dif)

	// умножение
	mult := new(big.Int)
	mult.Mul(a, b)
	fmt.Println("a * b:", mult)

	// деление
	div := new(big.Int)
	div.Div(b, a)
	fmt.Println("b / a:", div)

}
