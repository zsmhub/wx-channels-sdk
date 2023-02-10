package apis

import (
	"encoding/json"
)

// 更新商品
// 文档：https://developers.weixin.qq.com/doc/channels/API/product/update.html

type ReqProductUpdate struct {
	Attrs []struct {
		AttrKey   string `json:"attr_key"`
		AttrValue string `json:"attr_value"`
	} `json:"attrs"`
	Cats []struct {
		CatID string `json:"cat_id"`
	} `json:"cats"`
	DescInfo struct {
		Desc string   `json:"desc"`
		Imgs []string `json:"imgs"`
	} `json:"desc_info"`
	ExpressInfo struct {
		TemplateID string `json:"template_id"`
	} `json:"express_info"`
	HeadImgs  []string `json:"head_imgs"`
	ProductID string   `json:"product_id"`
	Skus      []struct {
		SalePrice int `json:"sale_price"`
		SkuAttrs  []struct {
			AttrKey   string `json:"attr_key"`
			AttrValue string `json:"attr_value"`
		} `json:"sku_attrs"`
		SkuCode  string `json:"sku_code"`
		SkuID    string `json:"sku_id"`
		StockNum int    `json:"stock_num"`
		ThumbImg string `json:"thumb_img"`
	} `json:"skus"`
	SubTitle string `json:"sub_title"`
	Title    string `json:"title"`
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
