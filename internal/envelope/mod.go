package envelope

import (
	"crypto/rand"
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/zsmhub/wx-channels-sdk/internal/encryptor"
	"github.com/zsmhub/wx-channels-sdk/internal/signature"
	"io"
	"net/url"
)

type Processor struct {
	token         string
	encryptor     *encryptor.WXEncryptor
	entropySource io.Reader
	timeSource    TimeSource
}

func NewProcessor(token string, encodingAESKey string, opts ...ProcessorOption) (*Processor, error) {
	obj := Processor{
		token:         token,
		encryptor:     nil, // XXX init later
		entropySource: rand.Reader,
		timeSource:    DefaultTimeSource{},
	}
	for _, o := range opts {
		o.applyTo(&obj)
	}

	enc, err := encryptor.NewWXEncryptor(
		encodingAESKey,
		encryptor.WithEntropySource(obj.entropySource),
	)
	if err != nil {
		return nil, err
	}
	obj.encryptor = enc

	return &obj, nil
}

var errInvalidSignature = errors.New("invalid signature")

// 处理回调事件的消息：支持明文模式/安全模式，也支持xml/json
func (p *Processor) HandleIncomingMsg(url *url.URL, body []byte) (Envelope, error) {
	bodyStr := string(body)
	if len(bodyStr) == 0 {
		return Envelope{}, nil
	}

	// check signature
	if !signature.VerifyHTTPRequestSignature(p.token, url) {
		return Envelope{}, errInvalidSignature
	}

	x, err := p.ParseMessage(bodyStr)
	if err != nil {
		return Envelope{}, err
	}

	// 区分消息加密方式：明文模式/安全模式
	var msg encryptor.WXPayload
	if p.IsSafeMode(url) && x.Encrypt != "" {
		msg, err = p.encryptor.Decrypt([]byte(x.Encrypt))
		if err != nil {
			return Envelope{}, err
		}
	} else {
		msg.Msg = body
	}

	// assemble envelope to return
	return Envelope{
		ToUserName: x.ToUserName,
		Msg:        msg.Msg,
		ReceiveID:  msg.ReceiveID,
	}, nil
}

// 判断消息加密方式是不是安全模式
func (p *Processor) IsSafeMode(url *url.URL) bool {
	query := url.Query()
	return query.Get("signature") != "" && "aes" == query.Get("encrypt_type")
}

// 解析安全模式的回调事件结构体
func (p *Processor) ParseMessage(content string) (RxEnvelope, error) {
	// 区分数据格式：xml/json
	var x RxEnvelope
	if content[0:1] == "<" {
		if err := xml.Unmarshal([]byte(content), &x); err != nil {
			return x, err
		}
	} else {
		if err := json.Unmarshal([]byte(content), &x); err != nil {
			return x, err
		}
	}
	return x, nil

}
