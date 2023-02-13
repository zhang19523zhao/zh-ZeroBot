package main

import (
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"

	zero "github.com/zhang19523zhao/zh-ZeroBot"
	"github.com/zhang19523zhao/zh-ZeroBot/driver"
	_ "github.com/zhang19523zhao/zh-ZeroBot/example/command"
	_ "github.com/zhang19523zhao/zh-ZeroBot/example/music"
	_ "github.com/zhang19523zhao/zh-ZeroBot/example/regex"
)

func init() {
	log.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[zero][%time%][%lvl%]: %msg% \n",
	})
	log.SetLevel(log.DebugLevel)
}

func main() {
	zero.RunAndBlock(&zero.Config{
		NickName: []string{"bot"},
		// 命令前缀
		CommandPrefix: "/",
		SuperUsers:    []int64{1234},
		Driver: []zero.Driver{
			// 正向 WS
			driver.NewWebSocketClient("ws://127.0.0.1:6700", ""),
			// 反向 WS
			driver.NewWebSocketServer(16, "ws://127.0.0.1:6701", ""),
		},
	}, nil)
}
