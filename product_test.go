package winit

import (
	"testing"
	"fmt"
)

func TestWinit_RegisterProduct(t *testing.T) {
	pp := &RegisterProductParam{}

	p := &Product{}
	p.ProductCode = "EAKKK004"
	p.ChineseName = "PowerBuyerDEF456"
	p.EnglishName = "PowerBuyerDEF456"
	p.CategoryOne = "1001"
	p.CategoryTwo = "1SDDW"
	p.CategoryThree = "categoryThree"
	p.RegisteredWeight = 2
	p.FixedVolumeWeight = "Y"
	p.RegisteredLength = 4
	p.RegisteredWidth = 4
	p.RegisteredHeight = 4
	p.Branded = "Y"
	p.BrandedName = "HHG"
	p.Model = "HG"
	p.DisplayPageUrl = "https://www.baidu.com"
	p.Remark = "e"
	p.ExportCountry = "CN"
	p.InporCountry = "AU"
	p.InportDeclaredValue = 33
	p.ExportDeclaredValue = 34
	p.Battery = "Y"

	pp.ProductList = append(pp.ProductList, p)

	results, err := client.RegisterProduct(pp)
	fmt.Println(err)
	if results != nil {
		fmt.Println(results.Code)
	}
}

func TestWinit_GetProductList(t *testing.T) {
	pp := &ProductListParam{}
	pp.PageNo = "1"
	pp.PageSize = "10"

	results, err := client.GetProductList(pp)
	fmt.Println(err)
	if results != nil {
		fmt.Println(results.Code)
	}
}