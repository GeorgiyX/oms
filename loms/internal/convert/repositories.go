package convert

import "route256/loms/internal/model"

func ToStocksItemInfo(in []model.Warehouse) []model.StocksItemInfo {
	out := make([]model.StocksItemInfo, 0, len(in))
	for _, item := range in {
		out = append(out, model.StocksItemInfo{
			WarehouseID: item.WarehouseID,
			Count:       item.AvailableToOrder,
		})
	}

	return out
}

func ToOrderItemsDB(orderID int64, in []model.OrderItemToCreate) []model.OrderItemDB {
	out := make([]model.OrderItemDB, 0, len(in))
	for _, item := range in {
		out = append(out, model.OrderItemDB{
			Sku:         item.SKU,
			Count:       item.Count,
			OrderInfoID: orderID,
		})
	}

	return out
}
