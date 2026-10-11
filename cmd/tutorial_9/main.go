package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.RWMutex{}
var wg = sync.WaitGroup{} // These basically work as counters
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)    //Increments counter
		go dbCall(i) // using the go keyword, we can spawn in tasks in the background to execute func calls
	}
	wg.Wait() // Waits for counter to go down to 0
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nThe results are: %v\n", results)
}

func dbCall(i int) {
	// Simulate DB call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
	save(dbData[i])
	log()
	wg.Done() // Decrements counter
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock() // Checks if there is a full lock (Lock()) if so, it can not access.
	// Otherwise, it allows other goroutines to go in and read but not to write
	fmt.Printf("The current results are: %v\n", results)
	m.RUnlock()
}
