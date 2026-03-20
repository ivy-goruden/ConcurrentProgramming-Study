package main

// 4 Adapt the program in the third exercise to continue searching recursively in
// any subdirectories. If you give your search goroutine a file, it should search for a
// string match in that file, just like in the previous exercises. Otherwise, if you
// give it a directory, it should recursively spawn a new goroutine for each file or
// directory found inside. Call the program grepdirrec.go, and execute it by running this command:
// go run grepdirrec.go bubbles ../../commonfiles
import (
	"fmt"
	"os"
	"strings"
	"time"
)

func proccess_file(file *os.File, search string) {
	var data []byte
	data = make([]byte, 1000)
	_, err := file.Read(data)
	if err != nil {
		fmt.Printf("File %s is invalid\n", file.Name())
		return
	}
	converted := string(data)
	if strings.Contains(converted, search) {
		fmt.Printf("%s contains %s\n\n", file.Name(), search)
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
	file, err := os.Open(folder)
	if err != nil {
		fmt.Printf("File is invalid\n")
		return
	}
	fileInfo, err := file.Stat()
	switch fileInfo.IsDir() {
	case true:
		iterateFolder(file, search)
	case false:
		proccess_file(file, search)
	}

	time.Sleep(2 * time.Second)

}

func iterateFolder(folder *os.File, search string) {

	entries, err := folder.ReadDir(-1)
	if err != nil {
		fmt.Printf("Dir is invalid\n")
		return
	}
	for i := 0; i < len(entries); i++ {
		path := folder.Name() + string(os.PathSeparator) + entries[i].Name()
		file, err := os.Open(path)
		if err != nil {
			continue
		}
		if entries[i].IsDir() {
			go iterateFolder(file, search)
		} else {
			go proccess_file(file, search)
		}
	}
}
