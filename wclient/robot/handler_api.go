package robot

import (
	"strings"

	"github.com/opentdp/go-helper/request"
	"github.com/opentdp/wechat-rest/args"
	"github.com/opentdp/wechat-rest/wcferry"
)

func apiHandler() {

	helpFix := func(text string) string {
		lines := strings.Split(text, "\n")
		for k, line := range lines {
			line = strings.TrimLeft(line, "/")
			line = strings.Replace(line, "/", " ", 1)
			lines[k] = "/api " + strings.TrimSpace(line)
		}
		return strings.Join(lines, "\n")
	}

	handlers["/api"] = &Handler{
		Level:    0,
		ChatAble: true,
		RoomAble: true,
		Describe: "调用远程接口",
		Callback: func(msg *wcferry.WxMsg) string {
			str := strings.Replace(strings.TrimSpace(msg.Content), " ", "/", 1)
			res, err := request.TextGet("https://api.rehi.org/format=yaml/"+str, request.H{
				"User-Agent": args.AppName + "/" + args.Version,
			})
			if err != nil {
				return err.Error()
			}
			if str == "" || str == "help" {
				return helpFix(res)
			}
			if fileReply(msg, res) == 0 {
				return ""
			}
			return res
		},
	}

}