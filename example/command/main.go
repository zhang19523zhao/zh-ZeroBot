package command

import (
	"fmt"
	zero "github.com/zhang19523zhao/zh-ZeroBot"
	"github.com/zhang19523zhao/zh-ZeroBot/config"
	"github.com/zhang19523zhao/zh-ZeroBot/example/gpt"
	"github.com/zhang19523zhao/zh-ZeroBot/message"
)

func init() {
	command := config.LoadConfig().Command
	zero.OnCommand(command).Handle(func(ctx *zero.Ctx) {

		//fmt.Println("test", ctx.Event.RawMessage)
		msg := ctx.Event.RawMessage
		replay, err := gtp.Completions(msg)
		if err != nil {
			fmt.Println("gpt err: ", err)
		}
		// 发送消息
		ctx.Send(message.Text(replay))
	})
}
