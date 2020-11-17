package main

import (
	"fmt"

	"github.com/davyxu/cellnet/bath/model"
)

func main() {
	fmt.Println("rec start")
	hotel := model.Rec{}
	hotel.Run()

	// waitInfo := model.ItemLevelWaitInfo{}
	// rec_w_max := waitInfo.GetSilgleBathLevel("rec_w_max") //等待人数

	// rp := model.RecProccess{}
	// num := rp.Len() //服务人数

	// fmt.Println("num:", rec_w_max, num)
}
