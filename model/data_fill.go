package model

import (
	"parsing-template/utils"
	"time"
)

func NewTicket() Ticket {
	return Ticket{
		Terminal:    "Terminal",
		Cashier:     "Chronical Do",
		PaymentDate: time.Now(),
		PaymentType: "cash",
		Tag: Tag{
			Pax:     10,
			PaxTime: time.Now(),
		},
		Payments: []Payment{
			{
				Name:     utils.StringRandom(10),
				Tendered: utils.StringRandom(10),
				PaymentInformation: PaymentInfo{
					RefNo:   int(utils.IntRandom(1, 1000)),
					RefTime: time.Now(),
				},
			},
			{
				Name:     utils.StringRandom(10),
				Tendered: utils.StringRandom(10),
				PaymentInformation: PaymentInfo{
					RefNo:   int(utils.IntRandom(1, 1000)),
					RefTime: time.Now(),
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
		Cashier:     utils.StringRandom(10),
		PaymentDate: time.Now(),
		PaymentType: utils.StringRandom(10),
		Tag: Tag{
			Pax:     int(utils.IntRandom(1, 1000)),
			PaxTime: time.Now(),
		},
		Payments: []Payment{
			{
				Name:     utils.StringRandom(10),
				Tendered: utils.StringRandom(10),
				PaymentInformation: PaymentInfo{
					RefNo:   int(utils.IntRandom(1, 1000)),
					RefTime: time.Now(),
				},
			},
			{
				Name:     utils.StringRandom(10),
				Tendered: utils.StringRandom(10),
				PaymentInformation: PaymentInfo{
					RefNo:   int(utils.IntRandom(1, 1000)),
					RefTime: time.Now(),
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
