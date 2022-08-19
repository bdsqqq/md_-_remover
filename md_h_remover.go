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
	// argsWithProg := os.Args
	// argsWithoutProg := argsWithProg[1:]
	// if len(argsWithoutProg) < 1 {
	// 	log.Fatal("Expected input file as argument")
	// }

	rawMD, err := os.ReadFile("./in/chapters/1.md")
	if err != nil {
		log.Fatal(err);
	}

	sliced := strings.Split(string(rawMD), "\n")

	fmt.Println(string(sliced[11]))
}