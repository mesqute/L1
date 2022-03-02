package main

import "fmt"

/*Задание:
Поменять местами два числа без создания временной переменной.*/

func main() {
	var a, b = 1, 2
	fmt.Println("a =", a, "b =", b)

	// средствами языка
	a, b = b, a // a = 2, b = 1
	fmt.Println("a =", a, "b =", b)

	// простой алгоритм
	a = a + b // a = 2 + 1 = 3
	b = a - b // b = 3 - 2 = 1
	a = a - b // a = 3 - 1 = 2
	fmt.Println("a =", a, "b =", b)

	// тот же алгоритм, но в более красивой форме
	a = -a - b // a = -1 - 2 = -3
	b = -a - b // b = -(-3) - 2 = 1
	a = -a - b // a = -(-3) - 1 = 2
	fmt.Println("a =", a, "b =", b)

	// побитовое исключающее или
	a = a ^ b // a = 01 ^ 10 = 11 = 3
	b = a ^ b // b = 11 ^ 10 = 01 = 1
	a = a ^ b // a = 11 ^ 01 = 10 = 2

	fmt.Println("a =", a, "b =", b)
}
