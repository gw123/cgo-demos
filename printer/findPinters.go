package main

import (
	"syscall"
	"unsafe"
	"strings"
	"fmt"
)

var (
	printerlibs  = syscall.MustLoadDLL("PrinterLibs.dll")
	Port_EnumUSB = printerlibs.MustFindProc("Port_EnumUSB")
	Port_EnumCOM = printerlibs.MustFindProc("Port_EnumCom")
	Port_EnumLPT = printerlibs.MustFindProc("Port_EnumLpt")
	Port_EnumPRN = printerlibs.MustFindProc("Port_EnumPrn")
)

func GetUsbList() ([]string, error) {
	num := 0
	pcnReturned := 0

	success, _, err := Port_EnumUSB.Call(0, 0, uintptr(unsafe.Pointer(&num)), uintptr(unsafe.Pointer(&pcnReturned)))
	if num == 0 {
		fmt.Println("Num ",num)
		return nil, nil
	}

	buffer := make([]byte, num)
	success, _, err = Port_EnumUSB.Call(uintptr(unsafe.Pointer(&buffer[0])), 1024, uintptr(unsafe.Pointer(&num)), uintptr(unsafe.Pointer(&pcnReturned)))
	if success == 0 {
		return nil, err
	}

	itemLen := num / pcnReturned
	fmt.Println("item len ", itemLen)
	list := make([]string, 0)
	start := 0
	end := 0
	count := 0
	for i := 0; i < num; i++ {
		if buffer[i] == 0 {
			end = i
			list = append(list, strings.Trim(string(buffer[start:end]), " "))
			start = i+1
			count++
			if count >= pcnReturned {
				break
			}
		}
		//fmt.Printf("%d %c \t", buffer[i], buffer[i])
	}
	return list, nil
}

func GetLPTList() ([]string, error) {
	num := 10
	pcnReturned := 10

	success, _, err := Port_EnumLPT.Call(0, 0, uintptr(unsafe.Pointer(&num)), uintptr(unsafe.Pointer(&pcnReturned)))
	if num == 0 {
		return nil, nil
	}

	buffer := make([]byte, num)
	success, _, err = Port_EnumLPT.Call(uintptr(unsafe.Pointer(&buffer[0])), 1024, uintptr(unsafe.Pointer(&num)), uintptr(unsafe.Pointer(&pcnReturned)))
	if success == 0 {
		return nil, err
	}
	//fmt.Println(num, pcnReturned, string(buffer))
	strArr := strings.Split(string(buffer), " ")
	return strArr, nil
}

func GetPRNList() ([]string, error) {
	num := 10
	pcnReturned := 10

	success, _, err := Port_EnumPRN.Call(0, 0, uintptr(unsafe.Pointer(&num)), uintptr(unsafe.Pointer(&pcnReturned)))
	if num == 0 {
		return nil, nil
	}

	buffer := make([]byte, num)
	success, _, err = Port_EnumPRN.Call(uintptr(unsafe.Pointer(&buffer[0])), 1024, uintptr(unsafe.Pointer(&num)), uintptr(unsafe.Pointer(&pcnReturned)))
	if success == 0 {
		return nil, err
	}
	//fmt.Println(num, pcnReturned, string(buffer))
	strArr := strings.Split(string(buffer), " ")
	return strArr, nil
}

func GetComList() ([]string, error) {
	num := 0
	pcnReturned := 0

	success, _, err := Port_EnumCOM.Call(0, 0, uintptr(unsafe.Pointer(&num)), uintptr(unsafe.Pointer(&pcnReturned)))
	if num == 0 {
		return nil, nil
	}

	buffer := make([]byte, num)
	success, _, err = Port_EnumCOM.Call(uintptr(unsafe.Pointer(&buffer[0])), 1024, uintptr(unsafe.Pointer(&num)), uintptr(unsafe.Pointer(&pcnReturned)))
	if success == 0 {
		return nil, err
	}
	//fmt.Println(num, pcnReturned, string(buffer))
	strArr := strings.Split(string(buffer), " ")
	return strArr, nil
}


func main() {
	//usb设备
	usbPathList, err := GetUsbList()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("usb 设备")
		for index, val := range usbPathList {
			fmt.Printf("[%d]%s||\n", index, val)
			//conn ,err := connection.NewUsbConnection(val)
			//if err != nil{
			//	fmt.Println(err)
			//	continue
			//}
			//defer conn.Close()
			//conn.Write([]byte("hello world\n\n"))
		}
	}

	//串口设备
	comPathList, err := GetComList()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("串口设备")
		for index, val := range comPathList {
			fmt.Printf("[%d] %s\n", index, val)
		}
	}

	//并口设备
	lptPathList, err := GetLPTList()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("并口设备")
		for index, val := range lptPathList {
			fmt.Printf("[%d] %s\n", index, val)
		}
	}

	//驱动设备
	prnPathList ,err := GetPRNList()
	if err != nil{
		fmt.Println(err)
		return
	}else{
		fmt.Println("驱动设备")
		for index, val := range prnPathList {
	 fmt.Printf("[%d] %s\n" ,index, val)
		}
	}

}
