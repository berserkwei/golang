package main

import (
	"fmt"
	"time"
)

func main() {
	ticks := time.Tick(time.Second)
	count := 10
	for count > 0 {
		<-ticks
		count--
		fmt.Println("I'm making log here! " + time.Now().String())
	}
}
