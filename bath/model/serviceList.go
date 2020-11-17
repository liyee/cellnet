//服务队列处理
package model

import (
	// "strconv"

	"github.com/davyxu/cellnet/bath/comm"
)

type ServiceListFunc interface {
	Len() int
	Lrem()
	Lpush()
	Lrange()
}

type RecProccess struct {
}

func (rp RecProccess) Len() int {
	val := comm.GetDataInt("LLEN", "1:rec_w:1")
	return val
}
