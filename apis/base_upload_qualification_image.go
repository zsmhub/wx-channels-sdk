package apis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/url"
	"path"
)

// 上传资质图片
// 文档：https://developers.weixin.qq.com/doc/channels/API/basics/qualificationupload.html

type ReqQualificationUpload struct {
	URL string `json:"url"`
}

type ReqQualificationUploadMedia struct {
	Media *Media `json:"media"`
}

var _ mediaUploader = ReqQualificationUploadMedia{}

func (x ReqQualificationUploadMedia) getMedia() *Media {
	return x.Media
}

type RespQualificationUpload struct {
	CommonResp
	Data struct {
		FileId string `json:"file_id"`
	} `json:"data"`
}

var _ bodyer = RespQualificationUpload{}

func (x RespQualificationUpload) intoBody() ([]byte, error) {
	result, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *ApiClient) ExecQualificationUpload(req ReqQualificationUpload) (RespQualificationUpload, error) {
	var resp RespQualificationUpload

	// 下载文件，转成二进制
	_, body, err := FastClient.Get(nil, req.URL)
	if err != nil {
		return RespQualificationUpload{}, err
	}

	urlInfo, err := url.Parse(req.URL)
	if err != nil {
		return RespQualificationUpload{}, err
	}

	filename := fmt.Sprintf("%s%s", uuid.New().String(), path.Ext(urlInfo.Path))

	media, err := NewMediaFromBuffer(filename, body)
	if err != nil {
		return RespQualificationUpload{}, err
	}

	err = c.executeWXApiMediaUpload("/channels/ec/basics/qualification/upload", ReqQualificationUploadMedia{Media: media}, &resp, true)
	if err != nil {
		return RespQualificationUpload{}, err
	}
	if bizErr := resp.TryIntoErr(); bizErr != nil {
		return RespQualificationUpload{}, bizErr
	}
	return resp, nil
}

func NewMediaFromBuffer(filename string, buf []byte) (*Media, error) {
	stream := bytes.NewReader(buf)
	return &Media{
		filename: filename,
		filesize: int64(len(buf)),
		stream:   stream,
	}, nil
}
