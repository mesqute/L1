package main

import (
	"bufio"
	"fmt"
	"os"
)

/*Задание:
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

func main() {
	// используем bufio.Scanner для считывания строки вместе с пробелами
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// создаем переменную сохраняющую ввод из консоли
		var str string
		// считываем данные из консоли
		fmt.Print("Введите строку для переворота: ")
		scanner.Scan()
		str = scanner.Text()

		// преобразуем строку в слайс rune
		strr := []rune(str)

		// переворачиваем слайс rune
		Reverse(strr)

		// конвертируем слайс обратно в строку
		str = string(strr)
		// выводим полученный результат
		fmt.Println("переворот:", str)
	}

}

// Reverse переворачивает слайс rune
func Reverse(data []rune) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
