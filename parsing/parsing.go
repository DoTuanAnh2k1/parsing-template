package parsing

import (
	"parsing-template/model"
	"parsing-template/utils"
	"strings"
	"time"
)

func ParsingFromFileToFile(filePathInput, filePathOutput string, ticket model.Ticket) {
	template := ParsingFromFile(filePathInput, ticket)
	err := utils.WriteToFile(filePathOutput, template)
	utils.CheckError(err)
}

func ParsingFromFile(filePath string, ticket model.Ticket) string {
	template, err := utils.ReadFromFile(filePath)
	utils.CheckError(err)
	return Parsing(template, ticket)
}

func ParsingToFile(filePathOutput string, template string, ticket model.Ticket) {
	template = Parsing(template, ticket)
	err := utils.WriteToFile(filePathOutput, template)
	utils.CheckError(err)
}

func Parsing(template string, ticket model.Ticket) string {
	timeNow := time.Now().Format("2006-01-02 15:04:23")
	parts := strings.Split(timeNow, " ")
	return ParsingWithDate(template, ticket, parts[0], parts[1])
}

func ParsingWithDate(template string, ticket model.Ticket, date, timeStr string) string {
	template = replaceValue(ticket, template, date, timeStr)
	template = findFormatDecimal(template)
	template = findFormateDate(template)
	return removeEmptyLines(template)
}
