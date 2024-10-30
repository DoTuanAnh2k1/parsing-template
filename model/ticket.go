package model

import "time"

type Ticket struct {
	Terminal    string
	Cashier     string // Doen't require
	PaymentDate time.Time
	PaymentType string
	Tag         Tag
	Payments    []Payment
	Orders      []Order
}

type Tag struct {
	Pax     int
	PaxTime time.Time
}

type Payment struct {
	Name               string
	Tendered           string
	PaymentInformation PaymentInfo
}

type PaymentInfo struct {
	RefNo   int
	RefTime time.Time
}

type Order struct {
	Name     string
	Quantity string
	Price    string
}
