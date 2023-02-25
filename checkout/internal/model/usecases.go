package model

type CartItem struct {
	Sku   uint32
	Count uint16
	Name  string
	Price uint32
}

type Cart struct {
	Items      []*CartItem
	TotalPrice uint32
}
