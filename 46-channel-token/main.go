// ASSIGNMENT
//
// Our Mailio server isn't able to boot up until it receives the signal that
// its databases are all online, and it learns about them being online by
// waiting for tokens (empty structs) on a channel.
//
// Run the code. It never exits! The channel passed to waitForDBs stays
// blocked.
//
// Fix the waitForDBs function. It should block until it receives numDBs
// tokens on the dbChan channel. Each time waitForDBs reads a token, the
// getDatabasesChannel goroutine will print a message to the console for you.

package main

import (
	"fmt"
)

// numDBs - количество сигналов, которые нужно получить
// через канал до момента завершения выполнения функции
func waitForDBs(numDBs int, dbChan chan struct{}) {
	// Если не конвертить в time.Duration, то возникает ошибка:
	// compiler: invalid operation: time.Millisecond * numDBs (mismatched types time.Duration and int)
	for dbCount := 0; dbCount < numDBs; dbCount++ {
		// Блокируем выполнение кода
		// До той поры пока не будет получено что угодно в канал
		<-dbChan
	}
}

// Чтобы посмотреть как это работает в связке с waitForDBs
// смотри файл main_test.go в текущей папке
func getDBsChannel(numDBs int) (chan struct{}, *int) {
	count := 0
	ch := make(chan struct{})

	go func() {
		for i := 0; i < numDBs; i++ {
			// пушим в channel токены
			// пустые структуры это унарные значения - токены
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
			count++
		}
	}()

	return ch, &count
}
