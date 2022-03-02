package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*Задание:
Реализовать все возможные способы остановки выполнения горутины.
*/

func main() {

	// методы завершения
	//CTXTimeout()
	//CTXCancel()
	// (по сути то же самое, но без использования контекста)
	//Timer()
	//CH()
	// методы блокировки
	//Mutex()
	//WG()
}

func CTXTimeout() {
	// создаем канал для контроля завершения дочерней функции
	doneCh := make(chan int, 1)
	// объявляем контекст с таймером на 5 сек
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	// запускаем в отдельной горутине функцию с бесконечным циклом
	go func() {
	loop:
		for {
			select {
			// функция каждую секунду отправляет сообщение
			case <-time.Tick(time.Second):
				fmt.Println("я предчувствую скорую смерть от CTXTimeout")
			case <-ctx.Done():
				// при получении сигнала из контекста выходим из цикла и завершает работу
				break loop
			}
		}
		fmt.Println("МЕНЯ УБИЛ CTXTimeout!")
		// отправляем сигнал о завершении работы
		doneCh <- 0
	}()
	// ждем сигнала о завершении работы от созданной горутины
	<-doneCh
}

func CTXCancel() {
	// создаем канал для контроля завершения дочерней функции
	doneCh := make(chan int, 1)
	// объявляем контекст с функцией отмены
	ctx, cancel := context.WithCancel(context.Background())
	// запускаем в отдельной горутине функцию с бесконечным циклом
	go func() {
	loop:
		for {
			select {
			// функция каждую секунду отправляет сообщение
			case <-time.Tick(time.Second):
				fmt.Println("только CTXCancel знает, сколько мне осталось")
			case <-ctx.Done():
				// при получении сигнала из контекста выходим из цикла и завершает работу
				break loop
			}
		}
		fmt.Println("МЕНЯ УБИЛ CTXCancel!")
		// отправляем сигнал о завершении работы
		doneCh <- 0
	}()
	// даем горутине пожить 5 секунд
	time.Sleep(time.Second * 5)
	// вызываем функцию сигнала завершения в контексте
	cancel()
	// ждем сигнала о завершении работы от созданной горутины
	<-doneCh

}

func Timer() {
	// определяем время жизни горутины
	secondsToDie := 5
	// создаем канал для контроля завершения дочерней функции
	doneCh := make(chan int, 1)
	// запускаем в отдельной горутине функцию с бесконечным циклом
	go func() {
		// переменная для подсчета пройденного времени
		var i int
		ticker := time.Tick(time.Second)
		timer := time.After(5 * time.Second)
	loop:
		for {
			select {
			// функция каждую секунду отправляет сообщение
			case <-ticker:
				fmt.Println("о нет, мне осталось всего", secondsToDie-i, "секунд!")
				i++
			case <-timer:
				// при получении сигнала от таймера выходим из цикла и завершаем работу
				break loop
			}
		}
		fmt.Println("МЕНЯ УБИЛ Timer!")
		// отправляем сигнал о завершении работы
		doneCh <- 0
	}()
	// ждем сигнала о завершении работы от созданной горутины
	<-doneCh
}

func CH() {
	// создаем канал для контроля завершения дочерней горутины
	doneCh := make(chan int, 1)
	// создаем канал для контроля работы дочерней горутины
	killCh := make(chan int)
	// запускаем в отдельной горутине функцию с бесконечным циклом
	go func() {
	loop:
		for {
			select {
			// функция каждую секунду отправляет сообщение
			case <-time.Tick(time.Second):
				fmt.Println("я предчувствую скорую смерть от CH")
			case <-killCh:
				// при получении сообщения из канала выходим из цикла и завершаем работу
				break loop
			}
		}
		fmt.Println("МЕНЯ УБИЛ CH!")
		// отправляем сигнал о завершении работы
		doneCh <- 0
	}()
	// даем горутине пожить 5 секунд
	time.Sleep(time.Second * 5)
	// отправляем в сигнальный канал сообщение
	killCh <- 0
	// ждем сигнала о завершении работы от созданной горутины
	<-doneCh

}

func Mutex() {
	// инициализируем Mutex
	var mtx sync.Mutex
	// запускаем в отдельной горутине функцию с бесконечным циклом
	go func() {
		for {
			select {
			// функция каждую секунду отправляет сообщение
			case <-time.Tick(time.Second):
				fmt.Println("я свободная горутина, Mutex меня не достать")
			}
			// на этом моменте срабатывает блокировка, если Mutex уже закрыт в другом месте
			mtx.Lock()
			mtx.Unlock()
		}
	}()
	// даем горутине пожить 5 секунд
	time.Sleep(time.Second * 5)
	// блокируем горутину
	mtx.Lock()
	fmt.Println("Mutex заблокировал горутину")
}

func WG() {
	// инициализируем Mutex
	var wg sync.WaitGroup
	// запускаем в отдельной горутине функцию с бесконечным циклом
	go func() {
		for {
			select {
			// функция каждую секунду отправляет сообщение
			case <-time.Tick(time.Second):
				fmt.Println("я свободная горутина, WaitGroup меня не достать")
			}
			// на этом моменте срабатывает блокировка, если счетчик wg больше 0
			wg.Wait()
		}
	}()
	// даем горутине пожить 5 секунд
	time.Sleep(time.Second * 5)
	// блокируем горутину
	wg.Add(1)
	fmt.Println("WG заблокировал горутину")
}
