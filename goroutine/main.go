package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 这种写法保证join操作一定会发生
func main() {
	// var wg sync.WaitGroup
	// sayHello := func() {
	// 	defer wg.Done()
	// 	fmt.Println("Hello")
	// }
	// wg.Add(1)
	// go sayHello()
	// wg.Wait()
	// // expect result: Hello

	// salutation := "hello"
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	salutation = "welcome"
	// }()

	// wg.Wait()
	// fmt.Println(salutation)
	// // expect result: welcome

	// for _, salutation := range []string{"hello", "greetings", "good day"} {
	// 	wg.Add(1)
	// 	go func(salutation string) {
	// 		defer wg.Done()
	// 		fmt.Println(salutation)
	// 	}(salutation)
	// }
	// wg.Wait()

	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg2 sync.WaitGroup
	// goroutine never exit, until the process is finished
	noop := func() { wg2.Done(); <-c }
	// the number of goroutines to create.
	const numGoroutines = 1e4
	wg2.Add(numGoroutines)
	before := memConsumed() // amount of memory consumed before creating our gorountines
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg2.Wait()
	after := memConsumed() // amount of memory consumed after creating our goroutines.
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)

}

// func sayHello() {
// 	fmt.Println("Hello")
// }

// alter function
