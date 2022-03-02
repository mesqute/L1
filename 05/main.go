package main

import (
	"context"
	"fmt"
	"time"
)

/*Задание:
Разработать программу, которая будет последовательно отправлять значения в канал,
а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

func main() {
	var num int
	// считываем из stdin время работы программы в секундах
	fmt.Print("Введите кол-во кол-во секунд: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println(err)
		return
	}
	// инициализируем канал для сигнала завершения работы
	doneCh := make(chan int)

	// выполнение задания с использованием контекста
	//go CTX(num, doneCh)
	// выполнение задания с использованием таймера
	//go Timer(num, doneCh)

	// тикающий каждую секунду цикл, ждущий окончания работы подпроцессов
loop:
	for {
		select {
		case <-time.Tick(time.Second):
			fmt.Println("Тик...")
		case <-doneCh:
			break loop
		}
	}
}

func CTX(num int, doneCh chan int) {
	// создаем контекст с таймером самоуничтожения (вызова функции cancel) в num секунд
	ctx, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(num))
	// инициализируем канал
	intCh := make(chan int)
	// запускаем писателя в отдельной горутине
	go WriterCTX(intCh, ctx)
	// цикл чтения из канала
loop1:
	for {
		select {
		//читаем из канала и проверяем открыт ли он
		case f, ok := <-intCh:
			// если открыт, то выводим данные в stdout
			if ok {
				fmt.Println("CTX:", f)
				break
			}
			// если закрыт, то выходим из цикла чтения
			fmt.Println("ReaderCTX is GONE")
			break loop1
		}
	}
	// посылаем сигнал о завершении работы функции закрытием канала
	close(doneCh)
}

func WriterCTX(intChan chan int, ctx context.Context) {
	var x, y = 1, 0

	// цикл записи в канал
loop2:
	for {
		// получаем значение Фибоначчи
		y, x = x, y+x
		select {
		// пишем значение в канал
		case intChan <- y:
			// ждем пол секунды для уменьшения захламляемости stdout
			time.Sleep(time.Millisecond * 500)
		case <-ctx.Done():
			// если контекст был закрыт, то закрываем канал и разрываем цикл записи
			fmt.Println("WriterCTX is GONE")
			close(intChan)
			break loop2
		}
	}
}

func Timer(num int, doneCh chan int) {
	// инициализируем канал
	intCh := make(chan int)
	// запускаем писателя в отдельной горутине
	go WriterTimer(intCh, num)
	// цикл чтения из канала
loop1:
	for {
		select {
		//читаем из канала и проверяем открыт ли он
		case f, ok := <-intCh:
			// если открыт, то выводим данные в stdout
			if ok {
				fmt.Println("Timer:", f)
				break
			}
			// если закрыт, то выходим из цикла чтения
			fmt.Println("ReaderTimer is GONE")
			break loop1
		}
	}
	// посылаем сигнал о завершении работы функции закрытием канала
	close(doneCh)

}

func WriterTimer(intChan chan int, num int) {
	var x, y = 1, 0
	timer := time.After(time.Second * time.Duration(num))
	// цикл записи в канал
loop2:
	for {
		// получаем значение Фибоначчи
		y, x = x, y+x
		select {
		// пишем значение в канал
		case intChan <- y:
			// ждем пол секунды для уменьшения захламляемости stdout
			time.Sleep(time.Millisecond * 500)

		case <-timer:
			// после полчения сигнала от таймера закрываем канал и выходим из цикла записи
			fmt.Println("WriterTimer is GONE")
			close(intChan)
			break loop2
		}
	}
}
