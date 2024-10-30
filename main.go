package main

import (
	"parsing-template/model"
	"parsing-template/utils"
)

func main() {
	// dataTicket := model.NewData()
	// fmt.Println(dataTicket)
	// parsingRes := parsing.Parsing(data.Template, dataTicket, "01-01-1970", "Monday")
	// fmt.Println(parsingRes)
	// for i := 0; i < 52; i++ {
	// 	fmt.Printf(" ")
	// }

	dataReceipt := model.NewReceipt()
	receiptStr := dataReceipt.Output()
	err := utils.WriteToFile("data/output/receipt.txt", receiptStr)
	if err != nil {
		panic(err)
	}
	// fmt.Println(dataReceipt.Output())
}
