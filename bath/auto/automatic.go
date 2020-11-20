package main

import (
	"fmt"

	"github.com/davyxu/cellnet/bath/model"
)

func main() {
	fmt.Println("rec start")

	//前台服务
	model.Service{Name: "rec", Next: "chr"}.Run()
}
