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
	fmt.Println("Мужик готовит")
	f.Eat()
}

type Meat struct {
}

func (m *Meat) Eat() {
	fmt.Println("Мужик съел мясо")
}

type Beer struct {
}

func (b *Beer) Drink() {
	fmt.Println("Мужик выпил пиво")
}

type DrinksAdapter struct {
	DrinksType Drinks
}

func (da *DrinksAdapter) Eat() {
	fmt.Println("Мужик разгрыз бутылку")
	da.DrinksType.Drink()
}

func main() {
	// инициализируем структуру "мужик"
	man := new(Man)

	// если мужик захочет мяса, то он его просто приготовит и съест
	meat := new(Meat)
	man.Cook(meat)

	// если мужик захочет пива, то он не сможет его съесть,
	// поэтому был добавлен адаптер, в котором мужик разгрызает бутылку, а потом пьет
	beer := new(Beer)
	adapter := &DrinksAdapter{DrinksType: beer}
	man.Cook(adapter)
}
