package main

import (
	"log"

	"github.com/go-toast/toast"
)

func main() {
	notification := toast.Notification{
		AppID:   "Microsoft.Windows.Shell.RunDialog",
		Title:   "标题111",
		Message: "这是消息内容，等等。。。",
		Icon:    "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1549464117983&di=286ad42d05b9ea9720daa1d62cd18ee5&imgtype=0&src=http%3A%2F%2Fimgsrc.baidu.com%2Fimgad%2Fpic%2Fitem%2F8326cffc1e178a8208b90d86fc03738da977e80b.jpg", // 文件必须存在
		Actions: []toast.Action{
			{"protocol", "按钮1", "https://www.baidu.com/"},
			{"protocol", "按钮2", "https://baidu.com/"},
		},
	}
	err := notification.Push()
	if err != nil {
		log.Fatalln(err)
	}
}
