package main

import (
	"fmt"
	"path"
)

func main() {
	var dir, file string

	dir, file = path.Split("css/main.css")
	// _, file = path.Split("css/main.css") If only desired return value is the file name.

	fmt.Println("dir: ", dir)
	fmt.Println("file: ", file)
}
