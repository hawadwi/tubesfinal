package mq

type DeliveryEvent struct {
	Resi          string `json:"resi"`
	WarehouseZone string `json:"warehouse_zone"`
	CourierID    int    `json:"courier_id"`
	AssignedZone string `json:"assigned_zone"`
}