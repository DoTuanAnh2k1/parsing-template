package model

import (
	"parsing-template/utils"
	"time"
)

func NewTicket() Ticket {
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
				Quantity: utils.StringRandom(10),
				Price:    utils.StringRandom(10),
			},
			{
				Name:     utils.StringRandom(10),
				Quantity: utils.StringRandom(10),
				Price:    utils.StringRandom(10),
			},
		},
	}
}
