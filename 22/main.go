package main

import "fmt"

/*Задание:
Разработать программу, которая перемножает, делит,
складывает, вычитает две числовых переменных a, b, значение которых > 2^20.
*/

func main() {
	// при значениях меньших чем 2^63

	pow := 60
	var a int64 = 1 << pow
	var b int64 = 2 << pow
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("a + b:", a+b)
	fmt.Println("b - a:", b-a)
	fmt.Println("a * b:", a*b)
	fmt.Println("b / a:", b/a)

	// при значениях 2^64 и больше
}
