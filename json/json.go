//yongze.chen json 处理
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	js := make(map[string]interface{})
	arrList := make([]interface{}, 0)

	arr1 := make([]interface{}, 0)
	arr1 = append(arr1, "data1")
	arr1 = append(arr1, 14)

	arr2 := make([]interface{}, 0)
	arr2 = append(arr2, "data2")
	arr2 = append(arr2, "red")

	arrList = append(arrList, arr1)
	arrList = append(arrList, arr2)

	js["params"] = arrList

	bt, _ := json.Marshal(js)

	fmt.Println(string(bt))
}
