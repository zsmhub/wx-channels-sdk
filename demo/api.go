package demo

import (
	"fmt"
	channels "github.com/zsmhub/wx-channels-sdk"
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
	resp3, err3 := channels.Sdk.ShopClient.ExecProductListGet(apis.ReqProductListGet{
		Status:   0,
		PageSize: 10,
		NextKey:  "",
	})
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Printf("%+v\n", resp3)
	}

	// 上传图片
	//resp4, err4 := channels.Sdk.ShopClient.ExecImgUpload(apis.ReqImgUpload{
	//	ImgUrl:   "https://img2.baidu.com/it/u=4106804942,1016065650&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=474",
	//	RespType: 1,
	//})
	//if err4 != nil {
	//	fmt.Println(err4)
	//} else {
	//	fmt.Printf("%+v\n", resp4)
	//}

	// 上传资质图片
	//resp5, err5 := channels.Sdk.ShopClient.ExecQualificationUpload(apis.ReqQualificationUpload{
	//	URL: "https://img2.baidu.com/it/u=4106804942,1016065650&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=474",
	//})
	//if err5 != nil {
	//	fmt.Println(err5)
	//} else {
	//	fmt.Printf("%+v\n", resp5)
	//}
}

// API客户端初始化
func initApiHandler() error {
	// 初始化sdk参数
	channels.Sdk.InitOptions(apis.Options{
		DcsToken: DcsTokenByRedis{},
		Logger:   Logger{},
	})

	// 视频号小店API客户端初始化
	channels.Sdk.NewShopApiClient(ShopAppId, ShopAppSecret)

	// 视频号橱窗API客户端初始化
	//channels.Sdk.NewWindowApiClient(WindowAppId, WindowAppSecret)

	return nil
}
