package main

import (
	"fmt"
	"os"
)

func main() {
	content, _ := os.ReadFile("sample.txt")
	fmt.Println(string(content))
}
