package parsing

import (
	"fmt"
	"parsing-template/model"
	"strings"
)

func Parsing(template string, ticket model.Ticket, date, timeStr string) string {
	// Replace simple fields
	result := template
	result = strings.ReplaceAll(result, "{Ticket.Terminal}", ticket.Terminal)
	result = strings.ReplaceAll(result, "{LoginUser}", ticket.Cashier)
	result = strings.ReplaceAll(result, "{Date}", date)
	result = strings.ReplaceAll(result, "{Time}", timeStr)
	result = strings.ReplaceAll(result, "{Ticket.PaymentDate}", ticket.PaymentDate.Format("2006-01-02"))
	result = strings.ReplaceAll(result, "{Ticket.PaymentTime}", ticket.PaymentDate.Format("15:04:05"))
	result = strings.ReplaceAll(result, "{Ticket.Tag!Pax}", fmt.Sprintf("%d", ticket.Tag.Pax))

	// Replace Payments
	paymentStr := ""
	for _, payment := range ticket.Payments {
		paymentStr += fmt.Sprintf("Tendered    :   %s\n", payment.Name)
		paymentStr += fmt.Sprintf("Change      :   %s\n", payment.Tendered)
		paymentStr += fmt.Sprintf("RefNo       :   %d\n", payment.PaymentInformation.RefNo)
	}
	result = strings.ReplaceAll(result, "#Ticket.Payments#", paymentStr)
	result = strings.ReplaceAll(result, "Tendered    :   {Payments.Name}", "")
	result = strings.ReplaceAll(result, "Change      :   {Payments.Tendered}", "")
	result = strings.ReplaceAll(result, "RefNo       :   {Payments.PaymentInformation!RefNo}", "")

	// Replace Orders
	orderStr := ""
	for _, order := range ticket.Orders {
		orderStr += fmt.Sprintf("Name %s [=%s] [=%s]\n", order.Name, order.Quantity, order.Price)
	}
	result = strings.ReplaceAll(result, "#Ticket.Orders#", orderStr)

	return result
}

// c: center
// 10: size
// f: full
// j: justifier
