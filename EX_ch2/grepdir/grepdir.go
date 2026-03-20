package main

// Change the program you wrote in the second exercise so that instead of passing
// a list of text filenames, you pass a directory path. The program will look inside
// this directory and list the files. For each file, you can spawn a goroutine that will
// search for a string match (the same as before). Call the program grepdir.go.
// Here’s how you can execute this Go program:
// go run grepdir.go bubbles ../../commonfiles
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
		fmt.Printf("Usage: go run . <search> <dir> \n")
		return
	}
	args := os.Args
	search := args[1]
	folder := args[2]
	if strings.LastIndex(folder, "/") != (len(folder) - 1) {
		folder += "/"
	}
	entries, err := os.ReadDir(folder)
	if err != nil {
		fmt.Printf("Dir is invalid\n")
		return
	}
	for i := 0; i < len(entries); i++ {
		if entries[i].IsDir() {
			continue
		}
		go proccess_file(folder+entries[i].Name(), search)
	}
	time.Sleep(2 * time.Second)

}
