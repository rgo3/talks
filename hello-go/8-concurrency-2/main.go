package main

import (
	"fmt"
	"time"
)

func main() {
	select {
	case n := <-warpDrive():
		fmt.Println(n)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Sadly no 100 ms! :(")
	}
}

func warpDrive() chan string {
	c := make(chan string, 0)
	go func() {
		time.Sleep(250 * time.Millisecond)
		c <- "100 ms? YES WE CAN! :)"
	}()
	return c
}
