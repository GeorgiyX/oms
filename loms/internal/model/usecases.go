package model

import "fmt"

type StocksItemInfo struct {
	WarehouseID int64
	Count       uint64
}

type OrderItem struct {
	SKU   uint32
	Count uint32
}

type OrderStatus int16

const (
	Invalid OrderStatus = iota
	New
	Failed
	AwaitingPayment
	Payed
	Cancelled
)

func (s OrderStatus) String() string {
	switch s {
	case New:
		return "New"
	case AwaitingPayment:
		return "AwaitingPayment"
	case Failed:
		return "Failed"
	case Payed:
		return "Payed"
	case Cancelled:
		return "Cancelled"
	}
	return fmt.Sprintf("invalid (%d)", s)
}

type Order struct {
	Status  OrderStatus
	User    int64
	Items   []OrderItem
	OrderID int64
}

type OrderItemToCreate struct {
	SKU   uint32
	Count uint32
}
