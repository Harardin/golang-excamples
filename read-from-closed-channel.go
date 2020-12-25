package main

import (
	"fmt"
)

func main() {

	/*
		If read from closed chanel it will return empty valu
		- int will = 0
		- struct will equal {}
	*/

	c := make(chan int, 1)
	fin := make(chan bool)

	go func() {
		for {
			num, done := <-c
			if done {
				fmt.Println("Num ", num)
			} else {
				fmt.Println("Done")
				return
			}
		}
	}()

	for i := 0; i <= 8; i++ {
		c <- i
	}
	close(c)

	<-fin
}
