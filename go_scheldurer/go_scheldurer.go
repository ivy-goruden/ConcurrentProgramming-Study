package main

import (
	"fmt"
	"runtime"
)

func sayHello() {
	fmt.Println("Hello")
}

func sayGoodbye() {
	fmt.Println("Bye")
}
func main() {
	go sayHello()
	go sayGoodbye()
	runtime.Gosched() //запуск планировщика прежде чем завершится главный тред.
	fmt.Println("Finished")
}
