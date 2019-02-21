package main

import "fmt"

var ch = make(chan int) 
var steps int = 0

func change() {
	for i := 0; i < 50000; i++ {
		steps += 1
		fmt.Print(steps,"\n")
	}

	ch <- 1
}

// func main() {
// 	go change()
// 	fmt.Println("Changed")
// 	<-ch
// 	fmt.Println("Finished")
// }