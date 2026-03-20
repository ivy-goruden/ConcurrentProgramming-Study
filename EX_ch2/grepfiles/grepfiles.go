package main

// Expand the program you wrote in the first exercise so that instead of printing
// the contents of the text files, it searches for a string match. The string to search
// for is the first argument on the command line. When you spawn a new goroutine, instead of printing the file’s contents, it should read the file and search for
// a match. If the goroutine finds a match, it should output a message saying that
// the filename contains a match. Call the program grepfiles.go. Here’s how you
// can execute this Go program (“bubbles” is the search string in this example):
// go run grepfiles.go bubbles txtfile1 txtfile2 txtfile3

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func proccess_file(filename string, search string) {
	dat, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("File %s is invalid\n", filename)
		return
	}
	converted := string(dat)
	if strings.Contains(converted, search) {
		fmt.Printf("%s contains %s\n\n", filename, search)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <search> <file1> <file2> ...\n", os.Args[0])
		return
	}
	args := os.Args
	search := args[1]
	args = args[2:]
	for i := 0; i < len(args); i++ {
		go proccess_file(args[i], search)
	}
	time.Sleep(2 * time.Second)

}
