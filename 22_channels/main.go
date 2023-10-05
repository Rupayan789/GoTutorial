package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := &sync.WaitGroup{}
	// mutex := &sync.Mutex{}

	// nums := []int8{0}
	// wg.Add(3)
	// go func(wg *sync.WaitGroup, mutex *sync.Mutex) {
	// 	mutex.Lock()
	// 	nums = append(nums, 1)
	// 	mutex.Unlock()
	// 	defer wg.Done()
	// }(wg, mutex)
	// go func(wg *sync.WaitGroup, mutex *sync.Mutex) {
	// 	mutex.Lock()
	// 	nums = append(nums, 2)
	// 	mutex.Unlock()
	// 	defer wg.Done()
	// }(wg, mutex)
	// go func(wg *sync.WaitGroup, mutex *sync.Mutex) {
	// 	mutex.Lock()
	// 	nums = append(nums, 3)
	// 	mutex.Unlock()
	// 	defer wg.Done()
	// }(wg, mutex)
	// wg.Wait()
	// fmt.Println(nums)

	ch := make(chan int, 2)

	wg.Add(2)
	//Send Only Channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("Producer started")
		time.Sleep(3 * time.Second)
		ch <- 5
		close(ch)
		fmt.Println("Producer ended")

	}(ch, wg)

	//Receive Only Channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("Consumer started")
		val, isOpen := <-ch
		fmt.Printf("Is open value :%v and value from channel:%v\n", isOpen, val)
		val2, isOpen2 := <-ch
		fmt.Printf("Is open value :%v and value from channel:%v\n", isOpen2, val2)
		fmt.Println("Consumer ended")
	}(ch, wg)
	wg.Wait()

}
