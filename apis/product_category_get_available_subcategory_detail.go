package apis

import (
	"encoding/json"
)

// 获取可用的子类目详情
// 文档：https://developers.weixin.qq.com/doc/channels/API/category/getavailablesoncategories.html

type ReqCategoryAvailablesoncategoriesGet struct {
	FCatID int `json:"f_cat_id"`
}

var _ bodyer = ReqCategoryAvailablesoncategoriesGet{}

func (x ReqCategoryAvailablesoncategoriesGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

type RespCategoryAvailablesoncategoriesGet struct {
	CatList []struct {
		CatID  int    `json:"cat_id"`
		FCatID int    `json:"f_cat_id"`
		Name   string `json:"name"`
	} `json:"cat_list"`
	CommonResp
}

var _ bodyer = RespCategoryAvailablesoncategoriesGet{}

func (x RespCategoryAvailablesoncategoriesGet) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecCategoryAvailablesoncategoriesGet(req ReqCategoryAvailablesoncategoriesGet) (RespCategoryAvailablesoncategoriesGet, error) {
	var resp RespCategoryAvailablesoncategoriesGet
	err := c.executeWXApiPost("/channels/ec/category/availablesoncategories/get", req, &resp, true)
	if err != nil {
		return RespCategoryAvailablesoncategoriesGet{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespCategoryAvailablesoncategoriesGet{}, bizErr
	}
	return resp, nil
}
