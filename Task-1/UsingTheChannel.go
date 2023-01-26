package main

import "fmt"

var global int

func main() {
	value := CreateSlice(5000, 1)
	ch := make(chan int)
	go Sum(value, ch)
	fmt.Println("global from chan = ", <-ch)
}

func CreateSlice(len int, elems int) []int {
	slice := []int{}
	for i := 0; i < len; i++ {
		slice = append(slice, elems)
	}
	return slice
}

func Sum(s []int, ch chan int) int {
	for _, value := range s {
		global += value
	}
	ch <- global
	return global
}
