# WecomBotGo
🤖 企业微信机器人Golang封装版本

[腾讯官方企业微信群机器人配置说明API](https://work.weixin.qq.com/api/doc/90000/90136/91770)

- [x] [文本类型](https://work.weixin.qq.com/api/doc/90000/90136/91770#文本类型)
- [x] [markdown类型](https://work.weixin.qq.com/api/doc/90000/90136/91770#markdown类型)
- [x] [图片类型](https://work.weixin.qq.com/api/doc/90000/90136/91770#图片类型)
- [x] [图文类型](https://work.weixin.qq.com/api/doc/90000/90136/91770#图文类型)
- [ ] [文件类型](https://work.weixin.qq.com/api/doc/90000/90136/91770#文件类型)

## 快速使用

可以查看 main_test.go 查看使用方法例子。

### 发送消息

```go
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
```

### 多个机器人

机器人列表为一个map，key为机器人名字，value为机器人的密钥。

#### 代码配置

你可以通过 **WecomBotGo.Bots["BotName"]="xxx-xxxxx-xxxxx-xxxxx"** 来设置对应的机器人。

````go
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
````

#### json配置

你可以通过编写一个形如

```json
[
    {
        "name": "yuki",
        "key": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
    },
    {
        "name": "haruhi",
        "key": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
    },
    {
        "name": "kyon",
        "key": "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
    }
]
```

的json文件，然后通过 **WecomBotGo.LoadBotsFromFile( "./xxxx.json" )** 函数来进行从json到 **WecomBotGo.Bots** 的转化，调用后可直接使用。

```go
WecomBotGo.Bots["yuki"].Send("text", "hello from bot send", nil, nil ) // just works
```

