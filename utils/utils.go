package utils

import "github.com/axgle/mahonia"

/***
 * 将utf-8 转码 gbk
 */
func ConvertToGbk(src string) string {
	enc := mahonia.NewEncoder("gbk")
	str := enc.ConvertString(src)
	return str
}
