package parsing

import (
	"fmt"
	"parsing-template/model"
	"strings"
)

func replaceValue(ticket model.Ticket, template string, date, timeStr string) string {
	// Delete some fields
	template = strings.ReplaceAll(template, "Tendered    :   {Payments.Name}", "")
	template = strings.ReplaceAll(template, "Change      :   {Payments.Tendered}", "")
	template = strings.ReplaceAll(template, "RefNo       :   {Payments.PaymentInformation!RefNo}", "")
	template = strings.ReplaceAll(template, "Name {Orders.Name} [=FormatDecimal({Orders.Quantity},2)] [=FormatDecimal({Orders.Price},2)]", "")

	// Replace simple fields of ticket
	template = strings.ReplaceAll(template, "{Ticket.Terminal}", ticket.Terminal)
	template = strings.ReplaceAll(template, "{LoginUser}", ticket.Cashier)
	template = strings.ReplaceAll(template, "{Date}", date)
	template = strings.ReplaceAll(template, "{Time}", timeStr)
	template = strings.ReplaceAll(template, "{Ticket.PaymentDate}", ticket.PaymentDate.Format("2006-01-02"))
	template = strings.ReplaceAll(template, "{Ticket.PaymentTime}", ticket.PaymentDate.Format("15:04:05"))
	template = strings.ReplaceAll(template, "{Ticket.Tag!Pax}", fmt.Sprintf("%d", ticket.Tag.Pax))
	template = strings.ReplaceAll(template, "{Tikcet.Tag|PaxTime}", ticket.Tag.PaxTime.Format("2006-01-02 15:04:05"))

	// Replace Payments
	paymentStr := "##Ticket.Payments##\n"
	for _, payment := range ticket.Payments {
		paymentStr += fmt.Sprintf("Tendered    :   %s\n", payment.Name)
		paymentStr += fmt.Sprintf("Change      :   %s\n", payment.Tendered)
		paymentStr += fmt.Sprintf("RefNo       :   %d\n", payment.PaymentInformation.RefNo)
		paymentStr += "\n"
	}
	template = strings.ReplaceAll(template, "##Ticket.Payments##", paymentStr)

	// Replace Orders
	orderStr := "##Ticket.Orders##\n"
	for _, order := range ticket.Orders {
		orderStr += fmt.Sprintf("Name %s [=FormatDecimal(%s,2)] [=FormatDecimal(%s,2)]\n", order.Name, order.Quantity, order.Price)
	}
	template = strings.ReplaceAll(template, "##Ticket.Orders##", orderStr)
	return template
}
