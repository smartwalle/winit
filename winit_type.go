package winit

type winitParam struct {
	Action     string      `json:"action"`
	AppKey     string      `json:"app_key"`
	Data       interface{} `json:"data,omitempty"`
	Format     string      `json:"format"`
	Language   string      `json:"language"`
	Platform   string      `json:"platform"`
	Sign       string      `json:"sign"`
	SignMethod string      `json:"sign_method"`
	Timestamp  string      `json:"timestamp"`
	Version    string      `json:"version"`
}

type WinitParam interface {
	Action() string
}

type WinitResults struct {
	Code interface{} `json:"code"`
	Msg  string      `json:"msg"`
}

type PageParams struct {
	TotalCount int `json:"totalCount"`
	PageNo     int `json:"pageNo"`
	PageSize   int `json:"pageSize"`
}
