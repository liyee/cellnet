package model

import (
	"fmt"
	"strconv"

	"github.com/davyxu/cellnet/bath/comm"
)

type Hotel interface {
	Run()
	Next()
}

type Rec struct {
}

func (rec Rec) Run() {
	// fmt.Println("rec")
	// waitInfo := ItemLevelWaitInfo{}
	// rec_w_max := waitInfo.GetSilgleBathLevel("rec_w_max") //等待人数

	rp := RecProccess{}
	num := rp.Len() //服务人数

	if num > 0 {
		userid := comm.GetDataInt("RPOP", "1:rec_w:1")
		useridStr := strconv.Itoa(userid)
		userInfo := comm.GetHash([]string{"1:customer:1", useridStr})

		fmt.Println(userInfo)
	}

}

func (rec Rec) Next() {
	fmt.Println("I am rec, next!")
}
