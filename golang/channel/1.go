package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	go printSomthing(("Hello world !"), done)
	go printSomthing(("Hello world for golang!"), done)
	go time.Sleep(3 * time.Second)
	go printSomthing(("Hello world take time !"), done)
	select{
		case 
	}
	<-done
	<-done
	<-done
}

func printSomthing(greet string, doneChan chan bool) string {
	fmt.Println(greet)
	doneChan <- true
	return greet
}
