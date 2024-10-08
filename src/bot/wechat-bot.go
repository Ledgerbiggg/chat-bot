package bot

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"os"
	"time"
)

type WechatBot struct {
	// 主人名字
	hostName string
}

func NewWechatBot() *WechatBot {
	w := &WechatBot{}
	return w
}

func (mb *WechatBot) Start() {
	bot := openwechat.DefaultBot(openwechat.Desktop)

	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl
	// 创建热存储容器对象
	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")
	// 注册消息处理函数
	bot.MessageHandler = func(msg *openwechat.Message) {
		if msg.IsText() && msg.Content == "ping" {
			text, _ := msg.ReplyText("pong")

			file, err := os.OpenFile("storage.json", os.O_RDONLY, 0)
			if err != nil {
				fmt.Println(err)
			}
			time.Sleep(5 * time.Second)
			text.Revoke()
			msg.ReplyFile(file)
		}
	}
	defer reloadStorage.Close()

	// 执行热登录
	bot.HotLogin(reloadStorage, openwechat.NewRetryLoginOption())
	self, _ := bot.GetCurrentUser()
	friends, _ := self.Friends()
	for _, v := range friends {
		if v.NickName == "ledgerbiggg" {
			self.SendTextToFriend(v, "机器人已经成功登录!!!")
		}
	}
	fmt.Println(friends)
	bot.Block()
}
