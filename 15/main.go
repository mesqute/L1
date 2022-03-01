package main

/*Задание:
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
Приведите корректный пример реализации.
Функция createHugeString() вернет []rune / string, из 1024+ элементов.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

var justString string

func someFunc() {
	// получение слайса rune или string размером 2^10
	v := createHugeString(1 << 10)
	// конвертация в string решает проблему неопределенности получаемых данных
	// ([]rune конвертируется в string, а string так и останется string),
	// а также освободит память от HugeString так как после конвертации
	// будет созданы новые данные в своей области памяти (не будет храниться ссылка на HugeString)
	justString = string(v[:100])
}

func main() {
	someFunc()
}
