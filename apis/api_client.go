package apis

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"sync"
)

// API调用客户端
type ApiClient struct {
	AppId     string // 应用唯一身份标识
	AppSecret string // 应用密钥

	accessTokenName string // token参数名，一般都是填access_token
	accessToken     *token

	logger Logger
}

// API客户端初始化
func NewApiClient(appId, appSecret string, opts Options) *ApiClient {
	accessTokenName := "access_token"
	c := ApiClient{
		AppId:           appId,
		AppSecret:       appSecret,
		accessTokenName: accessTokenName,
		accessToken: &token{
			mutex:         &sync.RWMutex{},
			dcsToken:      opts.DcsToken,
			tokenCacheKey: fmt.Sprintf("%s#%s", accessTokenName, appId),
		},
		logger: opts.Logger,
	}

	if c.logger == nil {
		c.logger = loggerPrint{}
	}

	c.accessToken.setGetTokenFunc(c.getAccessToken)

	return &c
}

func (c *ApiClient) composeWXApiURL(path string, req interface{}) *url.URL {
	values := url.Values{}
	if valuer, ok := req.(urlValuer); ok {
		values = valuer.intoURLValues()
	}

	base, err := url.Parse(DefaultWXAPIHost)
	if err != nil {
		panic(fmt.Sprintf("qyapiHost invalid: host=%s err=%+v", DefaultWXAPIHost, err))
	}

	base.Path = path
	base.RawQuery = values.Encode()

	return base
}

func (c *ApiClient) composeWXURLWithToken(path string, req interface{}, withAccessToken bool) *url.URL {
	wxApiURL := c.composeWXApiURL(path, req)

	if !withAccessToken {
		return wxApiURL
	}

	q := wxApiURL.Query()
	q.Set(c.accessTokenName, c.accessToken.getToken())
	wxApiURL.RawQuery = q.Encode()

	return wxApiURL
}

func (c *ApiClient) executeWXApiGet(path string, req urlValuer, objResp interface{}, withAccessToken bool) error {
	wxUrlWithToken := c.composeWXURLWithToken(path, req, withAccessToken)
	urlStr := wxUrlWithToken.String()

	httpReq := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(httpReq)

	httpResp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(httpResp)

	httpReq.SetRequestURI(urlStr)
	httpReq.Header.SetMethod(http.MethodGet)

	if err := FastClient.DoTimeout(httpReq, httpResp, HttpTTL); err != nil {
		return err
	}

	respBody := httpResp.Body()
	if len(respBody) == 0 { // 避免 json.Unmarshal 出现 unexpected end of JSON input 错误
		c.logger.Errorf("请求视频号路由=%s, resp=%s, err=返回空响应体", urlStr, string(respBody))
		return errors.New("http resp body is nil")
	}

	return json.Unmarshal(respBody, &objResp)
}

func (c *ApiClient) executeWXApiPost(path string, req bodyer, objResp interface{}, withAccessToken bool) error {
	wxUrlWithToken := c.composeWXURLWithToken(path, req, withAccessToken)
	urlStr := wxUrlWithToken.String()

	reqBody, err := req.intoBody()
	if err != nil {
		return err
	}

	httpReq := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(httpReq)

	httpResp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(httpResp)

	httpReq.SetRequestURI(urlStr)
	httpReq.Header.SetContentType("application/json")
	httpReq.SetBody(reqBody)
	httpReq.Header.SetMethod(http.MethodPost)

	if err := FastClient.DoTimeout(httpReq, httpResp, HttpTTL); err != nil {
		return err
	}

	respBody := httpResp.Body()
	if len(respBody) == 0 { // 避免 json.Unmarshal 出现 unexpected end of JSON input 错误
		c.logger.Errorf("请求视频号路由=%s, req=%s, resp=%s, err=返回空响应体", path, string(reqBody), respBody)
		return errors.New("http resp body is nil")
	}

	return json.Unmarshal(respBody, &objResp)
}

func (c *ApiClient) executeWXApiGetReturnByte(path string, req urlValuer, withAccessToken bool) ([]byte, error) {
	wxUrlWithToken := c.composeWXURLWithToken(path, req, withAccessToken)
	urlStr := wxUrlWithToken.String()

	httpReq := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(httpReq)

	httpResp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(httpResp)

	httpReq.SetRequestURI(urlStr)
	httpReq.Header.SetMethod(http.MethodGet)

	if err := FastClient.DoTimeout(httpReq, httpResp, HttpTTL); err != nil {
		return nil, err
	}

	return httpResp.Body(), nil
}

func (c *ApiClient) executeWXApiMediaUpload(path string, req mediaUploader, objResp interface{}, withAccessToken bool) error {
	wxUrlWithToken := c.composeWXURLWithToken(path, req, withAccessToken)

	urlStr := wxUrlWithToken.String()

	m := req.getMedia()

	httpReq := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(httpReq)

	httpResp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(httpResp)

	// 新建一个缓冲，用于存放文件内容
	bodyBufer := &bytes.Buffer{}
	// 创建一个multipart文件写入器，方便按照http规定格式写入内容
	bodyWriter := multipart.NewWriter(bodyBufer)
	// 从bodyWriter生成fileWriter,并将文件内容写入fileWriter,多个文件可进行多次
	fileWriter, err := bodyWriter.CreateFormFile("media", m.filename)
	if err != nil {
		c.logger.Error(err.Error())
		return err
	}

	_, err = io.Copy(fileWriter, m.stream)
	if err != nil {
		return err
	}

	// 停止写入
	_ = bodyWriter.Close()

	httpReq.SetRequestURI(urlStr)
	httpReq.Header.SetContentType(bodyWriter.FormDataContentType())
	httpReq.SetBody(bodyBufer.Bytes())
	httpReq.Header.SetMethod(http.MethodPost)

	if err := FastClient.DoTimeout(httpReq, httpResp, HttpTTL); err != nil {
		return err
	}

	respBody := httpResp.Body()
	if len(respBody) == 0 { // 避免 json.Unmarshal 出现 unexpected end of JSON input 错误
		c.logger.Errorf("请求视频号路由=%s, resp=%s, err=返回空响应体", path, respBody)
		return errors.New("http resp body is nil")
	}

	return json.Unmarshal(respBody, &objResp)
}
