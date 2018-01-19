package winit

// RegisterProductParam http://developer.winit.com.cn/Index/index.php?s=Index/index/id/19/l/zh-cn
type RegisterProductParam struct {
	ProductList []*Product `json:"productList"`
}

func (this *RegisterProductParam) Action() string {
	return "registerProduct"
}

type Product struct {
	Battery             string  `json:"battery"`
	Branded             string  `json:"branded"`
	BrandedName         string  `json:"brandedName"`
	CategoryOne         string  `json:"categoryOne"`
	CategoryThree       string  `json:"categoryThree"`
	CategoryTwo         string  `json:"categoryTwo"`
	ChineseName         string  `json:"chineseName"`
	DisplayPageUrl      string  `json:"displayPageUrl"`
	EnglishName         string  `json:"englishName"`
	ExportCountry       string  `json:"exportCountry"`
	ExportDeclaredValue float64 `json:"exportDeclaredvalue"`
	FixedVolumeWeight   string  `json:"fixedVolumeWeight"`
	InporCountry        string  `json:"inporCountry"`
	InportDeclaredValue float64 `json:"inportDeclaredvalue"`
	Model               string  `json:"model"`
	ProductCode         string  `json:"productCode"`
	RegisteredHeight    float32 `json:"registeredHeight"`
	RegisteredLength    float32 `json:"registeredLength"`
	RegisteredWeight    float32 `json:"registeredWeight"`
	RegisteredWidth     float32 `json:"registeredWidth"`
	Remark              string  `json:"remark"`
	Specification       string  `json:"specification"`
}

type RegisterProductResults struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data []*RegisterProductData `json:"data"`
}

type RegisterProductData struct {
	ProductCode string `json:"productCode"`
}

// ProductListParam https://developer.winit.com.cn/document/detail/id/17.html
type ProductListParam struct {
	PageNo   string `json:"pageNo"`
	PageSize string `json:"pageSize"`
	SKUCode  string `json:"skuCode"`
}

func (this *ProductListParam) Action() string {
	return "winit.mms.item.list"
}

type ProductListResults struct {
	Code string           `json:"code"`
	Msg  string           `json:"msg"`
	Data *ProductListData `json:"data"`
}

type ProductListData struct {
	List       []*ProductList `json:"list"`
	PageParams *PageParams    `json:"PageParams"`
}

type ProductList struct {
	Id                     string                `json:"id"`
	Code                   string                `json:"code"`
	SKUCode                string                `json:"skuCode"`
	Specification          string                `json:"specification"`
	CNName                 string                `json:"cnName"`
	Name                   string                `json:"name"`
	Length                 float32               `json:"length"`
	Width                  float32               `json:"width"`
	Height                 float32               `json:"height"`
	Volume                 float32               `json:"volume"`
	Weight                 float32               `json:"weight"`
	RegisterWeight         float32               `json:"registerWeight"`
	RegisterLength         float32               `json:"registerLength"`
	RegisterWidth          float32               `json:"registerWidth"`
	RegisterHeight         float32               `json:"registerHeight"`
	RegisterVolume         float32               `json:"registerVolume"`
	CustomsDeclarationList []*CustomsDeclaration `json:"customsDeclarationList"`
}

type CustomsDeclaration struct {
	CountryCode string  `json:"countryCode"`
	DeclareName string  `json:"declareName"`
	ImportPrice float32 `json:"importPrice"`
	ExportPrice float32 `json:"exportPrice"`
	RebateRate  float32 `json:"rebateRate"`
	ImportRate  float32 `json:"importRate"`
	VatRate     float32 `json:"vatRate"`
}
