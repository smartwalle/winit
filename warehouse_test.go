package winit

import (
	"testing"
)

func TestWinit_GetWarehouseList(t *testing.T) {
	//fmt.Println("===== GetWarehouseList =====")
	//pp := &WarehouseListParam{}
	//results, err := client.GetWarehouseList(pp)
	//fmt.Println(err)
	//if results != nil {
	//	for _, w := range results.Data {
	//		fmt.Println(w.WarehouseID, w.WarehouseName, w.WarehouseCode)
	//	}
	//}
}

func TestWinit_GetWarehouseStorageList(t *testing.T) {
	//fmt.Println("===== GetWarehouseStorageList =====")
	//pp := &WarehouseStorageListParam{}
	//pp.WarehouseID = "1000109"
	//pp.PageSize = 10
	//pp.PageNum = 1
	//results, err := client.GetWarehouseStorageList(pp)
	//fmt.Println(err)
	//if results != nil {
	//	for _, w := range results.Data {
	//		fmt.Println(w.WarehouseID, w.ProductName)
	//	}
	//}
}

func TestWinit_GetDeliveryWayList(t *testing.T) {
	//fmt.Println("===== GetDeliveryWayList =====")
	//pp := &DeliveryWayListParam{}
	//pp.WarehouseID = "1000008"
	//results, err := client.GetDeliveryWayList(pp)
	//fmt.Println(err)
	//if results != nil {
	//	for _, w := range results.Data {
	//		fmt.Println(w.WarehouseID, w.DeliveryID, w.DeliveryWay)
	//	}
	//}
}

func TestWinit_CreateOutboundInfo(t *testing.T) {
	//fmt.Println("===== CreateOutboundInfo =====")
	//pp := &CreateOutboundInfoParam{}
	//pp.WarehouseID = "1000008"
	//pp.Repeatable = "N"
	//pp.DeliveryWayID = "1000111"
	//pp.InsuranceTypeID = "1000000"
	//pp.RecipientName = "YF"
	//pp.PhoneNum = "3865622627"
	//pp.ZipCode = "96818"
	//pp.EmailAddress = "aa@qq.com"
	//pp.State = "US"
	//pp.Region = "Hawaii"
	//pp.City = "Honolulu"
	//pp.Address1 = "aaaaa"
	//pp.SellerOrderNo = "test-order3"
	//
	//var pl []*OutBoundProduct
	//var p1 = &OutBoundProduct{}
	//p1.ProductCode = "25663"
	//p1.ProductNum = 10
	//
	//pl = append(pl, p1)
	//pp.ProductList = pl
	//
	//
	//results, err := client.CreateOutboundInfo(pp)
	//fmt.Println(err)
	//if results != nil {
	//	fmt.Println(results.Data.OutboundOrderNum)
	//}
}

func TestWinit_CreateOutboundOrder(t *testing.T) {
	//fmt.Println("===== CreateOutboundOrder =====")
	//pp := &CreateOutboundOrderParam{}
	//pp.WarehouseID = "1000008"
	//pp.Repeatable = "N"
	//pp.DeliveryWayID = "1000111"
	//pp.InsuranceTypeID = "1000000"
	//pp.RecipientName = "YF"
	//pp.PhoneNum = "3865622627"
	//pp.ZipCode = "96818"
	//pp.EmailAddress = "aa@qq.com"
	//pp.State = "US"
	//pp.Region = "Hawaii"
	//pp.City = "Honolulu"
	//pp.Address1 = "aaaaa"
	//pp.SellerOrderNo = "test-order2"
	//
	//var pl []*OutBoundProduct
	//var p1 = &OutBoundProduct{}
	//p1.ProductCode = "25663"
	//p1.ProductNum = 10
	//
	//pl = append(pl, p1)
	//pp.ProductList = pl
	//
	//
	//results, err := client.CreateOutboundOrder(pp)
	//fmt.Println(err)
	//if results != nil {
	//	fmt.Println(results.Data.OutboundOrderNum)
	//}
}

func TestWinit_GetOrderVerdorTracking(t *testing.T) {
	//fmt.Println("===== GetOrderVerdorTracking =====")
	//pp := &GetOrderVerdorTrackingParam{}
	//pp.TrackingNos = "WO0000002370"
	//results, err := client.GetOrderVerdorTracking(pp)
	//fmt.Println(err)
	//if results != nil {
	//	for _, t := range results.Data {
	//		fmt.Println(t.OrderNo, t.Origin, t.Status)
	//		for _, tt := range t.Trace {
	//			fmt.Println(tt.Location, tt.EventDescription, tt.EventStatus)
	//		}
	//	}
	//}
}