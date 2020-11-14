package WecomBotGo

import (
	"io/ioutil"
	"testing"
)

var (
	key string
)

func init() {
	data, _ := ioutil.ReadFile("./key.txt")
	key = string(data)
}

func TestLoadBotsFromFile(t *testing.T) {
	e := LoadBotsFromFile( "./bots.json" )
	if e != nil {
		println( e )
		t.FailNow()
	}
	for k, v := range Bots {
		println( k, v.Key )
	}
}

func TestSend(t *testing.T) {
	Send( key, "text", "hello send", []string {"@all"}, nil )
}

func TestBot_Send(t *testing.T) {
	Bots["yuki"] = Bot{ Key: key }
	Bots["yuki"].Send("text", "hello from bot send", nil, nil )
	Bots["yuki"].SendNews( []Article {
		{
			Title:       "中秋节礼品领取",
			Description: "今年中秋节公司有豪礼相送",
			Url:         "www.qq.com",
			Picurl:      "http://res.mail.qq.com/node/ww/wwopenmng/images/independent/doc/test_pic_msg1.png",
		},
	})
	Bots["yuki"].SendMarkdown( `
		实时新增用户反馈<font color=\"warning\">132例</font>，请相关同事注意。\n
         >类型:<font color=\"comment\">用户反馈</font>
         >普通用户反馈:<font color=\"comment\">117例</font>
         >VIP用户反馈:<font color=\"comment\">15例</font>
	`)
}

