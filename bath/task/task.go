package main

import (
	"fmt"
	"time"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/timer"
)

func LoopTask() {

	q := cellnet.NewEventQueue()
	q.EnableCapturePanic(true)

	q.StartLoop()

	var times = 3

	timer.NewLoop(q, time.Millisecond*100, func(loop *timer.Loop) {
		times--
		if times == 0 {
			loop.Stop()
			q.StopLoop()
		}

		fmt.Println("before")

		fmt.Println("after")

	}, nil).Start()

	q.Wait()
}

func main() {
	LoopTask()
}
