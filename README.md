## 微信视频号 GO SDK

Go语言实现微信视频号sdk，集成了视频号小店和视频号橱窗的功能，使用简单，扩展灵活

- 支持一键生成sdk代码，包括api和回调事件
- 用缓存方案实现分布式 access_token，保证在多个服务中只有一个服务能成功调用API请求 access_token，减少API调用次数和服务重启需要重新获取的情况
    + 缓存方案支持自定义存储，默认内存存储
- 支持自定义日志存储，提供Logger interface：用于自行实现日志记录器，便于收集日志
    + 默认 log.Printf 输出

[点击查看视频号的开发博文](https://zsmhub.github.io/post/%E5%AE%9E%E6%88%98%E6%A1%88%E4%BE%8B/%E5%BE%AE%E4%BF%A1%E8%A7%86%E9%A2%91%E5%8F%B7/)

### 安装命令

```sh
go get github.com/zsmhub/wx-channels-sdk
```

### 一键生成sdk代码命令(`需手动格式化代码`)

todo

### sdk调用示例

**强烈建议去 ./demo 文件夹查看完整示例！**

[点击查看完整demo](https://github.com/zsmhub/wx-channels-sdk/tree/main/demo)

todo

### 目录结构

todo

### 注意点

- 如果你发现了sdk中，没有某个回调事件或某个api，可以使用一键生成sdk代码命令生成，然后提交下pr

### 兄弟项目

- [企业微信 GO SDK](https://github.com/zsmhub/workweixin)