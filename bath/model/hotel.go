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
			currentMax := Item{name + ":1", "max"}.GetItemOne()                                        //最大服务人数
			currentNum := comm.ListCommand{"LLEN", []string{userid + ":" + name + ":1"}}.LcommandInt() //当前服务中人数
			if currentMax > currentNum {
				pullCustomer(name, userid)
			}
			if currentNum > 0 { //服务内有客户时，需要进行检查是否完成
				recList := comm.ListCommand{"LRANGE", []string{userid + ":" + name + ":1", "0", "3"}}.Lcommand()
				//2.获取用户详情
				recListNew := make([]string, len(recList)+1, len(recList)+1)
				recListNew[0] = userid + ":customer:1"
				copy(recListNew[1:], recList)

				customer := comm.ListCommand{Cmd: "HMGET", Strs: recListNew}.Lcommand()
				checkCustomer(userid, recList, customer, name, nameNext)
			}

		}
	}
}

//检查客户服务情况
func checkCustomer(userid string, recList []string, customer []string, name string, nameNext string) {
	tmp := time.Now().Unix()
	if customer != nil {
		for i, userInfo := range customer {
			customerid := recList[i]
			fmt.Println("customerid:", customerid)
			detail := strings.Split(userInfo, "^")
			updatetime, _ := strconv.ParseInt(detail[0], 10, 64)
			waitingtime, _ := strconv.ParseInt(detail[1], 10, 64)
			if tmp-updatetime >= waitingtime {
				nextok := checkWait(userid, nameNext)
				if nextok == true {
					changeCustomer(userid, customerid, name, "w")
				}
			}
		}
	}
}

//检查等待区人数
func checkWait(userid, nameNext string) bool {
	currentlen := comm.ListCommand{"LLEN", []string{userid + ":" + nameNext + "_w:1"}}.LcommandInt()
	waitNum := Wait{nameNext + "_max"}.GetWaitNum()
	if currentlen < waitNum {
		return true
	}
	return false
}

//从等待区拉取客户
func pullCustomer(name, userid string) {
	customerid := comm.ListCommand{"RPOP", []string{userid + ":" + name + "_w:1"}}.LcommandStr()
	if customerid != "" {
		changeCustomer(userid, customerid, name, "p")
	}

}

/*更改用户服务流程程*/
func changeCustomer(userid, customerid, name, step string) {
	//updatetime^waitingtime^rec^chr^bap^spy^sua
	tmp := time.Now().Unix()
	keys := make(map[string]int)
	keys["rec"] = 2
	keys["chr"] = 3
	keys["bap"] = 4
	keys["spy"] = 5
	keys["sua"] = 6
	var nextName string

	switch name {
	case "rec":
		nextName = "chr"
	case "chr":
		nextName = "bap"
	case "bap":
		nextName = comm.Prob()
	default:
	}

	if step == "p" {
		comm.ListCommand{"LPUSH", []string{userid + ":" + name + ":1", customerid}}.CommandExec() //添加到服务区

	} else if step == "w" {
		comm.ListCommand{"LREM", []string{userid + ":" + name + ":1", strconv.Itoa(0), customerid}}.CommandExec() //从服务区移除
		comm.ListCommand{"LPUSH", []string{userid + ":" + nextName + "_w:1", customerid}}.CommandExec()           //添加下一个等待区
	}

	customerInfo := comm.ListCommand{"HGET", []string{userid + ":customer:1", customerid}}.LcommandStr() //获取用户详情
	detail := strings.Split(customerInfo, "^")
	detail[0] = strconv.FormatInt(tmp, 10)

	if step == "p" {
		speed := Item{name + ":1", "speed"}.GetItemOne()
		detail[1] = strconv.Itoa(speed)
		detail[keys[name]] = strconv.Itoa(2)
	} else {
		duration := Wait{name + "_duration"}.GetWaitNum()
		detail[1] = strconv.Itoa(duration)
		detail[keys[name]] = strconv.Itoa(1)
		detail[keys[nextName]] = strconv.Itoa(-1)
	}

	comm.ListCommand{"HSET", []string{userid + ":customer:1", customerid, strings.Join(detail, "^")}}.CommandExec() //更改用户状态
}
