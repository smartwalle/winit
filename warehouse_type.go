package winit

// WarehouseListParam https://developer.winit.com.cn/document/detail/id/43.html
type WarehouseListParam struct {
}

func (this *WarehouseListParam) Action() string {
	return "queryWarehouse"
}

type WarehouseListResults struct {
	Code int                  `json:"code"`
	Msg  string               `json:"msg"`
	Data []*WarehouseListData `json:"data"`
}

type WarehouseListData struct {
	WarehouseCode    string `json:"warehouseCode"`
	WarehouseName    string `json:"warehouseName"`
	WarehouseID      int    `json:"warehouseID"`
	WarehouseAddress string `json:"warehouseAddress"`
}

// WarehouseStorageListParam http://developer.winit.com.cn/document/detail/id/44.html
type WarehouseStorageListParam struct {
	IsActive    string `json:"isActive"`
	PageNum     int    `json:"pageNum"`
	PageSize    int    `json:"pageSize"`
	WarehouseID string `json:"warehouseID"`
}

func (this *WarehouseStorageListParam) Action() string {
	return "queryWarehouseStorage"
}

type WarehouseStorageListResults struct {
	Code int                         `json:"code"`
	Msg  string                      `json:"msg"`
	Data []*WarehouseStorageListData `json:"data"`
}

type WarehouseStorageListData struct {
	ReservedInventory string  `json:"reservedInventory"`
	ProductCode       string  `json:"productCode"`
	Inventory         string  `json:"inventory"`
	PipelineInventory string  `json:"pipelineInventory"`
	Specification     string  `json:"specification"`
	WarehouseID       string  `json:"warehouseID"`
	ProductName       string  `json:"productName"`
	ProductWeight     float32 `json:"productWeight"`
	ProductLength     float32 `json:"productLength"`
	ProductHeight     float32 `json:"productHeight"`
	ProductWidth      float32 `json:"productWidth"`
}

// DeliveryWayListParams http://developer.winit.com.cn/document/detail/id/45.html
type DeliveryWayListParam struct {
	WarehouseID string `json:"warehouseID"`
}

func (this *DeliveryWayListParam) Action() string {
	return "queryDeliveryWay"
}

type DeliveryWayListResults struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data []*DeliveryWayListData `json:"data"`
}

type DeliveryWayListData struct {
	DeliveryID            int    `json:"deliveryID"`
	IsMandoorplateNumbers string `json:"isMandoorplateNumbers"`
	DeliveryWay           string `json:"deliveryWay"`
	WarehouseID           int    `json:"warehouseID"`
	WinitProductCode      string `json:"winitProductCode"`
}

// CreateOutboundInfoParam http://developer.winit.com.cn/document/detail/id/50.html
type CreateOutboundInfoParam struct {
	WarehouseID      string             `json:"warehouseID"`
	Repeatable       string             `json:"repeatable"`
	DeliveryWayID    string             `json:"deliveryWayID"`
	InsuranceTypeID  string             `json:"insuranceTypeID"`
	SellerOrderNo    string             `json:"sellerOrderNo"`
	RecipientName    string             `json:"recipientName"`
	PhoneNum         string             `json:"phoneNum"`
	ZipCode          string             `json:"zipCode"`
	EmailAddress     string             `json:"emailAddress"`
	State            string             `json:"state"`
	Region           string             `json:"region"`
	City             string             `json:"city"`
	Address1         string             `json:"address1"`
	Address2         string             `json:"address2"`
	DoorplateNumbers string             `json:"doorplateNumbers"`
	ProductList      []*OutBoundProduct `json:"productList"`
}

func (this *CreateOutboundInfoParam) Action() string {
	return "createOutboundInfo"
}

type OutBoundProduct struct {
	ProductCode string `json:"productCode"`
	ProductNum  int    `json:"productNum"`
}

type CreateOutboundInfoResult struct {
	Code int                     `json:"code"`
	Msg  string                  `json:"msg"`
	Data *CreateOutboundInfoData `json:"data"`
}

type CreateOutboundInfoData struct {
	OutboundOrderNum string `json:"outboundOrderNum"`
}

// CreateOutboundOrderParam http://developer.winit.com.cn/document/detail/id/49.html
type CreateOutboundOrderParam struct {
	WarehouseID      string             `json:"warehouseID"`
	Repeatable       string             `json:"repeatable"`
	DeliveryWayID    string             `json:"deliveryWayID"`
	InsuranceTypeID  string             `json:"insuranceTypeID"`
	SellerOrderNo    string             `json:"sellerOrderNo"`
	RecipientName    string             `json:"recipientName"`
	PhoneNum         string             `json:"phoneNum"`
	ZipCode          string             `json:"zipCode"`
	EmailAddress     string             `json:"emailAddress"`
	State            string             `json:"state"`
	Region           string             `json:"region"`
	City             string             `json:"city"`
	Address1         string             `json:"address1"`
	Address2         string             `json:"address2"`
	DoorplateNumbers string             `json:"doorplateNumbers"`
	ProductList      []*OutBoundProduct `json:"productList"`
}

func (this *CreateOutboundOrderParam) Action() string {
	return "createOutboundOrder"
}

type CreateOutboundOrderResult struct {
	Code int                     `json:"code"`
	Msg  string                  `json:"msg"`
	Data *CreateOutboundOrderData `json:"data"`
}

type CreateOutboundOrderData struct {
	OutboundOrderNum string `json:"outboundOrderNum"`
}

// GetOrderVerdorTrackingParam http://developer.winit.com.cn/document/detail/id/56.html
type GetOrderVerdorTrackingParam struct {
	TrackingNos string `json:"trackingnos"`
}

func (this *GetOrderVerdorTrackingParam) Action() string {
	return "tracking.getOrderVerdorTracking"
}

type GetOrderVerdorTrackingResults struct {
	Code string                        `json:"code"`
	Msg  string                        `json:"msg"`
	Data []*GetOrderVerdorTrackingData `json:"data"`
}

type GetOrderVerdorTrackingData struct {
	OrderNo     string   `json:"orderNo"`
	TrackingNoa string   `json:"trackingNo"`
	Origin      string   `json:"origin"`
	Destination string   `json:"destination"`
	PickupMode  string   `json:"pickupMode"`
	Status      string   `json:"status"`
	VendorName  string   `json:"vendorName"`
	Trace       []*Trace `json:"trace"`
}

type Trace struct {
	Type             string `json:"type"`
	Date             string `json:"date"`
	Location         string `json:"location"`
	LastEvent        string `json:"lastEvent"`
	EventCode        string `json:"eventCode"`
	EventStatus      string `json:"eventStatus"`
	EventDescription string `json:"eventDescription"`
	Operator         string `json:"operator"`
}
