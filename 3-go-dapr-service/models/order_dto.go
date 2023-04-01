package models

import "github.com/google/uuid"

type OrderDto struct {
	Id     uuid.UUID  `json:"id"`
	Name   string     `json:"name"`
	Sku    string     `json:"sku"`
	Price  float64    `json:"price"`
	ShipTo AddressDto `json:"shipTo"`
	BillTo AddressDto `json:"billTo"`
}
