package main

import "fmt"

/*Задание:
Реализовать паттерн «адаптер» на любом примере.
*/

type Food interface {
	Eat()
}

type Drinks interface {
	Drink()
}

type Man struct {
}

func (m *Man) Cook(f Food) {
	fmt.Print("Человек начал готовить, ")
	f.Eat()
}

type Rice struct {
}

func (m *Rice) Eat() {
	fmt.Println("а потом съел миску риса")
}

type DrinksAdapter struct {
	DrinksType Drinks
}

func (da *DrinksAdapter) Eat() {
	fmt.Print("налил в кружку, ")
	da.DrinksType.Drink()
}

type Tea struct {
}

func (b *Tea) Drink() {
	fmt.Println("а потом и выпил чай.")
}

func main() {

	// инициализируем структуру "человек"
	man := new(Man)

	// если человек захочет рис, то он его просто приготовит и съест
	rice := new(Rice)
	man.Cook(rice)

	// если человек захочет чай, то он не сможет его съесть,
	// поэтому был добавлен адаптер, в котором человек наливает чай, а потом уже пьет
	tea := new(Tea)
	adapter := &DrinksAdapter{DrinksType: tea}
	man.Cook(adapter)
}
