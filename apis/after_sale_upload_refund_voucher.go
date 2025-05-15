package apis

import (
	"encoding/json"
)

// 上传退款凭证
// 文档：https://developers.weixin.qq.com/doc/channels/API/aftersale/uploadrefundcertificate.html

type ReqAftersaleUploadrefundcertificate struct {
	AfterSaleOrderID   string   `json:"after_sale_order_id"`
	RefundCertificates []string `json:"refund_certificates"`
	Desc               string   `json:"desc"`
}

var _ bodyer = ReqAftersaleUploadrefundcertificate{}

func (x ReqAftersaleUploadrefundcertificate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespAftersaleUploadrefundcertificate struct {
	CommonResp
}

var _ bodyer = RespAftersaleUploadrefundcertificate{}

func (x RespAftersaleUploadrefundcertificate) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecAftersaleUploadrefundcertificate(req ReqAftersaleUploadrefundcertificate) (RespAftersaleUploadrefundcertificate, error) {
	var resp RespAftersaleUploadrefundcertificate
	err := c.executeWXApiPost("/channels/ec/aftersale/uploadrefundcertificate", req, &resp, true)
	if err != nil {
		return RespAftersaleUploadrefundcertificate{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespAftersaleUploadrefundcertificate{}, bizErr
	}
	return resp, nil
}
