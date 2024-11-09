package model

import (
	"parsing-template/utils"
	"time"
)

func NewTicket() Ticket {
	return Ticket{
		Terminal:    "Terminal",
		LoginUser:   "Chronical Do",
		PaymentDate: time.Now().Format("2006-01-02"),
		PaymentType: "cash",
		Tag:         "{\"Pax\": \"100\", \"PaxTime\": \"2020/10/10\"}",
		Payments: []Payment{
			{
				Name:     utils.StringRandom(10),
				Tendered: utils.StringRandom(10),
				PaymentInformation: PaymentInfo{
					RefNo:   "100",
					RefTime: "2020/10/10",
				},
			},
			{
				Name:     utils.StringRandom(10),
				Tendered: utils.StringRandom(10),
				PaymentInformation: PaymentInfo{
					RefNo:   "100",
					RefTime: "2020/10/10",
				},
			},
		},
		Orders: []Order{
			{
				Name:     "Highland Coffee",
				Quantity: "1",
				Price:    "15",
			},
			{
				Name:     "Red Bull",
				Quantity: "1",
				Price:    "18.50",
			},
		},
	}
}

func NewTicketRandom() Ticket {
	return Ticket{
		Terminal:    utils.StringRandom(10),
		LoginUser:   utils.StringRandom(10),
		PaymentDate: time.Now().Format("2006-01-02"),
		PaymentType: utils.StringRandom(10),
		Tag:         "{\"Pax\": \"100\", \"PaxTime\": \"2020/10/10\"}",
		Payments: []Payment{
			{
				Name:     utils.StringRandom(10),
				Tendered: utils.StringRandom(10),
				PaymentInformation: PaymentInfo{
					RefNo:   "100",
					RefTime: "2020/10/10",
				},
			},
			{
				Name:     utils.StringRandom(10),
				Tendered: utils.StringRandom(10),
				PaymentInformation: PaymentInfo{
					RefNo:   "100",
					RefTime: "2020/10/10",
				},
			},
		},
		Orders: []Order{
			{
				Name:     utils.StringRandom(10),
				Quantity: "1",
				Price:    "3.99",
			},
			{
				Name:     utils.StringRandom(10),
				Quantity: "1",
				Price:    "3.99",
			},
		},
	}
}
