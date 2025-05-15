package apis

import (
	"encoding/json"
)

// 电子面单取号
// 文档：https://developers.weixin.qq.com/doc/channels/API/ewaybill/create_order.html

type ReqLogisticsEwaybillBizOrderCreate struct {
	DeliveryID  string `json:"delivery_id"`
	EcOrderList []struct {
		EcOrderID int `json:"ec_order_id"`
		GoodsList []struct {
			GoodCount int    `json:"good_count"`
			GoodName  string `json:"good_name"`
			ProductID int    `json:"product_id"`
			SkuID     int    `json:"sku_id"`
		} `json:"goods_list"`
	} `json:"ec_order_list"`
	EwaybillAcctID  string `json:"ewaybill_acct_id"`
	EwaybillOrderID string `json:"ewaybill_order_id"`
	Receiver        struct {
		Address  string `json:"address"`
		City     string `json:"city"`
		County   string `json:"county"`
		Mobile   string `json:"mobile"`
		Name     string `json:"name"`
		Province string `json:"province"`
		Street   string `json:"street"`
	} `json:"receiver"`
	Remark        string `json:"remark"`
	ReturnAddress struct {
		Address  string `json:"address"`
		City     string `json:"city"`
		County   string `json:"county"`
		Mobile   string `json:"mobile"`
		Name     string `json:"name"`
		Province string `json:"province"`
		Street   string `json:"street"`
	} `json:"return_address"`
	Sender struct {
		Address  string `json:"address"`
		City     string `json:"city"`
		County   string `json:"county"`
		Mobile   string `json:"mobile"`
		Name     string `json:"name"`
		Province string `json:"province"`
		Street   string `json:"street"`
	} `json:"sender"`
	ShopID     string `json:"shop_id"`
	SiteCode   string `json:"site_code"`
	TemplateID string `json:"template_id"`
}

var _ bodyer = ReqLogisticsEwaybillBizOrderCreate{}

func (x ReqLogisticsEwaybillBizOrderCreate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespLogisticsEwaybillBizOrderCreate struct {
	DeliveryErrorMsg string `json:"delivery_error_msg"`
	CommonResp
	EwaybillOrderID string `json:"ewaybill_order_id"`
	WaybillID       string `json:"waybill_id"`
}

var _ bodyer = RespLogisticsEwaybillBizOrderCreate{}

func (x RespLogisticsEwaybillBizOrderCreate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecLogisticsEwaybillBizOrderCreate(req ReqLogisticsEwaybillBizOrderCreate) (RespLogisticsEwaybillBizOrderCreate, error) {
	var resp RespLogisticsEwaybillBizOrderCreate
	err := c.executeWXApiPost("/channels/ec/logistics/ewaybill/biz/order/create", req, &resp, true)
	if err != nil {
		return RespLogisticsEwaybillBizOrderCreate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespLogisticsEwaybillBizOrderCreate{}, bizErr
	}
	return resp, nil
}
