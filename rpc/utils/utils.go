/*
 * @Author       : SpadesA.yanjuan9998@gmail.com
 * @Date         : 2022-06-16 18:19:06
 * @LastEditors  : SpadesA.yanjuan9998@gmail.com
 * @LastEditTime : 2022-06-21 13:43:15
 * @FilePath     : \antd-pro-amis-server\rpc\utils\utils.go
 */
package utils

import (
	"strconv"
	"strings"
)

func IfThen[T any](cond bool, trueVal T, falseVal T) T {
	if cond {
		return trueVal
	} else {
		return falseVal
	}
}

//string数组转int数组
func ToIntArray(str string) []int {
	arr := strings.Split(str, ",")
	var arrInt []int
	for _, v := range arr {
		i, _ := strconv.Atoi(v)
		arrInt = append(arrInt, i)
	}
	return arrInt
}
