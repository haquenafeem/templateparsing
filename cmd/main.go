package main

import (
	"fmt"

	"github.com/haquenafeem/templateparsing"
)

func main() {
	data := struct{ Name string }{Name: "Anexa"}

	fmt.Println(templateparsing.ParseTemplate(templateparsing.FileLocation, data))
}
