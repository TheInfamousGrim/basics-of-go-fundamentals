package main

import (
	"fmt"
	"os"

	fileutils "frontendmasters.com/go/files/file-management"
)

func main() {
	dir, _ := os.Getwd()

	filePath := dir + "/data/text.txt"

	contents, err := fileutils.ReadTextFile(filePath)

	if (err != nil) {
		fmt.Printf("Error Panic!!! %v", err)
	} else {
		fmt.Println(contents)

		newContent := fmt.Sprintf("Original: %v\n\nDouble Original: %v, %v", contents, contents, contents)
		fileutils.WriteToFile(dir + "/data/output.txt", newContent)
	}
}