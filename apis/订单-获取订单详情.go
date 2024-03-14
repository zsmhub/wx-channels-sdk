package apis

import (
	"encoding/json"
)

// 获取订单详情
// 文档：https://developers.weixin.qq.com/doc/channels/API/order/get.html

type ReqOrderGet struct {
	OrderID string `json:"order_id"`
}

var _ bodyer = ReqOrderGet{}

func (x ReqOrderGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespOrderGet struct {
	CommonResp
	Order struct {
		AftersaleDetail struct {
			AftersaleOrderList []struct {
				AftersaleOrderID string `json:"aftersale_order_id"`
				Status           int    `json:"status"`
			} `json:"aftersale_order_list"`
			OnAftersaleOrderCnt int `json:"on_aftersale_order_cnt"`
		} `json:"aftersale_detail"`
		CreateTime  int    `json:"create_time"`
		Openid      string `json:"openid"`
		Unionid     string `json:"unionid"`
		OrderDetail struct {
			ProductInfos []struct {
				ProductID             string `json:"product_id"`
				SkuID                 string `json:"sku_id"`
				ThumbImg              string `json:"thumb_img"`
				SkuCnt                int    `json:"sku_cnt"`
				SalePrice             int    `json:"sale_price"`
				Title                 string `json:"title"`
				OnAftersaleSkuCnt     int    `json:"on_aftersale_sku_cnt"`
				FinishAftersaleSkuCnt int    `json:"finish_aftersale_sku_cnt"`
				SkuCode               string `json:"sku_code"`
				MarketPrice           int    `json:"market_price"`
				SkuAttrs              []struct {
					AttrKey   string `json:"attr_key"`
					AttrValue string `json:"attr_value"`
				} `json:"sku_attrs"`
				RealPrice      int    `json:"real_price"`
				OutProductId   string `json:"out_product_id"`
				OutSkuId       string `json:"out_sku_id"`
				IsDiscounted   bool   `json:"is_discounted"`
				EstimatePrice  int    `json:"estimate_price"`
				IsChangePrice  bool   `json:"is_change_price"`
				ChangePrice    int    `json:"change_price"`
				OutWarehouseId string `json:"out_warehouse_id"`
			} `json:"product_infos"`
			PriceInfo struct {
				ProductPrice         int  `json:"product_price"`
				OrderPrice           int  `json:"order_price"`
				Freight              int  `json:"freight"`
				DiscountedPrice      int  `json:"discounted_price"`
				IsDiscounted         bool `json:"is_discounted"`
				OriginalOrderPrice   int  `json:"original_order_price"`
				EstimateProductPrice int  `json:"estimate_product_price"`
				ChangeDownPrice      int  `json:"change_down_price"`
				ChangeFreight        int  `json:"change_freight"`
				IsChangeFreight      bool `json:"is_change_freight"`
			} `json:"price_info"`
			PayInfo struct {
				PayTime       int    `json:"pay_time"`
				PrepayID      string `json:"prepay_id"`
				PrepayTime    int    `json:"prepay_time"`
				TransactionID string `json:"transaction_id"`
			} `json:"pay_info"`
			DeliveryInfo struct {
				AddressInfo         OrderDetailAddress `json:"address_info"`
				DeliverMethod       int                `json:"deliver_method"`
				DeliveryProductInfo []struct {
					WaybillID    string `json:"waybill_id"`
					DeliveryID   string `json:"delivery_id"`
					ProductInfos []struct {
						ProductCnt int    `json:"product_cnt"`
						ProductID  string `json:"product_id"`
						SkuID      string `json:"sku_id"`
					} `json:"product_infos"`
					DeliveryName    string             `json:"delivery_name"`
					DeliveryTime    int                `json:"delivery_time"`
					DeliverType     int                `json:"deliver_type"`
					DeliveryAddress OrderDetailAddress `json:"delivery_address"`
				} `json:"delivery_product_info"`
				ShipDoneTime int `json:"ship_done_time"`
			} `json:"delivery_info"`
			CouponInfo struct {
				UserCouponID string `json:"user_coupon_id"`
			} `json:"coupon_info"`
			ExtInfo struct {
				CustomerNotes string `json:"customer_notes"`
				MerchantNotes string `json:"merchant_notes"`
			} `json:"ext_info"`
			CommissionInfos []struct {
				SkuID    int    `json:"sku_id"`
				Nickname string `json:"nickname"`
				Type     int    `json:"type"`
				Status   int    `json:"status"`
				Amount   int    `json:"amount"`
				FinderId string `json:"finder_id"`
			} `json:"commission_infos"`
			SharerInfo struct {
				SharerOpenid  string `json:"sharer_openid"`
				SharerUnionid string `json:"sharer_unionid"`
				SharerType    int    `json:"sharer_type"`
				ShareScene    int    `json:"share_scene"`
			} `json:"sharer_info"`
			SettleInfo struct {
				PredictCommissionFee int `json:"predict_commission_fee"`
				CommissionFee        int `json:"commission_fee"`
			} `json:"settle_info"`
		} `json:"order_detail"`
		OrderID    string `json:"order_id"`
		Status     int    `json:"status"`
		UpdateTime int    `json:"update_time"`
	} `json:"order"`
}

type OrderDetailAddress struct {
	UserName              string            `json:"user_name"`
	PostalCode            string            `json:"postal_code"`
	ProvinceName          string            `json:"province_name"`
	CityName              string            `json:"city_name"`
	CountyName            string            `json:"county_name"`
	DetailInfo            string            `json:"detail_info"`
	NationalCode          string            `json:"national_code"`
	TelNumber             string            `json:"tel_number"`
	HouseNumber           string            `json:"house_number"`
	VirtualOrderTelNumber string            `json:"virtual_order_tel_number"`
	TelNumberExtInfo      *TelNumberExtInfo `json:"tel_number_ext_info"`
	UseTelNumber          int               `json:"use_tel_number"`
	HashCode              string            `json:"hash_code"`
}

type TelNumberExtInfo struct {
	RealTelNumber        string `json:"real_tel_number"`
	VirtualTelNumber     string `json:"virtual_tel_number"`
	VirtualTelExpireTime int    `json:"virtual_tel_expire_time"`
	GetVirtualTelCnt     int    `json:"get_virtual_tel_cnt"`
}

var _ bodyer = RespOrderGet{}

func (x RespOrderGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecOrderGet(req ReqOrderGet) (RespOrderGet, error) {
	var resp RespOrderGet
	err := c.executeWXApiPost("/channels/ec/order/get", req, &resp, true)
	if err != nil {
		return RespOrderGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespOrderGet{}, bizErr
	}
	return resp, nil
}
