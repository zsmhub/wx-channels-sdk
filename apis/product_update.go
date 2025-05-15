package apis

import (
	"encoding/json"
)

// 更新商品
// 文档：https://developers.weixin.qq.com/doc/channels/API/product/update.html

type ReqProductUpdate struct {
	ProductID   string             `json:"product_id"`
	Attrs       []ProductAttrs     `json:"attrs"`
	Cats        []ProductCats      `json:"cats"`
	DescInfo    ProductDescInfo    `json:"desc_info"`
	ExpressInfo ProductExpressInfo `json:"express_info"`
	HeadImgs    []string           `json:"head_imgs"`
	Skus        []ProductSkus      `json:"skus"`
	SubTitle    string             `json:"sub_title"`
	Title       string             `json:"title"`
}

var _ bodyer = ReqProductUpdate{}

func (x ReqProductUpdate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespProductUpdate struct {
	Data struct {
		ProductID  string `json:"product_id"`
		UpdateTime string `json:"update_time"`
	} `json:"data"`
	CommonResp
}

var _ bodyer = RespProductUpdate{}

func (x RespProductUpdate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecProductUpdate(req ReqProductUpdate) (RespProductUpdate, error) {
	var resp RespProductUpdate
	err := c.executeWXApiPost("/channels/ec/product/update", req, &resp, true)
	if err != nil {
		return RespProductUpdate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespProductUpdate{}, bizErr
	}
	return resp, nil
}
