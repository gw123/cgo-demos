package main

import "C"
import (
	"syscall"
	"fmt"
	"unsafe"
	"github.com/gw123/cgo-demos/utils"
)

//测试MP3播报
func main() {
	module, err := syscall.LoadDLL("libtestModule.dll")
	if err != nil {
		fmt.Println(err)
		return
	}
	start, err := module.FindProc("start")
	if err != nil {
		fmt.Println(err)
		return
	}
	str := utils.ConvertToGbk("{'name':'测试'}")
	config, err := syscall.BytePtrFromString(str)
	if err != nil {
		fmt.Println(err)
	}

	call := func(event *C.char) uintptr {
		fmt.Println("callback", C.GoString(event))
		return 0
	}

	status, _, err := start.Call(uintptr(unsafe.Pointer(config)), uintptr(syscall.NewCallback(call)))
	if status != 0 {
		fmt.Println("start call", err)
	} else {
		fmt.Println("start call success")
	}

	///
	handel, err := module.FindProc("handel")
	if err != nil {
		fmt.Println(err)
		return
	}

	t, err := syscall.BytePtrFromString("{}")
	if err != nil {
		fmt.Println(err)
	}
	//s, err := syscall.UTF16PtrFromString("hello")
	status, _, err = handel.Call(uintptr(unsafe.Pointer(t)))
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
