package apis

import (
	"encoding/json"
)

// 获取商品
// 文档：https://developers.weixin.qq.com/doc/channels/API/product/get.html

type ReqProductGet struct {
	DataType  int    `json:"data_type"`
	ProductID string `json:"product_id"`
}

var _ bodyer = ReqProductGet{}

func (x ReqProductGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespProductGet struct {
	CommonResp
	Product struct {
		Attrs []struct {
			AttrKey   string `json:"attr_key"`
			AttrValue string `json:"attr_value"`
		} `json:"attrs"`
		Cats []struct {
			CatID string `json:"cat_id"`
		} `json:"cats"`
		DescInfo struct {
			Imgs []string `json:"imgs"`
		} `json:"desc_info"`
		EditStatus  int `json:"edit_status"`
		ExpressInfo struct {
			TemplateID string `json:"template_id"`
		} `json:"express_info"`
		HeadImgs     []string `json:"head_imgs"`
		MinPrice     int      `json:"min_price"`
		OutProductID string   `json:"out_product_id"`
		ProductID    string   `json:"product_id"`
		Skus         []struct {
			OutSkuID  string        `json:"out_sku_id"`
			SalePrice int           `json:"sale_price"`
			SkuAttrs  []interface{} `json:"sku_attrs"`
			SkuCode   string        `json:"sku_code"`
			SkuID     string        `json:"sku_id"`
			Status    int           `json:"status"`
			StockNum  int           `json:"stock_num"`
			ThumbImg  string        `json:"thumb_img"`
		} `json:"skus"`
		SpuCode  string `json:"spu_code"`
		Status   int    `json:"status"`
		SubTitle string `json:"sub_title"`
		Title    string `json:"title"`
	} `json:"product"`
}

var _ bodyer = RespProductGet{}

func (x RespProductGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecProductGet(req ReqProductGet) (RespProductGet, error) {
	var resp RespProductGet
	err := c.executeWXApiPost("/channels/ec/product/get", req, &resp, true)
	if err != nil {
		return RespProductGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespProductGet{}, bizErr
	}
	return resp, nil
}
