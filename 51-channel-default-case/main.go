// ASSIGNMENT
//
// Like all good back-end engineers, we frequently save backup snapshots of the
// Mailio database.
//
// Complete the saveBackups function.
//
// It should read values from the snapshotTicker and saveAfter channels
// simultaneously and continuously.
//
// - If a value is received from snapshotTicker, call takeSnapshot()
// - If a value is received from saveAfter, call saveSnapshot() and return from
//   the function: you're done.
// - If neither channel has a value ready, call waitForData() and then
//   time.Sleep() for 500 milliseconds. After all, we want to show in the logs
//   that the snapshot service is running.

package main

import (
	"time"
)

func saveBackups(snapshotTicker, saveAfter <-chan time.Time, logChan chan string) {
	for {
		select {
		case <-snapshotTicker:
			takeSnapshot(logChan)
		case <-saveAfter:
			saveSnapshot(logChan)
			return
		default:
			waitForData(logChan)
			// Если просто передать 500 как в js, то time.Sleep будет как будто игнорироваться
			// То есть это не будет блокировать goroutine
			// ❌ time.Sleep(500)
			time.Sleep(time.Millisecond * 500)
		}
	}
}

// don't touch below this line

func takeSnapshot(logChan chan string) {
	logChan <- "Taking a backup snapshot..."
}

func saveSnapshot(logChan chan string) {
	logChan <- "All backups saved!"
	close(logChan)
}

func waitForData(logChan chan string) {
	logChan <- "Nothing to do, waiting..."
}
