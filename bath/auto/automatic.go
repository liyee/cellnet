package main

import (
	"time"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/bath/model"
	"github.com/davyxu/cellnet/timer"
)

func main() {
	q := cellnet.NewEventQueue()
	q.EnableCapturePanic(true)

	q.StartLoop()
	timer.NewLoop(q, time.Millisecond*1000, func(loop *timer.Loop) {
		model.Service{"rec", "rec", "chr"}.Run() //当前服务项目->下一个等待区（前端表现为移动）
		model.Service{"chr", "chr", "bap"}.Run()
		model.Service{"bap", "bap", "spy1sau"}.Run()
		model.Service{"spy1sau", "spy", "spy1sau"}.Run()
		model.Service{"spy1sau", "sau", "spy1sau"}.Run()
	}, nil).Start()

	q.Wait()
}
