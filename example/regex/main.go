package regex

import (
	"fmt"
	zero "github.com/zhang19523zhao/zh-ZeroBot"
	"github.com/zhang19523zhao/zh-ZeroBot/config"
	"github.com/zhang19523zhao/zh-ZeroBot/example/gpt"
	"github.com/zhang19523zhao/zh-ZeroBot/filter"
	"github.com/zhang19523zhao/zh-ZeroBot/message"
)

func init() {
	var (
		err    error
		replay string
	)
	conf := config.LoadConfig()
	zero.OnRegex(conf.Regexp).Handle(func(ctx *zero.Ctx) {
		msg := ctx.Event.RawMessage
		check, fmsg := filter.FilterWord(msg, conf.Filter)
		if check {
			replay, err = gtp.Completions(msg)
			if err != nil {
				fmt.Println("gpt err: ", err)
			}
		} else {
			replay = fmsg
		}
		// 发送消息
		ctx.Send(message.Text(replay))
	})

}
