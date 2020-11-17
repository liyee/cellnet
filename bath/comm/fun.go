package comm

import (
	"crypto/rand"
	"strings"

	"math/big"
)

//概率事件
func Prob() string {
	var items = [4]string{"sau", "spy", "sau1spy", "sau0spy"}
	var item string
	result, _ := rand.Int(rand.Reader, big.NewInt(100))
	//fmt.Println(result)
	num := result.Uint64() / uint64(10)

	switch num {
	case 8, 9:
		item = items[3]
	case 4, 5, 6, 7:
		item = items[2]
	case 2, 3:
		item = items[1]
	case 0, 1:
		item = items[0]
	}
	return item
}

//解析项目数据
func UnItems(items map[string]string) ItemLevelInfo {
	var levelInfo ItemLevelInfo
	var itemV map[string]string

	for _, value := range items {
		v := strings.Split(value, "^")

		itemV["max"] = v[0]
		itemV["cost"] = v[1]
		itemV["exp"] = v[2]
		itemV["speed"] = v[3]
		itemV["income"] = v[4]

		levelInfo.rec = itemV
	}

	return levelInfo
}
