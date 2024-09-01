package grtn

import (
	"fmt"
	"sync"
	"time"
)

var dbdata = []string{"id1", "id2", "id3", "id4", "id5"}
var result = []string{}
var wg = sync.WaitGroup{}
var m = sync.RWMutex{}

func GoRoutineExmaple() {

	start := time.Now()
	for i := 0; i < len(dbdata); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\n Total time : %v", time.Since(start))
	fmt.Printf("\n Result time : %v", result)
}
func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Microsecond)
	save(dbdata[i])
	log()
	wg.Done()
}
func save(data string) {
	m.Lock()
	result = append(result, data)
	m.Unlock()
}
func log() {
	m.RLock()
	fmt.Printf("\n the current result %v", result)
	m.RUnlock()
}

func count() {
	var res int
	for i := 0; i < 100000000; i++ {
		res += 1
	}
	wg.Done()
}
func GoRoutineExmaple2() {

	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go count()
	}
	wg.Wait()
	fmt.Printf("\n Total time : %v", time.Since(start))
}
