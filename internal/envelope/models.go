package envelope

// 安全模式的回调事件结构体
type RxEnvelope struct {
	ToUserName string `xml:"ToUserName" json:"ToUserName"`
	Encrypt    string `xml:"Encrypt" json:"Encrypt"`
}

type Envelope struct {
	ToUserName string
	Msg        []byte
	ReceiveID  []byte
}
