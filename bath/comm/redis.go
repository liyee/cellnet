package comm

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	_ "github.com/davyxu/cellnet/peer/redix"
	"github.com/mediocregopher/radix.v2/redis"
)

var (
	p = peer.NewGenericPeer("redix.Connector", "redis", "10.8.189.203:6379", nil)
)

func init() {
	p.(cellnet.RedisConnector).SetConnectionCount(1)
	p.Start()

	for i := 0; i < 5; i++ {

		if p.(cellnet.PeerReadyChecker).IsReady() {
			break
		}

		time.Sleep(time.Millisecond * 200)
	}
}

//设置单个数据
func SetData(command string, name string, key string, value string) {
	newNum, error := strconv.Atoi(value)
	if error == nil {
		//return newNum
	}
	p.(cellnet.RedisPoolOperator).Operate(func(rawClient interface{}) interface{} {
		client := rawClient.(*redis.Client)
		fmt.Println(key, value)
		client.Cmd(command, name, key, newNum)
		return nil
	})
}

//设置单个数据
func SetDataStr(command string, name string, key string, value string) {
	p.(cellnet.RedisPoolOperator).Operate(func(rawClient interface{}) interface{} {
		client := rawClient.(*redis.Client)
		fmt.Println(key, value)
		client.Cmd(command, name, key, value)
		return nil
	})
}

//获取单个数据
func GetDataStr(command string, name string, key string) string {
	var str string
	p.(cellnet.RedisPoolOperator).Operate(func(rawClient interface{}) interface{} {
		client := rawClient.(*redis.Client)
		fmt.Println(command, name)
		v, error := client.Cmd(command, name, key).Str()
		if error == nil {
			str = v
		}
		return str
	})

	return str
}

//获取单个数据
func GetDataInt(command string, name string) int {
	var str int
	p.(cellnet.RedisPoolOperator).Operate(func(rawClient interface{}) interface{} {
		client := rawClient.(*redis.Client)
		fmt.Println(command, name)
		v, error := client.Cmd(command, name).Int()
		if error == nil {
			str = v
		}
		return str
	})

	return str
}

//初始化数据
func GetHash(strs []string) map[string]string {
	var bath = make(map[string]string)
	p.(cellnet.RedisPoolOperator).Operate(func(rawClient interface{}) interface{} {
		client := rawClient.(*redis.Client)

		args := make([]interface{}, len(strs))
		for i, s := range strs {
			args[i] = s
		}

		v, error := client.Cmd("HMGET", args).Array()
		if error == nil {
			for k := range v {
				elemStr, _ := v[k].Str()
				bath[strs[k+1]] = elemStr
			}
		}
		return bath
	})

	return bath
}

//队列添加数值
func SetList(command string, name string, strs []int64) {
	p.(cellnet.RedisPoolOperator).Operate(func(rawClient interface{}) interface{} {
		client := rawClient.(*redis.Client)
		fmt.Println(command, name)

		args := make([]interface{}, len(strs))
		for i, s := range strs {
			args[i] = s
		}

		client.Cmd(command, name, args)
		return nil
	})
}

func MapToJson(initInfo Info) string {
	dataType, _ := json.Marshal(initInfo)
	dataString := string(dataType)
	return dataString
}
