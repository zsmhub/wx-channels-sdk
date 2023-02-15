package callbacks

import (
    "errors"
    "github.com/zsmhub/wx-channels-sdk/internal/envelope"
    "github.com/zsmhub/wx-channels-sdk/internal/signature"
    "io/ioutil"
    "net/http"
    "net/url"
    "strconv"
)

type callbackUrlVars struct {
    Signature string
    Timestamp int64
    Nonce     string
    EchoStr   string
}

type CallbackHandler struct {
    token string // 回调 token
    ep    *envelope.Processor
}

func NewCallbackHandler(token string, encodingAESKey string) (*CallbackHandler, error) {
    ep, err := envelope.NewProcessor(token, encodingAESKey)
    if err != nil {
        return nil, err
    }
    return &CallbackHandler{token: token, ep: ep}, nil
}

// 解析并获取回调数据
func (cb *CallbackHandler) GetCallbackMsg(r *http.Request) (CallbackMessage, error) {
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        return CallbackMessage{}, err
    }

    // 验签
    ev, err := cb.ep.HandleIncomingMsg(r.URL, body)
    if err != nil {
        return CallbackMessage{}, err
    }

    // 解析json
    message, err := CallbackMessage{}.ParseMessageFromJson(ev.Msg)
    if err != nil {
        return message, err
    }

    return message, nil
}

// 后台回调配置URL，申请校验
func (cb *CallbackHandler) EchoTestHandler(rw http.ResponseWriter, r *http.Request) {
    if !signature.VerifyHTTPRequestSignature(cb.token, r.URL) {
        rw.WriteHeader(http.StatusBadRequest)
        return
    }

    args, err := cb.parseUrlVars(r.URL.Query())
    if err != nil {
        rw.WriteHeader(http.StatusBadRequest)
        return
    }

    rw.WriteHeader(http.StatusOK)
    _, _ = rw.Write([]byte(args.EchoStr))
}

func (cb *CallbackHandler) parseUrlVars(urlVars url.Values) (callbackUrlVars, error) {
    var errMalformedArgs = errors.New("malformed arguments for echo test API")

    var msgSignature string
    {
        l := urlVars["signature"]
        if len(l) != 1 {
            return callbackUrlVars{}, errMalformedArgs
        }
        msgSignature = l[0]
    }

    var timestamp int64
    {
        l := urlVars["timestamp"]
        if len(l) != 1 {
            return callbackUrlVars{}, errMalformedArgs
        }
        timestampStr := l[0]

        timestampInt, err := strconv.ParseInt(timestampStr, 10, 64)
        if err != nil {
            return callbackUrlVars{}, errMalformedArgs
        }

        timestamp = timestampInt
    }

    var nonce string
    {
        l := urlVars["nonce"]
        if len(l) != 1 {
            return callbackUrlVars{}, errMalformedArgs
        }
        nonce = l[0]
    }

    var echoStr string
    {
        l := urlVars["echostr"]
        if len(l) != 1 {
            return callbackUrlVars{}, errMalformedArgs
        }
        echoStr = l[0]
    }

    return callbackUrlVars{
        Signature: msgSignature,
        Timestamp: timestamp,
        Nonce:     nonce,
        EchoStr:   echoStr,
    }, nil
}
