package main

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
func setData(command string, name string, key string, value string) {
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

//获取单个数据
func getData(command string, name string, value string) {
	p.(cellnet.RedisPoolOperator).Operate(func(rawClient interface{}) interface{} {
		client := rawClient.(*redis.Client)
		fmt.Println(name, value)
		v, error := client.Cmd(command, name, value).Str()
		if error == nil {
			return v
		}
		return nil
	})
}

//初始化数据
func getBathInfo(name string, key1 string, key2 string, key3 string, key4 string, key5 string, key6 string, key7 string) (data string) {
	var bath = make(map[string]string)
	p.(cellnet.RedisPoolOperator).Operate(func(rawClient interface{}) interface{} {
		client := rawClient.(*redis.Client)
		fmt.Println(name, key1, key2, key3)
		v, error := client.Cmd("HMGET", name, key1, key2, key3, key4, key5, key6, key7).Array()
		if error == nil {
			var bathKey = [7]string{"level", "earnings", "rec_num", "chr_num", "bap_num", "sau_num", "spy_num"}
			for k := range v {
				elemStr, _ := v[k].Str()
				bath[bathKey[k]] = elemStr
			}
		}
		return bath
	})
	data = MapToJson(bath)
	return
}

func MapToJson(param map[string]string) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}
