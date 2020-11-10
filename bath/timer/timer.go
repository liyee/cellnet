package main

import (
	"fmt"
	"time"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/timer"
)

func LoopTimer() {
	queue := cellnet.NewEventQueue()

	// 启动消息循环
	queue.StartLoop()

	var count int

	// 启动计时循环
	timer.NewLoop(queue, time.Millisecond*10, func(ctx *timer.Loop) {
		count++
		fmt.Println("Hello, World!")
		if count >= 10 {
			ctx.Stop()
		}
	}, nil).Start()
}

func main() {
	LoopTimer()
}
