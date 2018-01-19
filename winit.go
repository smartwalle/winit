package winit

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	k_WINIT_SANDBOX_API_URL    = "http://erp.sandbox.winit.com.cn/ADInterface/api"
	k_WINIT_PRODUCTION_API_URL = "http://api.winit.com.cn/ADInterface/api"
)

const (
	k_WINIT_SANDBOX_OPEN_API_URL    = "http://openapi.sandbox.winit.com.cn/openapi/service"
	k_WINIT_PRODUCTION_OPEN_API_URL = "http://openapi.winit.com.cn/openapi/service"
)

const (
	k_WINIT_API_TYPE_AD       = 1
	k_WINIT_API_TYPE_OPEN_API = 2
)

var apiMap map[int]map[bool]string

func init() {
	apiMap = make(map[int]map[bool]string)
	apiMap[k_WINIT_API_TYPE_AD] = make(map[bool]string)
	apiMap[k_WINIT_API_TYPE_AD][true] = k_WINIT_PRODUCTION_API_URL
	apiMap[k_WINIT_API_TYPE_AD][false] = k_WINIT_SANDBOX_API_URL

	apiMap[k_WINIT_API_TYPE_OPEN_API] = make(map[bool]string)
	apiMap[k_WINIT_API_TYPE_OPEN_API][true] = k_WINIT_PRODUCTION_OPEN_API_URL
	apiMap[k_WINIT_API_TYPE_OPEN_API][false] = k_WINIT_SANDBOX_OPEN_API_URL
}

type Winit struct {
	appKey       string
	token        string
	isProduction bool
	Client       *http.Client
}

func New(appKey, token string, isProduction bool) (client *Winit) {
	client = &Winit{}
	client.appKey = appKey
	client.token = token
	client.Client = http.DefaultClient
	client.isProduction = isProduction
	return client
}

func (this *Winit) doRequest(apiType int, method string, param WinitParam, results interface{}) (err error) {
	var buf io.Reader
	if param == nil {
		return errors.New("请求参数不能为空")
	}

	var p = &winitParam{}
	p.Action = param.Action()
	p.AppKey = this.appKey
	p.Timestamp = ""
	p.Version = "1.0"
	p.SignMethod = "md5"
	p.Format = "json"
	p.Platform = "SELLERERP"
	p.Language = "zh_CN"
	p.Data = param

	var now = time.Now()
	p.Timestamp = now.Format("2006-01-02 15:04:05")

	sign, err := signMD5(this.appKey, this.token, p)
	if err != nil {
		return err
	}
	p.Sign = sign

	var pBytes, _ = json.Marshal(p)
	buf = bytes.NewReader(pBytes)

	fmt.Println("--------------------------------------------------")
	fmt.Println("param", string(pBytes))
	fmt.Println("--------------------------------------------------")

	var apiDomain = apiMap[apiType][this.isProduction]
	req, err := http.NewRequest(method, apiDomain, buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	resp, err := this.Client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(resp.Status)

	fmt.Println("--------------------------------------------------")
	fmt.Println("results", string(data))
	fmt.Println("--------------------------------------------------")

	err = json.Unmarshal(data, results)
	if err != nil {
		return err
	}

	return err
}

func signMD5(appKey, token string, p *winitParam) (s string, err error) {
	if p == nil {
		return "", errors.New("签名参数不能为空")
	}
	var sBuffer = &bytes.Buffer{}
	sBuffer.WriteString(token)
	sBuffer.WriteString("action" + p.Action)
	sBuffer.WriteString("app_key" + appKey)
	var dataBytes, _ = json.Marshal(p.Data)
	sBuffer.WriteString("data" + string(dataBytes))
	sBuffer.WriteString("format" + p.Format)
	sBuffer.WriteString("platform" + p.Platform)
	sBuffer.WriteString("sign_method" + p.SignMethod)
	sBuffer.WriteString("timestamp" + p.Timestamp)
	sBuffer.WriteString("version" + p.Version)
	sBuffer.WriteString(token)

	var m = md5.New()
	m.Write(sBuffer.Bytes())
	s = hex.EncodeToString(m.Sum(nil))
	s = strings.ToUpper(s)
	return s, nil
}
