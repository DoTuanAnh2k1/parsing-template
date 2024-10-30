package main

import (
	"parsing-template/model"
	"parsing-template/utils"

	"github.com/mect/go-escpos"
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

	p, err := escpos.NewUSBPrinterByPath("") // empry string will do a self discovery
	if err != nil {
		panic(err)
	}

	p.Init()
	p.PrintLn(receiptStr)
	p.Cut()
	p.End()
	// // Write buffer content to a file for review
	// if err := os.WriteFile("data/output/output.txt", buf.Bytes(), 0644); err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Print simulation saved to output.txt")
}
