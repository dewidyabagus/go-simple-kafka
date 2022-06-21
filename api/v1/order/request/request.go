package request

import "learn/kafka/business/order"

type (
	NewOrder struct {
		TransactionNo string `json:"transaction_no"`
		Items         []struct {
			ItemID    uint    `json:"item_id"`
			ItemPrice float64 `json:"item_price"`
			Qty       uint    `json:"qty"`
		}
		Date string `json:"date"`
	}
)

func (n *NewOrder) ToBusinessNewOrder() *order.NewOrder {
	response := &order.NewOrder{
		TransactionNo: n.TransactionNo,
		Date:          n.Date,
	}
	for _, item := range n.Items {
		response.Items = append(response.Items, order.Item{
			ItemID:    item.ItemID,
			ItemPrice: item.ItemPrice,
			Qty:       item.Qty,
		})
	}
	return response
}
