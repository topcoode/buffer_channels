package main

import (
	"fmt"
	"sync"
)

func main() {
	data := make(chan string, 4)
	var wg sync.WaitGroup
	wg.Add(4)
	go Routine1(data, &wg)
	go Routine2(data, &wg)
	go Routine3(data, &wg)
	go Routine4(data, &wg)
	wg.Wait()
	//signal completion by sending a value on the channel
	<-data
	fmt.Println("data is binary...", data)

}
func Waitlock() {
	var wl sync.Mutex
	wl.Lock()
	defer wl.Unlock()

}
func Routine1(data chan string, wg *sync.WaitGroup) {
	fmt.Println("go routine1:")
	defer wg.Done()
	data <- "go routine"
}
func Routine2(data chan string, wg *sync.WaitGroup) {
	fmt.Println("go routine2:")
	defer wg.Done()
	//Waitlock()
	data <- "go routine"
}
func Routine3(data chan string, wg *sync.WaitGroup) {
	fmt.Println("go routine3:")
	wg.Done()
	//Waitlock()
	data <- "go routine"
}
func Routine4(data chan string, wg *sync.WaitGroup) {
	fmt.Println("go routine4:")
	wg.Done()
	//Waitlock()
	data <- "go routine"
}
