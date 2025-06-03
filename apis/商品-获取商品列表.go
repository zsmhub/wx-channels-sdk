package apis

import (
	"encoding/json"
)

// 获取商品列表
// 文档：https://developers.weixin.qq.com/doc/channels/API/product/list_get.html

type ReqProductListGet struct {
	Status   int    `json:"status,omitempty"` // 商品状态，不填默认拉全部商品（不包含回收站）
	PageSize int    `json:"page_size"`        // 每页数量（默认10，不超过30），必填
	NextKey  string `json:"next_key"`         // 由上次请求返回，记录翻页的上下文。传入时会从上次返回的结果往后翻一页，不传默认拉取第一页数据。
}

var _ bodyer = ReqProductListGet{}

func (x ReqProductListGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespProductListGet struct {
	CommonResp
	ProductIds []int  `json:"product_ids"` // 商品 id 列表
	NextKey    string `json:"next_key"`    // 本次翻页的上下文，用于请求下一页
	TotalNum   int    `json:"total_num"`   // 商品总数
}

var _ bodyer = RespProductListGet{}

func (x RespProductListGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecProductListGet(req ReqProductListGet) (RespProductListGet, error) {
	var resp RespProductListGet
	err := c.executeWXApiPost("/channels/ec/product/list/get", req, &resp, true)
	if err != nil {
		return RespProductListGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespProductListGet{}, bizErr
	}
	return resp, nil
}
