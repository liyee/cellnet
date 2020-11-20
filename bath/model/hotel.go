package model

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/davyxu/cellnet/bath/comm"
)

type Hotel interface {
	Run()
	Next()
}

type Service struct {
	Name, Next string
}

func (rervice Service) Run() {
	name := rervice.Name
	nameNext := rervice.Next
	online := comm.ListCommand{"LRANGE", []string{"userOnline", "0", "100"}}.Lcommand()

	if online != nil {
		for _, userid := range online {
			//1.检查完成人员，2.补充人员
			//1.拉取队列数据
			recMax := Item{name, "max"}.GetItemOne()                                      //最大服务人数
			recNum := comm.ListCommand{"LLEN", []string{userid + ":rec:1"}}.LcommandInt() //当前服务中人数
			if recMax > recNum {

			}
			recList := comm.ListCommand{"LRANGE", []string{userid + ":rec:1", "0", "3"}}.Lcommand()
			//2.获取用户详情
			cDetailKey := userid + ":customer:1"

			app := recList[:]
			recListNew := make([]string, len(recList)+1, len(recList)+1)
			recListNew[0] = cDetailKey
			copy(recListNew[1:], app)
			commandCustomer := comm.ListCommand{Cmd: "HMGET", Strs: recListNew}
			customer := commandCustomer.Lcommand()
			checkCustomer(userid, recList, customer, name, nameNext)
		}
	}
}

//检查客户情况
//更新时间^等待时长^前台^更衣室^浴池^SPY^桑拿
//updatetime^waitingtime^rec^chr^bap^spy^sua
//1605579939^5^-1^0^0^0^0
func checkCustomer(userid string, recList []string, customer []string, name string, nameNext string) {
	fmt.Println("checkCustomer:", recList, name, nameNext)
	tmp := time.Now().Unix()
	if customer != nil {
		for i, userInfo := range customer {
			customerid := recList[i]
			fmt.Println("customerid:", customerid)
			detail := strings.Split(userInfo, "^")
			updatetime, _ := strconv.ParseInt(detail[0], 10, 64)
			waitingtime, _ := strconv.ParseInt(detail[1], 10, 64)
			if tmp-updatetime >= waitingtime {
				detail[0] = strconv.FormatInt(tmp, 10)
				detail[1] = strconv.Itoa(10)
				detail[2] = strconv.Itoa(1)
				detail[3] = strconv.Itoa(-1)
				userStr := strings.Join(detail, "^")
				fmt.Println("userStr:", userStr)

				nextok := checkWait(userid, nameNext)
				if nextok == true {
					//执行事务
					comm.TrCommand{"MULTI"}.Transaction()
					comm.ListCommand{"LREM", []string{userid + ":" + name + "_w:1", customerid}}.CommandExec() //从等待区取出
					comm.ListCommand{"LPUSH", []string{userid + ":" + name + ":1", customerid}}.CommandExec()  //添加到服务区
					comm.ListCommand{"HSET", detail}.CommandExec()                                             //更改用户状态

					comm.TrCommand{"EXEC"}.Transaction()
					fmt.Println("userStr:", nextok)
				}

			}
		}
	}
}

//检查等待区人数
func checkWait(userid, nameNext string) bool {
	len := comm.ListCommand{"LLEN", []string{userid + ":chr_w:1"}}.LcommandInt()
	waitNum := Wait{nameNext}.GetWaitNum()
	if len < waitNum {
		return true
	}
	return false
}

func (rec Rec) Next() {
	fmt.Println("I am rec, next!")
}
