package main

import (
	"fmt"
	"reflect"
)

/*Задание:
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/

func main() {
	// создаем подопытную переменную
	var val interface{}
	// создаем переменную дял получения результатов работы определяющих функций
	var str string

	// присваиваем переменной значение с типом int
	val = 5

	str = Switch(val)
	fmt.Println("Switch:", str)
	str = Reflect(val)
	fmt.Println("Reflect:", str)
	str = FMT(val)
	fmt.Println("FMT:", str)

	// присваиваем переменной значение с типом string
	val = "str"

	str = Switch(val)
	fmt.Println("Switch:", str)
	str = Reflect(val)
	fmt.Println("Reflect:", str)
	str = FMT(val)
	fmt.Println("FMT:", str)

	// присваиваем переменной значение с типом bool
	val = true

	str = Switch(val)
	fmt.Println("Switch:", str)
	str = Reflect(val)
	fmt.Println("Reflect:", str)
	str = FMT(val)
	fmt.Println("FMT:", str)

	// присваиваем переменной значение с типом chan int
	val = make(chan int)

	str = Switch(val)
	fmt.Println("Switch:", str)
	str = Reflect(val)
	fmt.Println("Reflect:", str)
	str = FMT(val)
	fmt.Println("FMT:", str)

}

func Switch(obj interface{}) string {
	// создаем переменную для вывода
	var str string
	// используем switch по типу элемента
	switch obj.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	case chan string:
		return "chan string"
	case chan bool:
		return "chan bool"
	}

	return str
}

func Reflect(obj interface{}) string {
	// создаем переменную для вывода
	var str string
	// используем reflect для получения типа объекта
	str = reflect.TypeOf(obj).String()
	return str
}

func FMT(obj interface{}) string {
	// создаем переменную для вывода
	var str string
	// используем форматирование пакета fmt
	str = fmt.Sprintf("%T", obj)
	return str

}
