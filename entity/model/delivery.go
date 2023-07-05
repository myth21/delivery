package model

type Delivery struct {
	ID                 int64   `json:"id"`
	ReceiverName       string  `json:"receiver_name"`
	ReceiverPhone      string  `json:"receiver_phone"`
	ReceiverAddress    string  `json:"receiver_address"`
	DateTimeFrom       string  `json:"date_time_from"`
	DateTimeTo         string  `json:"date_time_to"`
	DeliveredAt        *string `json:"delivered_at"`
	NonDeliveredReason *string `json:"non_delivered_reason"`
	Comment            string  `json:"comment"`
	Price              int     `json:"price"`
	Status             int     `json:"status"`
	CreatedAt          string  `json:"created_at"`
	UpdatedAt          *string `json:"updated_at"`
}

func (delivery Delivery) GetStatusName() string {
	if delivery.Status == 1 {
		return string(OrderPendingStatus)
	}
	if delivery.Status == 2 {
		return string(OrderDeliveredStatus)
	}
	if delivery.Status == 3 {
		return string(DeliveryStatusFailed)
	}
	return "Undefined"
}

func (delivery *Delivery) SetPendingStatus() {
	delivery.Status = 1
}

func (delivery *Delivery) SetFailedStatus() {
	delivery.Status = 3
}

// DeliveryStatus is like enum
type DeliveryStatus string

const (
	OrderPendingStatus   DeliveryStatus = "Order Pending"
	OrderDeliveredStatus                = "Order Delivered"
	DeliveryStatusFailed                = "Order Failed"
)

// TODO see Stringer package solution
