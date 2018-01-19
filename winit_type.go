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