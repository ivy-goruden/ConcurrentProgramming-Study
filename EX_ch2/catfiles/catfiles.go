package main

// Write a program similar to the one in listing 2.3 that accepts a list of text filenames as arguments. For each filename, the program should spawn a new
// goroutine that will output the contents of that file to the console. You can use
// the time.Sleep() function to wait for the child goroutines to complete (until
// you know how to do this better). Call the program catfiles.go. Here’s how you
// can execute this Go program:
// go run catfiles.go txtfile1 txtfile2 txtfile3

import (
	"fmt"
	"os"
	"time"
)

func proccess_file(filename string) {
	dat, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("File %s is invalid\n", filename)
		return
	}
	fmt.Printf("%s\n\n", dat)
}

func main() {
	args := os.Args
	args = args[1:]
	for i := 0; i < len(args); i++ {
		go proccess_file(args[i])
	}
	time.Sleep(2 * time.Second)

}
