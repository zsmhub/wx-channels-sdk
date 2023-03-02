package callbacks

var callbackParseExtraInfoMap = make(map[string]CallbackExtraInfoInterface)

type CallbackExtraInfoInterface interface {
	GetMessageType() string
	GetEventType() string
	GetTypeKey() string
	ParseFromJson(data []byte) (CallbackExtraInfoInterface, error)
}

func supportCallback(item CallbackExtraInfoInterface) {
	callbackParseExtraInfoMap[item.GetTypeKey()] = item
}
