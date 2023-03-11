package apis

import (
	"encoding/json"
)

// 添加商品
// 文档：https://developers.weixin.qq.com/doc/channels/API/product/add.html

type ReqProductAdd struct {
	Attrs       []ProductAttrs     `json:"attrs"`
	Cats        []ProductCats      `json:"cats"`
	DescInfo    ProductDescInfo    `json:"desc_info"`
	ExpressInfo ProductExpressInfo `json:"express_info"`
	HeadImgs    []string           `json:"head_imgs"`
	Skus        []ProductSkus      `json:"skus"`
	SubTitle    string             `json:"sub_title"`
	Title       string             `json:"title"`
}

type ProductAttrs struct {
	AttrKey   string `json:"attr_key"`
	AttrValue string `json:"attr_value"`
}

type ProductCats struct {
	CatID string `json:"cat_id"`
}

type ProductDescInfo struct {
	Desc string   `json:"desc"`
	Imgs []string `json:"imgs"`
}

type ProductExpressInfo struct {
	TemplateID string `json:"template_id"`
}

type ProductSkus struct {
	SalePrice int               `json:"sale_price"`
	SkuAttrs  []ProductSkuAttrs `json:"sku_attrs"`
	SkuCode   string            `json:"sku_code"`
	StockNum  int               `json:"stock_num"`
	ThumbImg  string            `json:"thumb_img"`
}

type ProductSkuAttrs struct {
	AttrKey   string `json:"attr_key"`
	AttrValue string `json:"attr_value"`
}

var _ bodyer = ReqProductAdd{}

func (x ReqProductAdd) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespProductAdd struct {
	Data struct {
		CreateTime string `json:"create_time"`
		ProductID  int    `json:"product_id"`
	} `json:"data"`
	CommonResp
}

var _ bodyer = RespProductAdd{}

func (x RespProductAdd) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecProductAdd(req ReqProductAdd) (RespProductAdd, error) {
	var resp RespProductAdd
	err := c.executeWXApiPost("/channels/ec/product/add", req, &resp, true)
	if err != nil {
		return RespProductAdd{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespProductAdd{}, bizErr
	}
	return resp, nil
}
