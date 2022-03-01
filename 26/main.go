package main

import (
	"fmt"
	"strings"
)

/*Задание:
Разработать программу, которая проверяет, что все символы в строке уникальные
(true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

func main() {
	// создаем переменную для сохранения проверяемой строки
	var str string

	// проверяем различные строки
	str = "abcd"
	fmt.Printf("%s - %v \n", str, UniCharCheckMap(str))
	str = "abCdefAaf"
	fmt.Printf("%s - %v \n", str, UniCharCheckMap(str))
	str = "aabcd"
	fmt.Printf("%s - %v \n", str, UniCharCheckMap(str))

}

func UniCharCheckMap(str string) bool {
	// создаем map[rune] для хранения символов строки
	runes := make(map[rune]bool)

	// переводим все символы в нижний регистр
	str = strings.ToLower(str)
	// обходим строку
	for _, s := range str {
		// если символ уже есть в map, то значит символ не уникальный
		if _, ok := runes[s]; ok {
			return false
		}
		// если символа нет, то добавляем его
		runes[s] = true
	}
	// возвращаем true если не нашло повторяющихся символов
	return true
}
