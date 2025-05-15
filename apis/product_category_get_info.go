package apis

import (
	"encoding/json"
)

// 获取类目信息
// 文档：https://developers.weixin.qq.com/doc/channels/API/category/getcategorydetail.html

type ReqCategoryDetail struct {
	CatID int `json:"cat_id"`
}

var _ bodyer = ReqCategoryDetail{}

func (x ReqCategoryDetail) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespCategoryDetail struct {
	Attr struct {
		AccessPermitRequired bool `json:"access_permit_required"`
		BrandList            []struct {
			BrandID string `json:"brand_id"`
		} `json:"brand_list"`
		Deposit         string `json:"deposit"`
		PreSale         bool   `json:"pre_sale"`
		ProductAttrList []struct {
			IsRequired bool   `json:"is_required"`
			Name       string `json:"name"`
			Type       string `json:"type"`
			Value      string `json:"value"`
		} `json:"product_attr_list"`
		ShopNoShipment     bool `json:"shop_no_shipment"`
		TransactionfeeInfo struct {
			BasisPoint         int `json:"basis_point"`
			IncentiveType      int `json:"incentive_type"`
			OriginalBasisPoint int `json:"original_basis_point"`
		} `json:"transactionfee_info"`
	} `json:"attr"`
	CommonResp
	Info struct {
		CatID string `json:"cat_id"`
		Name  string `json:"name"`
	} `json:"info"`
}

var _ bodyer = RespCategoryDetail{}

func (x RespCategoryDetail) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecCategoryDetail(req ReqCategoryDetail) (RespCategoryDetail, error) {
	var resp RespCategoryDetail
	err := c.executeWXApiPost("/channels/ec/category/detail", req, &resp, true)
	if err != nil {
		return RespCategoryDetail{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCategoryDetail{}, bizErr
	}
	return resp, nil
}
