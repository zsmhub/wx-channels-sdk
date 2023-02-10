package apis

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
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

// 视频号小店API客户端初始化
func NewShopApiClient(appId, appSecret string, opts Options) *ApiClient {
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

	c.accessToken.setGetTokenFunc(c.getShopAccessToken)

	return &c
}

// 视频号橱窗API客户端初始化
func NewWindowApiClient(appId, appSecret string, opts Options) *ApiClient {
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

	c.accessToken.setGetTokenFunc(c.getWindowAccessToken)

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
