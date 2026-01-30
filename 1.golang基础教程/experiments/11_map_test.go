package experiments

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	//1.初始化
	var map1 = make(map[string]int)
	fmt.Println(map1)

	//2.赋值
	map1["key1"] = 2
	fmt.Println(map1)

	map1["key2"] = 3
	fmt.Println(map1)

	//3.取值
	var value = map1["key2"]
	fmt.Println(value)

	//4.删除某个key
	delete(map1, "key2")
	var value1 = map1["key2"]
	fmt.Println(value1)

	//5.清空map
	fmt.Println(map1)
	clear(map1)
	fmt.Println(map1)
}
