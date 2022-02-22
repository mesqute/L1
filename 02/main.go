package main

import (
	"fmt"
	"sync"
)

func main() {
	// реализация через WaitGroup
	WG()

	// реализация через каналы
	CH()
}

func WG() {
	fmt.Println("Использование пакета sync:")

	// создаем массив с заданными числами
	var arr = [...]int{2, 4, 6, 8, 10}
	// создаем экземпляр sync.WaitGroup для контроля окончания выполнения группы горутин
	var wg sync.WaitGroup
	// передаем в wg число горутин в группе
	wg.Add(len(arr))
	// перебираем элементы массива и передаем значения в функцию вычисления квадрата числа
	for _, val := range arr {
		// запускаем функцию в отдельной горутине
		go SqrWG(&wg, val)
	}
	// ждем сигналы завершения от всех горутин в группе
	wg.Wait()

}

func SqrWG(wg *sync.WaitGroup, val int) {
	// по окончанию выполнения функции отправляет в wg сигнал,
	// что все необходимые операции были выполнены
	defer wg.Done()
	// выводит в stdout оформленный результат вычисления квадрата val
	fmt.Println(val, "*", val, "=", val*val)
}

func CH() {
	fmt.Println("Использование каналов:")

	// создаем массив с заданными числами
	var arr = [...]int{2, 4, 6, 8, 10}
	// создаем буферизированный канал выполняющий роль синхронизатора
	intCh := make(chan int, len(arr))
	defer close(intCh)
	// перебираем элементы массива и передаем значения в функцию вычисления квадрата числа
	for _, val := range arr {
		// запускаем функцию в отдельной горутине
		go SqrCH(intCh, val)
	}
	// ждем завершения работы горутин путем считывания из канала кол-ва результатов,
	// равное количеству запущенных горутин
	for i := 0; i < len(arr); i++ {
		<-intCh
	}
}

func SqrCH(ch chan int, val int) {
	// выводит в stdout оформленный результат вычисления квадрата val
	fmt.Println(val, "*", val, "=", val*val)
	// отправляет в канал данные
	ch <- 0
}
