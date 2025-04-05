package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, _ := os.Open("sample.txt")
	content, _ := ioutil.ReadAll(f)
	fmt.Println(string(content))
}
