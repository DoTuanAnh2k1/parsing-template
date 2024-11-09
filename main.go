package main

import (
	"parsing-template/model"
	"parsing-template/parsing"
)

func main() {
	dataTicket := model.NewTicket()
	inputFile := "data/template/template.txt"
	outputFile := "data/output/output.txt"
	err := parsing.ParsingFromFileToFile(inputFile, outputFile, dataTicket)
	if err != nil {
		panic(err)
	}
}
