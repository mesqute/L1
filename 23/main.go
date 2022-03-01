package main

import "fmt"

/*Задание:
Удалить i-ый элемент из слайса.
*/

func main() {
	// определяем слайс с данными
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(data)
	// задаем индекс удаляемого элемента
	i := 2

	// способ с нарушением порядка (но с фиксированной скоростью)
	data = DirtyDel(data, i)
	fmt.Println("drt out:", data)

	// чистый способ (но с линейной скоростью)
	data = ClearDel(data, i)
	fmt.Println("clr out:", data)
}

func DirtyDel(data []int, i int) []int {
	fmt.Println("drt in:", data)
	// копируем элемент с конца слайса на позицию удаляемого элемента
	data[i] = data[len(data)-1]
	// зануляем конец среза
	data[len(data)-1] = 0
	// усекаем конец среза
	data = data[:len(data)-1]
	return data
}

func ClearDel(data []int, i int) []int {
	fmt.Println("clr in:", data)
	// сдвигаем все элементы справа от i на один шаг влево
	copy(data[i:], data[i+1:])

	// зануляем конец среза
	data[len(data)-1] = 0

	// усекаем конец среза
	data = data[:len(data)-1]
	return data
}
