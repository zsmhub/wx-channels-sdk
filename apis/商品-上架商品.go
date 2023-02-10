package apis

import (
	"encoding/json"
)

// 上架商品
// 文档：https://developers.weixin.qq.com/doc/channels/API/product/listing.html

type ReqProductListing struct {
	ProductID string `json:"product_id"`
}

var _ bodyer = ReqProductListing{}

func (x ReqProductListing) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespProductListing struct {
	CommonResp
}

var _ bodyer = RespProductListing{}

func (x RespProductListing) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecProductListing(req ReqProductListing) (RespProductListing, error) {
	var resp RespProductListing
	err := c.executeWXApiPost("/channels/ec/product/listing", req, &resp, true)
	if err != nil {
		return RespProductListing{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespProductListing{}, bizErr
	}
	return resp, nil
}
