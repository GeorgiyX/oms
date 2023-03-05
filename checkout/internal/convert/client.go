package convert

import (
	"route256/checkout/internal/model"
	descProductService "route256/checkout/pkg/product-service"
	descLoms "route256/loms/pkg/loms"
)

func ToProduct(in *descProductService.GetProductResponse) *model.Product {
	return &model.Product{
		Name:  in.GetName(),
		Price: in.GetPrice(),
	}
}

func ToStocks(in *descLoms.StocksResponse) []model.Stock {
	out := make([]model.Stock, 0, len(in.GetItems()))
	for _, stock := range in.GetItems() {
		out = append(out, model.Stock{
			WarehouseID: stock.GetWarehouseId(),
			Count:       stock.GetCount(),
		})
	}

	return out
}

func ToCreateOrderRequest(user int64, items []model.CreateOrderItem) *descLoms.CreateOrderRequest {
	out := &descLoms.CreateOrderRequest{
		User:  user,
		Items: make([]*descLoms.CreateOrderRequest_Item, 0, len(items)),
	}

	for _, item := range items {
		out.Items = append(out.Items, &descLoms.CreateOrderRequest_Item{
			Sku:   item.SKU,
			Count: item.Count,
		})
	}

	return out
}
