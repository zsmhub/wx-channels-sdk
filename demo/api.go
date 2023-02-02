package demo

import (
	"fmt"
	"github.com/zsmhub/wx-channels-sdk"
	"github.com/zsmhub/wx-channels-sdk/apis"
)

// 调用 视频号API 示例
func ApiMain() {
	if err := initApiHandler(); err != nil {
		fmt.Println(err)
	}

	// 获取 access_token
	resp, err := channels.Sdk.ShopClient.GetToken()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("access_token: ", resp)
	}

	// 获取店铺信息
	resp2, err2 := channels.Sdk.ShopClient.ExecEcBasicsInfoGet(apis.ReqEcBasicsInfoGet{})
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Printf("%+v\n", resp2)
	}

	// 获取商品列表
	resp3, err3 := channels.Sdk.ShopClient.ExecEcProductListGet(apis.ReqEcProductListGet{
		Status:   0,
		PageSize: 10,
		NextKey:  "",
	})
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Printf("%+v\n", resp3)
	}
}

// API客户端初始化
func initApiHandler() error {
	// 初始化sdk参数
	channels.Sdk.InitOptions(apis.Options{
		DcsToken: DcsTokenByRedis{},
		Logger:   Logger{},
	})

	// API客户端初始化
	channels.Sdk.NewShopApiClient(ShopAppId, ShopAppSecret)

	return nil
}
