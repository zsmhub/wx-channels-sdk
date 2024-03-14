package apis

import (
	"encoding/json"
)

// 查询面单详情
// 文档：https://developers.weixin.qq.com/doc/channels/API/ewaybill/get_order.html

type ReqLogisticsEwaybillBizOrderGet struct {
	EwaybillOrderID string `json:"ewaybill_order_id"`
}

var _ bodyer = ReqLogisticsEwaybillBizOrderGet{}

func (x ReqLogisticsEwaybillBizOrderGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespLogisticsEwaybillBizOrderGet struct {
	CommonResp
	OrderInfo struct {
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
		PathInfo        []struct {
			Desc       string `json:"desc"`
			Status     int    `json:"status"`
			UpdateTime int    `json:"update_time"`
		} `json:"path_info"`
		Receiver struct {
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
		ShopID    string `json:"shop_id"`
		SiteCode  string `json:"site_code"`
		Status    int    `json:"status"`
		WaybillID string `json:"waybill_id"`
	} `json:"order_info"`
}

var _ bodyer = RespLogisticsEwaybillBizOrderGet{}

func (x RespLogisticsEwaybillBizOrderGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecLogisticsEwaybillBizOrderGet(req ReqLogisticsEwaybillBizOrderGet) (RespLogisticsEwaybillBizOrderGet, error) {
	var resp RespLogisticsEwaybillBizOrderGet
	err := c.executeWXApiPost("/channels/ec/logistics/ewaybill/biz/order/get", req, &resp, true)
	if err != nil {
		return RespLogisticsEwaybillBizOrderGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespLogisticsEwaybillBizOrderGet{}, bizErr
	}
	return resp, nil
}
