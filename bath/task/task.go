package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/bath/comm"
	"github.com/davyxu/cellnet/timer"
)

func LoopTask() {

	q := cellnet.NewEventQueue()
	q.EnableCapturePanic(true)

	q.StartLoop()

	rec_w_max := comm.GetDataStr("HGET", "bathLevel:1", "rec_w_max")

	index := 0
	timer.NewLoop(q, time.Millisecond*200, func(loop *timer.Loop) {
		fmt.Println("---------------")
		// times--
		index--

		userid := comm.GetDataStr("LINDEX", "userOnline", strconv.Itoa(index))
		if userid == "" {
			index = 0
		} else {
			tmp := time.Now().Unix()
			customeridNew := time.Now().UnixNano() / 1e6

			rec_w := userid + ":rec_w:1" //前台等待区
			customerLen := comm.GetDataInt("LLEN", rec_w)
			customer_0 := comm.GetDataStr("LINDEX", rec_w, strconv.Itoa(0))
			customer_0_int, _ := strconv.ParseInt(customer_0, 10, 64)

			if tmp-customer_0_int/1e3 > 0 {
				max, _ := strconv.Atoi(rec_w_max)
				if max > customerLen {
					comm.SetList("LPUSH", rec_w, []int64{customeridNew})
					createCustomer(userid, customeridNew, tmp)
				}
			} else {
				fmt.Println("end")
			}
		}
	}, nil).Start()

	q.Wait()
}

//记录用户状态
func createCustomer(userid string, customeridNew int64, tmp int64) {
	// rec:1
	var info map[string]string
	info = make(map[string]string)

	itemParam := []string{"items", "rec:1"}

	fmt.Println(itemParam)
	info = comm.GetHash(itemParam)
	kv := strings.Split(info["rec:1"], "^")
	customeridNewStr := strconv.FormatInt(customeridNew, 10)
	tmpStr := strconv.FormatInt(tmp, 10)
	process := tmpStr + "^" + kv[3] + "^-1^0^0^0^0"
	comm.SetDataStr("HSET", userid+":customer:1", customeridNewStr, process)
	fmt.Println(process)

}
func main() {
	LoopTask()
}
