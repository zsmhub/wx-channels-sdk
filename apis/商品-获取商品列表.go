package apis

import (
	"encoding/json"
)

// 获取商品列表
// 文档：https://developers.weixin.qq.com/doc/channels/API/product/list_get.html

type ReqEcProductListGet struct {
	// 商品状态，不填默认拉全部商品（不包含回收站）
	Status int `json:"status"`
	// 每页数量（默认10，不超过30），必填
	PageSize int `json:"page_size"`
	// 由上次请求返回，记录翻页的上下文。传入时会从上次返回的结果往后翻一页，不传默认拉取第一页数据。
	NextKey string `json:"next_key"`
}

var _ bodyer = ReqEcProductListGet{}

func (x ReqEcProductListGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespEcProductListGet struct {
	CommonResp
	// 商品 id 列表
	ProductIds []int `json:"product_ids"`
	// 本次翻页的上下文，用于请求下一页
	NextKey string `json:"next_key"`
	// 商品总数
	TotalNum int `json:"total_num"`
}

var _ bodyer = RespEcProductListGet{}

func (x RespEcProductListGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecEcProductListGet(req ReqEcProductListGet) (RespEcProductListGet, error) {
	var resp RespEcProductListGet
	err := c.executeWXApiPost("/channels/ec/product/list/get", req, &resp, true)
	if err != nil {
		return RespEcProductListGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespEcProductListGet{}, bizErr
	}
	return resp, nil
}
