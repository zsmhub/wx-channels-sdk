package apis

import (
	"encoding/json"
	
)

// 电子面单预取号
// 文档：https://developers.weixin.qq.com/doc/channels/API/ewaybill/precreate_order.html

type ReqLogisticsEwaybillBizOrderPrecreate struct {
	DeliveryID  string `json:"delivery_id"`
	EcOrderList []struct {
		EcOrderID int `json:"ec_order_id"`
		GoodsList []struct {
			GoodCount int  `json:"good_count"`
			GoodName  string `json:"good_name"`
			ProductID int  `json:"product_id"`
			SkuID     int  `json:"sku_id"`
		} `json:"goods_list"`
	} `json:"ec_order_list"`
	EwaybillAcctID string `json:"ewaybill_acct_id"`
	Receiver       struct {
		Address  string `json:"address"`
		City     string `json:"city"`
		County   string `json:"county"`
		Mobile   string `json:"mobile"`
		Name     string `json:"name"`
		Province string `json:"province"`
		Street   string `json:"street"`
	} `json:"receiver"`
	Remark string `json:"remark"`
	Sender struct {
		Address  string `json:"address"`
		City     string `json:"city"`
		County   string `json:"county"`
		Mobile   string `json:"mobile"`
		Name     string `json:"name"`
		Province string `json:"province"`
		Street   string `json:"street"`
	} `json:"sender"`
	ShopID   string `json:"shop_id"`
	SiteCode string `json:"site_code"`
}



var _ bodyer = ReqLogisticsEwaybillBizOrderPrecreate{}

func (x ReqLogisticsEwaybillBizOrderPrecreate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}


type RespLogisticsEwaybillBizOrderPrecreate struct {
	CommonResp
	EwaybillOrderID string `json:"ewaybill_order_id"`
}

var _ bodyer = RespLogisticsEwaybillBizOrderPrecreate{}

func (x RespLogisticsEwaybillBizOrderPrecreate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecLogisticsEwaybillBizOrderPrecreate(req ReqLogisticsEwaybillBizOrderPrecreate) (RespLogisticsEwaybillBizOrderPrecreate, error) {
	var resp RespLogisticsEwaybillBizOrderPrecreate
	err := c.executeWXApiPost("/channels/ec/logistics/ewaybill/biz/order/precreate", req, &resp, true)
	if err != nil {
		return RespLogisticsEwaybillBizOrderPrecreate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespLogisticsEwaybillBizOrderPrecreate{}, bizErr
	}
	return resp, nil
}