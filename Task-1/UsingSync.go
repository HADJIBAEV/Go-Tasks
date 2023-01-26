package main

import (
	"fmt"
	"sync"
)

var global int

func main() {

	var Wg sync.WaitGroup
	var m sync.Mutex

	value := CreateSlice(50000, 1)
	value2 := CreateSlice(50000, 2)

	Wg.Add(1)
	go Sum(value, &m, &Wg)
	Wg.Add(1)
	go Sum(value2, &m, &Wg)
	Wg.Wait()
	fmt.Println("After 2 go GoRutine = ", global)

}

func CreateSlice(len int, elems int) []int {
	slice := []int{}
	for i := 0; i < len; i++ {
		slice = append(slice, elems)
	}
	return slice
}
func Sum(s []int, mu *sync.Mutex, wg *sync.WaitGroup) int { //ch chan int

	for _, value := range s {
		mu.Lock()
		global += value
		mu.Unlock()
	}

	wg.Done()
	return global
}
