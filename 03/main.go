package main

import (
	"fmt"
	"sync"
)

func main() {
	// реализация через WaitGroup и Mutex
	WG()

	// реализация через каналы
	CH()
}

func WG() {
	fmt.Println("Использование пакета sync:")

	// инициализируем переменную, хранящую сумму чисел
	var sum int
	var mtx sync.Mutex
	// создаем массив с заданными числами
	var arr = [...]int{2, 4, 6, 8, 10}
	// создаем экземпляр sync.WaitGroup для контроля окончания выполнения группы горутин
	var wg sync.WaitGroup
	// передаем в wg число горутин в группе
	wg.Add(len(arr))
	// перебираем элементы массива и передаем значения в функцию вычисления квадрата числа
	for _, value := range arr {
		// запускаем анонимную функцию в отдельной горутине
		go func(val int) {
			// по окончанию выполнения функции отправляет в wg сигнал,
			// что все необходимые операции были выполнены
			defer wg.Done()
			// блокируем Mutex
			mtx.Lock()
			// записываем сумму во внешнюю (для анонимной функции) переменную
			sum += val * val
			// разблокируем Mutex
			mtx.Unlock()
		}(value)
	}
	// ждем сигналы завершения от всех горутин в группе
	wg.Wait()
	// выводим сумму
	fmt.Println("сумма:", sum)
}

func CH() {
	fmt.Println("Использование каналов:")

	// инициализируем переменную, хранящую сумму чисел
	var sum int
	// создаем массив с заданными числами
	var arr = [...]int{2, 4, 6, 8, 10}
	// создаем буферизированный канал
	intCh := make(chan int, len(arr))
	defer close(intCh)
	// перебираем элементы массива и передаем значения в функцию вычисления квадрата числа
	for _, value := range arr {
		// запускаем функцию в отдельной горутине
		go func(val int) {
			// вычисляемм квадрат
			sq := val * val
			// отправляет в канал данные
			intCh <- sq
		}(value)
	}
	// считываем и суммируем из канала кол-во результатов,
	// равное количеству запущенных горутин
	for i := 0; i < len(arr); i++ {
		sum += <-intCh
	}
	// выводим сумму
	fmt.Println("сумма:", sum)
}