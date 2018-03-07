package main

import (
	"log"

	"github.com/chanyipiaomiao/hltool"
)

func main() {
	dingtalk := hltool.NewDingTalkClient("https://oapi.dingtalk.com/robot/send?access_token=3a5569bf04362a5a0d14731e0ce8ba6359292c0033de14dce83c6ce14652548a", "> 消息内容", "markdown")
	ok, err := hltool.SendMessage(dingtalk)
	if err != nil {
		log.Fatalf("发送钉钉通知失败了: %s", err)
	} else {
		log.Fatalf("success: %t", ok)
	}
}
