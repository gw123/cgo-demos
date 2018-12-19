package main

import (
	"syscall"
	"fmt"
	"unsafe"
	"github.com/fpay/erp-client-s/common"
	"encoding/json"
)

func main() {
	module, err := syscall.LoadDLL("LibMP3DLL.dll")
	if err != nil {
		fmt.Println(err)
		return
	}
	start, err := module.FindProc("start")
	if err != nil {
		fmt.Println(err)
		return
	}

	config, err := syscall.BytePtrFromString("{}")
	if err != nil {
		fmt.Println(err)
	}
	call := func(event common.Event) {
		fmt.Println("callback", event.EventType)
	}
	status, _, err := start.Call(uintptr(unsafe.Pointer(config)), uintptr(unsafe.Pointer(&call)))
	if status != 0 {
		fmt.Println("start call", err)
	} else {
		fmt.Println("start call success")
	}

	///
	handle, err := module.FindProc("handle")
	if err != nil {
		fmt.Println(err)
		return
	}
	event:= common.Event{EventType:"",Payload:[]byte("123")}
	json ,err := json.Marshal(event)
	json = []byte(`{"MsgId":"","EventType":"","Payload":"123"}`)
	fmt.Println(string(json))

	//t, err := syscall.BytePtrFromString(string(json))
	//if err != nil {
	//	fmt.Println(err)
	//}
	//s, err := syscall.UTF16PtrFromString("hello")
	status, _, err = handle.Call(uintptr(unsafe.Pointer(&json[0])))
	if status != 0 {
		fmt.Println("handel call", err)
	} else {
		fmt.Println("handel call success")
	}

	///
	stop, err := module.FindProc("stop")
	if err != nil {
		fmt.Println(err)
		return
	}
	status, _, err = stop.Call()
	if status != 0 {
		fmt.Println("stop call", err)
	} else {
		fmt.Println("stop call success")
	}
}
