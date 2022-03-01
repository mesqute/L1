package main

import "fmt"

/*Задание:
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

type MapSet struct {
	data map[string]bool
}

func main() {
	items := []string{"cat", "cat", "dog", "cat", "tree"}

	// объявляем переменную интерфейса
	var s MapSet

	fmt.Println("Последовательность строк:", items)
	s.Insert(items...)
	fmt.Println("Созданное множество:", s.GetAll())
	fmt.Println("Есть ли элемент 'sun'?", s.Check("sun"))
	fmt.Println("Есть ли элемент 'dog'?", s.Check("dog"))
	s.Delete("dog")
	fmt.Println("Есть ли элемент 'dog' после удаления?", s.Check("dog"))

}

// Реализация через map

// Insert добавляет в множество новые элементы (или не добавляет если они там уже есть)
func (s *MapSet) Insert(strs ...string) {
	for _, str := range strs {
		// проверка проинициализирована ли map
		if s.data == nil {
			s.data = make(map[string]bool)
		}
		// если элемент уже был добавлен, то ничего не изменится,
		// если не был, то он добавится (особенность использования map в основе множества)
		s.data[str] = true
	}
}

// Check проверят наличие элемента во множестве
func (s *MapSet) Check(str string) bool {
	_, ok := s.data[str]
	return ok
}

// Delete удаляет элемент из множества
func (s *MapSet) Delete(str string) {
	// проверка существования элемента
	if s.Check(str) {
		// удаление
		delete(s.data, str)
	}
}

// GetAll возвращает слайс всех элементов множества
func (s *MapSet) GetAll() []string {
	// создаем возвращаемый слайс
	var str []string
	// обходим все элементы map и наполняем слайс
	for item, _ := range s.data {
		str = append(str, item)
	}
	return str
}
