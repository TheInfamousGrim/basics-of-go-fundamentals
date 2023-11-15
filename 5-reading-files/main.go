package main

import (
	"fmt"
	"os"

	fileutils "frontendmasters.com/go/files/file-management"
)

func main() {
	dir, _ := os.Getwd()

	contents, err := fileutils.ReadTextFile(dir + "/data/text.txt")

	if (err != nil) {
		fmt.Printf("Error Panic!!! %v", err)
	} else {
		fmt.Println(contents)
	}
}