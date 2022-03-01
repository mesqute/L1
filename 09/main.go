package main

import (
	"fmt"
	"sync"
)

/*Задание:
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

func main() {
	// создаем WG для ожидания завершения работы читателя
	var wg sync.WaitGroup
	// создаем каналы
	firstCh, secondCh := make(chan int), make(chan int)
	// создаем массив (x)
	data := [...]int{0, 5, 8, 13, 91}

	// выводим в консоль знания массива
	fmt.Println("Конвейерные числа:", data)

	// создаем писателя
	go func() {
		for _, datum := range data {
			firstCh <- datum
		}
		// закрываем первый канал после передачи всех значений
		close(firstCh)
	}()
	// создаем вычислителя
	go func() {
		for {
			// считываем значения из первого канала пока канал не закроется
			if v, ok := <-firstCh; ok {
				// записываем квадрат полученного значения во второй канал
				secondCh <- v * v
				continue
			}
			// если первый канал закрыт, закрываем второй канал и заканчиваем работу
			close(secondCh)
			break
		}
	}()
	// увеличиваем счетчик ожидаемых сигналов WaitGroup
	wg.Add(1)
	// создаем читателя
	go func() {
		for {
			// считываем значения из второго канала пока канал не закроется
			if v, ok := <-secondCh; ok {
				// выводим полученное значение в stdout
				fmt.Println("Конвейерный квадрат:", v)
				continue
			}
			// если второй канал закрыт, отправляем сигнал о завершении работы и заканчиваем работу
			wg.Done()
			break
		}
	}()
	// ждем завершения работы
	wg.Wait()
}
