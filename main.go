package main

import (
	"fmt"
	"os"
)

func saveTextToFile(file_path string, content string) {

	file, err := os.OpenFile(file_path, os.O_CREATE|os.O_RDWR, 0755)

	if err != nil {
		fmt.Println("Error opening file.")
		file, err = os.Create(file_path)
		if err != nil {
			fmt.Println("Error creating file.")
			return
		}
	}

	file.Truncate(0)

	_, err = file.WriteString(content)

	if err != nil {
		fmt.Println("Error writing to file.")
		return
	}
}

func main() {
	saveTextToFile("./test.txt", "Hello World!")
}
