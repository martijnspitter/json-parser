package main

import (
	"fmt"
	"json-parser/cli"
	"json-parser/parser"
	"json-parser/reader"
)

func main() {
	cli := cli.NewCLI()
	cli.Run()
	path := cli.GetPath()

	newReader := reader.NewReader(path)
	err := newReader.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	content := newReader.GetContent()

	tokenizer := parser.NewTokenizer(content)
	parser := parser.NewParser(tokenizer)

	result, err := parser.Parse()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Parsed JSON: %v\n", result)
	}

}
