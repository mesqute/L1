package main

import (
	"fmt"
	"strings"
)

/*Задание:
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func main() {
	// создаем строку со словами для переворота
	str := "snow dog sun"
	fmt.Println("Строка для переворота:", str)
	// разделяем строку на слайс срок используя символ " " как разделитель
	sstr := strings.Split(str, " ")
	// переворачиваем слайс
	Reverse(sstr)
	// восстанавливаем строку с пробелами
	str = strings.Join(sstr, " ")
	// выводим результат в консоль
	fmt.Println("переворот:", str)
}

// Reverse переворачивает слайс string
func Reverse(data []string) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}
