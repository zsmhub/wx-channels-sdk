package callbacks

import (
	"encoding/json"
	"errors"
	"strings"
)

// CallbackMessage 一条接收到的消息
type CallbackMessage struct {
	// ToUserName 接收成员
	ToUserName string `json:"ToUserName"`
	// FromUserName 发送成员
	FromUserName string `json:"FromUserName"`
	// CreateTime 消息创建时间，秒级时间戳
	CreateTime int64 `json:"CreateTime"`
	// MsgType 消息类型
	MsgType MessageType `json:"MsgType"`
	// MsgID 消息id，64位整型
	MsgID int64 `json:"MsgId"`
	// EventType 事件类型 MsgType为event存在
	EventType EventType `json:"Event"`

	// 额外的信息
	Extras CallbackExtraInfoInterface `json:"-"`

	// 原始回调内容，方便排查问题
	OriginalMessage string `json:"OriginalMessage"`
}

func (m CallbackMessage) ParseMessageFromJson(body []byte) (CallbackMessage, error) {
	err := json.Unmarshal(body, &m)
	if err != nil {
		return m, err
	}

	m.OriginalMessage = string(body)

	if m.MsgType == "" {
		m.MsgType = MessageTypeUnknown
	}

	m.EventType = EventType(strings.Trim(string(m.EventType), " "))

	extraParser, ok := callbackParseExtraInfoMap[m.GetTypeKey()]
	if ok {
		m.Extras, err = extraParser.ParseFromJson(body)
	} else {
		err = errors.New("回调事件解析失败，请去SDK生成对应的回调事件：" + m.GetTypeKey() + "，json: " + m.OriginalMessage)
	}
	return m, err
}

func (m CallbackMessage) GetTypeKey() string {
	return string(m.MsgType) + ":" + string(m.EventType)
}

func (m CallbackMessage) GetStructName() string {
	return m.snakeToCamel(string(m.MsgType)) + m.snakeToCamel(string(m.EventType))
}

func (m CallbackMessage) snakeToCamel(s string) string {
	news := ""
	for k, v := range s {
		if k == 0 {
			news += strings.ToUpper(string(v))
			continue
		}

		if v == '_' {
			continue
		}

		if s[k-1] == '_' {
			news += strings.ToUpper(string(v))
		} else {
			news += string(v)
		}

	}
	return news
}
