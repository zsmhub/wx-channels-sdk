package apis

import (
	"io"
	"net/url"
)

// urlValuer 可转化为 url.Values 类型的 trait
type urlValuer interface {
	intoURLValues() url.Values
}

// Bodyer 可转化为 API 请求体的 trait
type bodyer interface {
	intoBody() ([]byte, error)
}

// mediaUploader 携带 *Media 对象，可转化为 multipart 文件上传请求体的 trait
type mediaUploader interface {
	getMedia() *Media
}

// 欲上传的图片等媒体
type Media struct {
	filename string
	filesize int64
	stream   io.Reader
}
