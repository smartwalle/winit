package winit

func (this *Winit) GetWarehouseList(param *WarehouseListParam) (results *WarehouseListResults, err error) {
	err = this.doRequest(k_WINIT_API_TYPE_AD, "POST", param, &results)
	return results, err
}

func (this *Winit) GetWarehouseStorageList(param *WarehouseStorageListParam) (results *WarehouseStorageListResults, err error) {
	err = this.doRequest(k_WINIT_API_TYPE_AD, "POST", param, &results)
	return results, err
}

func (this *Winit) GetDeliveryWayList(param *DeliveryWayListParam) (results *DeliveryWayListResults, err error) {
	err = this.doRequest(k_WINIT_API_TYPE_AD, "POST", param, &results)
	return results, err
}

func (this *Winit) CreateOutboundInfo(param *CreateOutboundInfoParam) (results *CreateOutboundInfoResult, err error) {
	err = this.doRequest(k_WINIT_API_TYPE_AD, "POST", param, &results)
	return results, err
}

func (this *Winit) CreateOutboundOrder(param *CreateOutboundOrderParam) (results *CreateOutboundOrderResult, err error) {
	err = this.doRequest(k_WINIT_API_TYPE_AD, "POST", param, &results)
	return results, err
}

func (this *Winit) GetOrderVerdorTracking(param *GetOrderVerdorTrackingParam) (results *GetOrderVerdorTrackingResults, err error) {
	err = this.doRequest(k_WINIT_API_TYPE_OPEN_API, "POST", param, &results)
	return results, err
}