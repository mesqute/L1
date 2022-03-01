package main

import (
	"math/rand"
	"sync"
	"time"
)

/*Задание:
Реализовать конкурентную запись данных в map.*/

// Data вспомогательная структура
type Data struct {
	id  int
	val int
}

func main() {
	// реализация через каналы
	//CH()
	// реализация через Mutex
	//Mutex()
	// реализация через sync.Map
	//SyncMap()
}

func CH() {
	// задаем кол-во писателей
	num := 5
	// создаем мапу для записи
	data := make(map[int]int)
	// создаем канал для передачи данных
	dataChan := make(chan Data, num)
	// запускаем num горутин писателей
	for i := 0; i < num; i++ {
		go func(ch chan Data) {
			for {
				// пишем в канал структуру с рандомизированными величинами полей
				ch <- Data{id: rand.Int(), val: rand.Int()}
			}
		}(dataChan)
	}

	// бесконечный цикл чтения из канала
	for {
		tmp := <-dataChan
		data[tmp.id] = tmp.val
	}
}

func Mutex() {
	// задаем кол-во писателей
	num := 5
	// создаем мапу для записи
	data := make(map[int]int)
	// создаем Mutex
	var mtx sync.Mutex
	// запускаем num горутин писателей
	for i := 0; i < num; i++ {
		go func(data map[int]int, mtx *sync.Mutex) {
			for {
				id := rand.Int()
				val := rand.Int()
				// блокируем Mutex
				mtx.Lock()
				// записываем данные в data
				data[id] = val
				// разблокируем Mutex
				mtx.Unlock()
			}
		}(data, &mtx)
	}
	time.Sleep(time.Minute)
}

func SyncMap() {
	// задаем кол-во писателей
	num := 5
	// создаем мапу
	var data sync.Map
	// запускаем num горутин писателей
	for i := 0; i < num; i++ {
		go func(data *sync.Map) {
			for {
				// записываем данные в data
				data.Store(rand.Int(), rand.Int())
			}
		}(&data)
	}
	time.Sleep(time.Minute)

}
