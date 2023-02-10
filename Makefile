# 生成API代码
api:
	go run generate/api.go -doc=$(doc) -prefix=$(prefix)

# 生成回调事件代码
callback:
	go run generate/callback.go -doc=$(doc)