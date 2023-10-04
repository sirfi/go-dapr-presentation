package models

import (
	"encoding/json"

	"github.com/google/uuid"
)

type OrderDto struct {
	Id     uuid.UUID  `json:"id"`
	Name   string     `json:"name"`
	Sku    string     `json:"sku"`
	Price  float64    `json:"price"`
	ShipTo AddressDto `json:"shipTo"`
	BillTo AddressDto `json:"billTo"`
}

func (order OrderDto) MarshalBinary() ([]byte, error) {
	return json.Marshal(order)
}

func (order *OrderDto) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, order)
}
