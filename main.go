// 更为具体的参数说明，请查看
// https://work.weixin.qq.com/api/doc/90000/90136/91770
package WecomBotGo

import (
	"bytes"
	"encoding/json"
	"github.com/kpango/glg"
	"io/ioutil"
	"net/http"
)

type (
	H map[string]interface{}
	Bot struct {
		Key string
	}
	Article struct {
		Title string `json:"title"`
		Description string `json:"description"`
		Url string `json:"url"`
		Picurl string `json:"picurl"`
	}
	bot struct {
		Name string `json:"name"`
		Key string `json:"key"`
	}
)

var (
	Bots map[string] Bot
)

func init() {
	Bots = make( map[string]Bot )
}

func LoadBotsFromFile( path string ) error {
	var bots []bot
	data, e := ioutil.ReadFile( path )
	if e = json.Unmarshal( data, &bots ); e != nil {
		glg.Log( e )
		return e
	} else {
		for _, b := range bots {
			Bots[b.Name] = Bot {Key: b.Key}
		}
		return nil
	}
}

func (b Bot) Send( msgType, content string, mentionedList, mentionedMobileList []string ) ( *http.Response, error ) {
	return Send( b.Key, msgType, content, mentionedList, mentionedMobileList )
}

func (b Bot) SendImage( base64, md5 string ) ( *http.Response, error ) {
	return SendImage( b.Key, base64, md5 )
}

func (b Bot) SendMarkdown( markdown string ) ( *http.Response, error ) {
	return SendMarkdown( b.Key, markdown )
}

func (b Bot) SendText( content string, mentionedList, mentionedMobileList []string ) ( *http.Response, error ) {
	return Send( b.Key, "text", content, mentionedList, mentionedMobileList )
}

func (b Bot) SendNews( articles []Article ) ( *http.Response, error ) {
	return SendNews( b.Key, articles )
}

func SendMarkdown( key, markdown string ) ( *http.Response, error ) {
	var (
		bj []byte
		e error
	)
	if bj, e = json.Marshal( &H {
		"msgtype": "markdown",
		"markdown": &H {
			"content": markdown,
		},
	} ); e != nil {
		return nil, e
	}
	return SendByte( key, bj )
}

func SendImage( key, base64, md5 string ) ( *http.Response, error ) {
	var (
		bj []byte
		e error
	)
	if bj, e = json.Marshal( &H {
		"msgtype": "image",
		"image": &H {
			"base64": base64,
			"md5": md5,
		},
	} ); e != nil {
		return nil, e
	}
	return SendByte( key, bj )
}

func SendNews( key string, articles []Article ) ( *http.Response, error ) {
	var (
		bj []byte
		e error
	)
	if bj, e = json.Marshal( &H {
		"msgtype": "news",
		"news": &H {
			"articles": articles,
		},
	} ); e != nil {
		return nil, e
	}
	return SendByte( key, bj )
}

func SendText( key, content string, mentionedList, mentionedMobileList []string ) ( *http.Response, error ) {
	return Send( key, "text", content, mentionedList, mentionedMobileList )
}

func Send( key, msgType, content string, mentionedList, mentionedMobileList []string ) ( *http.Response, error ) {
	var (
		bj []byte
		e error
	)
	if bj, e = json.Marshal( &H {
		"msgtype": msgType,
		msgType: &H {
			"content": content,
			"mentioned_list": mentionedList,
			"mentioned_mobile_list": mentionedMobileList,
		},
	} ); e != nil {
		return nil, e
	}

	return SendByte( key, bj )
}

func SendByte( key string, jsonByte[]byte ) ( *http.Response, error ) {
	if resp, e := http.Post("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + key,
		"application/json",
		bytes.NewBuffer( jsonByte ) ); e != nil {
		return nil, e
	} else {
		return resp, nil
	}
}