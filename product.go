package winit

func (this *Winit) RegisterProduct(param *RegisterProductParam) (results *RegisterProductResults, err error) {
	err = this.doRequest(k_WINIT_API_TYPE_AD, "POST", param, &results)
	return results, err
}

func (this *Winit) GetProductList(param *ProductListParam) (results *ProductListResults, err error) {
	err = this.doRequest(k_WINIT_API_TYPE_OPEN_API, "GET", param, &results)
	return results, err
}