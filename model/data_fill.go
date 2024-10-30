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

func NewReceipt() ReceiptLayOut {
	return ReceiptLayOut{
		Header: HeaderReceipt{
			Something:   "Curry Village",
			StoreName:   "Banana leaf p/l",
			Address:     "8 Lim Teck Kim Road",
			PhoneNumber: "6226 2562",
		},
		ReceiptInfo: ReceiptInfo{
			ReceiptNo: 3,
			Date:      time.Now(),
		},
		Entities: Entities{
			EntityName: "B14",
			Header:     [4]string{"Qty", "Items", "Price", "Amount"},
			Orders: []OrderInfo{
				{
					Quantity:    1,
					Name:        "Toasted Bagel Jam",
					Price:       1.50,
					TotalAmount: 1.50,
				},
				{
					Quantity:    1,
					Name:        "Toast and Jam",
					Price:       1.50,
					TotalAmount: 1.50,
				},
				{
					Quantity:    1,
					Name:        "Bacon and Tomato",
					Price:       3.49,
					TotalAmount: 3.49,
				},
				{
					Quantity:    100,
					Name:        "Egg, Bacon Cheese",
					Price:       3.99,
					TotalAmount: 399.00,
				},
			},
			PaymentInfo: PaymentInfor{
				Name:   "Total",
				Amount: 405.49,
			},
		},
	}
}
