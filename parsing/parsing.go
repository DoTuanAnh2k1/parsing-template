package parsing

import (
	"fmt"
	"parsing-template/model"
	"strings"
)

func Parsing(template string, ticket model.Ticket, date, timeStr string) string {
	result := template

	// Delete some fields
	result = strings.ReplaceAll(result, "Tendered    :   {Payments.Name}", "")
	result = strings.ReplaceAll(result, "Change      :   {Payments.Tendered}", "")
	result = strings.ReplaceAll(result, "RefNo       :   {Payments.PaymentInformation!RefNo}", "")
	result = strings.ReplaceAll(result, "Name {Orders.Name} [=FormatDecimal({Orders.Quantity},2)] [=FormatDecimal({Orders.Price},2)]", "")

	// Replace simple fields of ticket
	result = strings.ReplaceAll(result, "{Ticket.Terminal}", ticket.Terminal)
	result = strings.ReplaceAll(result, "{LoginUser}", ticket.Cashier)
	result = strings.ReplaceAll(result, "{Date}", date)
	result = strings.ReplaceAll(result, "{Time}", timeStr)
	result = strings.ReplaceAll(result, "{Ticket.PaymentDate}", ticket.PaymentDate.Format("2006-01-02"))
	result = strings.ReplaceAll(result, "{Ticket.PaymentTime}", ticket.PaymentDate.Format("15:04:05"))
	result = strings.ReplaceAll(result, "{Ticket.Tag!Pax}", fmt.Sprintf("%d", ticket.Tag.Pax))
	result = strings.ReplaceAll(result, "{Tikcet.Tag|PaxTime}", ticket.Tag.PaxTime.Format("2006-01-02 15:04:05"))

	// Replace Payments
	paymentStr := ""
	for _, payment := range ticket.Payments {
		paymentStr += fmt.Sprintf("Tendered    :   %s\n", payment.Name)
		paymentStr += fmt.Sprintf("Change      :   %s\n", payment.Tendered)
		paymentStr += fmt.Sprintf("RefNo       :   %d\n", payment.PaymentInformation.RefNo)
		paymentStr += "\n"
	}
	result = strings.ReplaceAll(result, "##Ticket.Payments##", paymentStr)

	// Replace Orders
	orderStr := ""
	for _, order := range ticket.Orders {
		orderStr += fmt.Sprintf("Name %s [%d] [%0.2f]\n", order.Name, order.Quantity, order.Price)
	}
	result = strings.ReplaceAll(result, "##Ticket.Orders##", orderStr)

	return result
}

// c: center
// 10: size
// f: full
// j: justifier
