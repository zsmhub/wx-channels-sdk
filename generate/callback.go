package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/tidwall/gjson"
	"github.com/zsmhub/wx-channels-sdk/generate/tool"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

// 生成回调事件代码
type Callback struct {
	DocUrl     string
	Name       string
	StructName string
	JsonStr    string
	StructCode string
	MsgType    string
	EventType  string
}

var callbackDocVar = flag.String("doc", "", "[必填]微信文档地址")

func main() {
	flag.Parse()

	var docURL, savePath string

	if callbackDocVar != nil {
		docURL = *callbackDocVar
	}
	if docURL == "" {
		fmt.Println("请输入参数doc(微信文档地址):")
		_, _ = fmt.Scanf("%s", &docURL)
	}
	if docURL == "" {
		tool.Die("必传参数 doc=?")
	}

	doc := tool.GetDoc(docURL)

	titleHtml, err := doc.Find("title").Html()
	if err != nil {
		tool.Die("failed to get html: %+v\n", err)
	}
	titleHtml = titleHtml[:strings.Index(titleHtml, " ")]
	fmt.Println("开始抓取和生成回调事件代码，文档地址:", docURL)

	var cb Callback
	cb.DocUrl = docURL
	cb.Name = titleHtml

	cb.JsonStr = doc.Find("pre[class=language-json] > code").First().Text()
	if cb.JsonStr == "" {
		tool.Die("回调事件的json为空")
	}

	cb.MsgType = gjson.Get(cb.JsonStr, "MsgType").String()
	cb.EventType = gjson.Get(cb.JsonStr, "Event").String()
	eventSlice := strings.Split(cb.EventType, "_")
	for _, v := range eventSlice {
		cb.StructName += strcase.ToCamel(v)
	}
	savePath = fmt.Sprintf("./callbacks/%s%s.go", cb.MsgType, cb.Name)

	result, err := generateCallbackCode(cb)
	if err != nil {
		tool.Die("generateCallbackCode failed: %+v\n", err)
	}

	err = ioutil.WriteFile(savePath, result, os.ModePerm)
	if err != nil {
		tool.Die("ioutil.WriteFile failed: %+v\n", err)
	}

	addCallbackConstantToFile(cb)

	fmt.Printf("保存文件成功，文件路径: %s\n", savePath)
}

func generateCallbackCode(cb Callback) (result []byte, err error) {
	tpl, err := template.ParseFiles("./generate/template/callback.tmpl")
	if err != nil {
		fmt.Println("template.ParseFiles failed:", err)
		return
	}

	cb.StructCode, err = tool.GenerateStruct(cb.JsonStr, cb.StructName, false)
	if err != nil {
		fmt.Printf("generate StructCode failed: %+v\n", err)
		return
	}

	buf := bytes.NewBufferString("")
	err = tpl.Execute(buf, cb)
	if err != nil {
		return
	}

	result = buf.Bytes()
	return
}

func addCallbackConstantToFile(cb Callback) {
	filename := "./callbacks/callback_constant.go"
	content := `
// %s
const EventType%s EventType = "%s"
`
	tool.AddContentToFile(filename, fmt.Sprintf(content, cb.Name, cb.StructName, cb.EventType))
}
