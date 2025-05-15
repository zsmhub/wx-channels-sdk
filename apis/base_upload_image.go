package apis

import (
	"encoding/json"
	"fmt"
)

// 上传图片，固定上传类型upload_type=1
// 文档：https://developers.weixin.qq.com/doc/channels/API/basics/img_upload.html

type ReqImgUpload struct {
	ImgUrl   string `json:"img_url"`   // upload_type=1时必填，图片url
	RespType int    `json:"resp_type"` // 返回数据类型。0:media_id和pay_media_id；1:图片链接（商品信息相关图片请务必使用此参数得到链接）
}

var _ bodyer = ReqImgUpload{}

func (x ReqImgUpload) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespImgUpload struct {
	CommonResp
	PicFile struct {
		MediaId    string `json:"media_id"`
		PayMediaId string `json:"pay_media_id"`
		ImgUrl     string `json:"img_url"`
	} `json:"pic_file"`
}

var _ bodyer = RespImgUpload{}

func (x RespImgUpload) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecImgUpload(req ReqImgUpload) (RespImgUpload, error) {
	var resp RespImgUpload
	reqUrl := fmt.Sprintf("/channels/ec/basics/img/upload?upload_type=1&resp_type=%d", req.RespType)
	err := c.executeWXApiPost(reqUrl, req, &resp, true)
	if err != nil {
		return RespImgUpload{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespImgUpload{}, bizErr
	}
	return resp, nil
}
