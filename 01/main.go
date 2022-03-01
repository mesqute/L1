package main

import "fmt"

/*Задание:
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action
от родительской структуры Human (аналог наследования).*/

// Human это родительская структура, обладающая свойствами и методами
type Human struct {
	Age,
	Weight int
}

// Action это структура паразит, использующая методы и свойства структуры Human
type Action struct {
	Human
}

// Eat повышает вес Human
func (h *Human) Eat() {
	h.Weight++
}

// Wait повышает возраст Human
func (h *Human) Wait() {
	h.Age++
}

func main() {
	// инициализация структуры паразита
	action := Action{Human{}}
	// вывод данных, не принадлежащих структуре
	fmt.Println("Возраст:", action.Age, "Вес:", action.Weight)
	// выполнение действий
	action.Eat()
	action.Wait()
	// повторный вывод для демонстрации изменений
	fmt.Println("Возраст:", action.Age, "Вес:", action.Weight)

}
