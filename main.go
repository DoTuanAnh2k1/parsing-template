package main

import (
	"fmt"
	"parsing-template/data"
	"parsing-template/model"
	"parsing-template/parsing"
)

func main() {
	dataTicket := model.NewData()
	fmt.Println(dataTicket)
	parsingRes := parsing.Parsing(data.Template, dataTicket, "01-01-1970", "Monday")
	fmt.Println(parsingRes)
}
