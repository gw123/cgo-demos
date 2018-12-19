package main

/*
#cgo CFLAGS: -I .
#include "cgo_wrap.h"
*/
import "C"

import (
	"fmt"
)

var g_handel_fun C.t_handel
var g_config string

//export Start
func Start(config *C.char, fn C.t_handel) int {
	g_handel_fun = fn
	g_config = C.GoString(config)

	fmt.Printf("Go.Start(): config = %s\n", g_config)
	return 1
}

//export Push
func Push(event *C.char) {
	//把消息推到总线上面去
	fmt.Printf("Go.Push(): config = %s\n", C.GoString(event))
	C.call_handel(g_handel_fun, event)
}

func main() {
	//fmt.Printf("Go.main(): calling C function with callback to us\n")
	//C.some_c_func((C.callback_fcn)(unsafe.Pointer(C.callOnMeGo_cgo)))
}
