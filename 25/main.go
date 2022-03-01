package main

import (
	"context"
	"fmt"
	"time"
)

/*Задание:
Реализовать собственную функцию sleep.
*/

func main() {
	durat := time.Second * 5

	// Sleep с использованием таймера
	SleepTimer(durat)
	// Sleep с использованием context (по сути то же самое)
	SleepCTX(durat)

}

func SleepTimer(duration time.Duration) {
	// создаем таймер
	timer := time.After(duration)

	// начинаем спать
	fmt.Println("начало SleepTimer")
	// ждем сигнала таймера
	<-timer
	// просыпаемся
	fmt.Println("конец SleepTimer")
}

func SleepCTX(duration time.Duration) {
	// создаем контекст с таймаутом
	ctx, _ := context.WithTimeout(context.Background(), duration)

	// начинаем спать
	fmt.Println("начало SleepCTX")
	// ждем сигнал закрытия контекста
	<-ctx.Done()
	// просыпаемся
	fmt.Println("конец SleepCTX")
}
