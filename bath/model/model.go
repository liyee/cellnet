package model

import (
	"strconv"
	"strings"

	"github.com/davyxu/cellnet/bath/comm"
)

type ModelInfo interface {
	GetItemOne() int
	GetItemAll() map[string]string

	GetWaitNum() int
}

type Item struct {
	Name string
	Key  string
}

type Wait struct {
	Name string
}

func (wait Wait) GetWaitNum() int {
	command := comm.ListCommand{Cmd: "HGET", Strs: []string{"bathLevel:1", wait.Name}}
	val := command.LcommandInt()
	return val
}

// 最大值 价格  产生经验 时长/每人 收益/每人
// max	  cost exp	  speed	   income
var itemkey = [5]string{"max", "cost", "exp", "speed", "income"}
var items map[string]int

func (para Item) GetItemOne() int {
	items = make(map[string]int)
	command := comm.ListCommand{Cmd: "HGET", Strs: []string{"items", para.Name}}
	itemStr := command.LcommandStr()
	itemArr := strings.Split(itemStr, "^")
	for i, item := range itemArr {
		v, _ := strconv.Atoi(item)
		items[itemkey[i]] = v
	}
	return items[para.Key]

}

func (para Item) GetItemAll() map[string]int {
	items = make(map[string]int)
	command := comm.ListCommand{Cmd: "HGET", Strs: []string{"items", para.Name}}
	itemStr := command.LcommandStr()
	itemArr := strings.Split(itemStr, "^")
	for i, item := range itemArr {
		v, _ := strconv.Atoi(item)
		items[itemkey[i]] = v
	}
	return items
}
