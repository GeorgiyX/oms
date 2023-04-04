package convert

import "route256/loms/internal/model"

func ToStatusChangeKafka(in model.StatusChange) model.StatusChangeKafka {
	return model.StatusChangeKafka{
		OrderID: in.OrderID,
		Status:  int16(in.Status),
	}
}
