package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
}

/* test defer, FILO
func foo() {
	defer fmt.Println("Done!")
	defer fmt.Println("Are we done?")
	fmt.Println("Doing some stuff, who knows what?")

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
*/

func main() {
	//foo()

	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("There")
	wg.Wait()
	wg.Add(1)
	go say("Hi")
	wg.Wait()
}
