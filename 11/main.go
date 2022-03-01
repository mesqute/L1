package main

import (
	"fmt"
	"sort"
)

/*Задание:
Реализовать пересечение двух неупорядоченных множеств.
*/

func main() {
	// создаем множества
	a := []int{8, 489, 5, 15, 74, 25, 46, 98}
	b := []int{0, 25, 7, 956, 4, 8, 3}
	var c []int
	// пересечение с использованием мап (хеш-таблиц) (сложность алгоритма O(a+b)
	c = Hash(a, b)
	// пересечение с использованием сортировки и параллельным обходом по множествам
	// (сложность алгоритма O(a+b) + две сортировки Шелла (от O(n log^2 n) до O(n^2) каждая)
	//c = SortComp(a, b)

	fmt.Println("Пересечение:", c)
}

func Hash(a, b []int) []int {
	// создаем хеш-таблицу (мапу)
	m := make(map[int]bool)
	// создаем слайс для вывода
	var c []int
	// заполняем таблицу (ключи таблицы это значения множества)
	for _, i := range a {
		m[i] = true
	}

	// обходим значения второго множества и проверяем, есть ли в мапе значение с таким ключом
	for _, i := range b {
		if _, ok := m[i]; ok {
			// если есть, то добавляем в возвращаемый слайс
			c = append(c, i)
		}
	}
	return c
}

func SortComp(a, b []int) []int {
	// сортируем оба множества встроенными методами
	sort.Ints(a)
	sort.Ints(b)

	// создаем слайс для вывода
	var c []int

	// создаем переменные для индексов
	var i, j int

	// цикл прохода по обоим множествам
	for {
		// если индексы за пределами множества, то значит пересечений больше нет
		if i < len(a) && j < len(b) {
			// сдвиг на 1 шаг во втором множестве если элемент первого больше элемента второго
			if a[i] > b[j] {
				j++
				continue
			}
			// сдвиг на 1 шаг в первом множестве если элемент второго больше элемента первого
			if a[i] < b[j] {
				i++
				continue
			}
			// если они равны, то добавляем значение в возвращаемый слайс и сдвигаемся в обоих множествах
			c = append(c, a[i])
			i++
			j++
			continue
		}
		break
	}
	return c
}
