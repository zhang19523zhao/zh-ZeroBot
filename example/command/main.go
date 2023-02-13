package command

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
		replay string
		err    error
	)

	//command := config.LoadConfig().Command
	conf := config.LoadConfig()
	zero.OnCommand(conf.Command).Handle(func(ctx *zero.Ctx) {

		//fmt.Println("test", ctx.Event.RawMessage)
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
