package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Thanks https://stackoverflow.com/a/59961623
func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		 return nil, err
	}
	return os.Create(p)
}

func main() {
	argsWithProg := os.Args
	argsWithoutProg := argsWithProg[1:]
	if len(argsWithoutProg) < 1 {
		log.Fatal("Expected input file as argument")
	}

	// I know how many chapters we have and the name of the files is the number of the chapter soooo...
	for i := 1; i <= 53; i++ {
		rawMD, err := os.ReadFile(fmt.Sprintf("./%s/chapters/%d.md", argsWithoutProg[0], i))
		if err != nil {
			log.Fatal(err);
		}
		sliced := strings.Split(string(rawMD), "\n")
		fmt.Println(string(sliced[11]))
		// fuck line 12, all my homies hate line 12.
		newRawMD := strings.Join(sliced[:11], "\n") + strings.Join(sliced[12:], "\n")

		f, err := create(fmt.Sprintf("./%s/chapters/%d.md", argsWithoutProg[0], i))
		if err != nil {
			log.Fatal(err)
			return
		}
		defer f.Close()
		_, err2 := f.WriteString(newRawMD)
		if err2 != nil {
			log.Fatal(err2)
			return
		}
	}
}