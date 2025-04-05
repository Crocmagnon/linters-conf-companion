package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}
