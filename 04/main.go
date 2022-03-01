package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*Задание:
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

func main() {
	var num int

	// счиываем из stdin кол-во необходимых работяг
	fmt.Print("Введите кол-во работяг: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println(err)
		return
	}
	// инициализируем главный поток
	mainChan := make(chan int, 100)
	defer close(mainChan)

	// инициализируем контекст с функцией отмены
	ctx, cancel := context.WithCancel(context.Background())

	// инициалииуем канал для контроля окончания работы работяг
	doneChan := make(chan bool, num)
	defer close(doneChan)

	// запускаем N работяг
	for i := 0; i < num; i++ {
		go Worker(mainChan, doneChan, ctx, i)
	}
	// запускаем функцию перехвата сигнала о закрытии приложения
	go InterceptionExitSignal(cancel)

	// используем лейблы для назначения циклу названия
loop:
	for {
		select {
		// отправляет в канал случайные чила от 0 до 9
		case mainChan <- rand.Intn(10):
			// небольшое замедление
			time.Sleep(time.Second)
		case <-ctx.Done():
			// выходим из цикла loop если контекст закрыл канал
			break loop
		}
	}

	// ждем пока все работяги не завершат работу
	for i := 0; i < num; i++ {
		<-doneChan
	}

	// завершаем приложение с кодом 0
	os.Exit(0)
}

// Worker считывает данные из канаа и выводит их в stdout
func Worker(mChan <-chan int, doneChan chan<- bool, ctx context.Context, id int) {
	// используем лейблы для назначения циклу названия
loop1:
	for {
		select {
		// считываем данные из канала
		case b := <-mChan:
			fmt.Println("Работяга номер", id, "-", b)
		case <-ctx.Done():
			// выходим из цикла loop1 если контекст закрыл канал
			break loop1
		}
	}
	fmt.Println("Работяга номер", id, "завершил свою работу")
	// отправляем в канал данные, что работяга закончил работу
	doneChan <- true
}

// InterceptionExitSignal перехватывает сигнал о закрытии приложения
// и запускает процесс окончания работы
func InterceptionExitSignal(cancel context.CancelFunc) {
	// инициализируем канал, в который перенаправим сигнал завершения работы программы
	c := make(chan os.Signal, 1)
	// перехватываем сингал
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	// ждем получение этого сигнала
	fmt.Println("Получен сигнал", <-c)
	// вызываем функцию завершения в контексте
	cancel()
}
