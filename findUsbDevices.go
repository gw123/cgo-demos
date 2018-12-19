package main

import (
	"syscall"
	"unsafe"
	"log"
	"fmt"
)

func main() {
	calltest3()
}

type DWORD uintptr

type ULONG_PTR uintptr

type SP_DEVICE_INTERFACE_DATA struct {
	cbSize             uintptr
	InterfaceClassGuid syscall.GUID
	Flags              uintptr
	Reserved           ULONG_PTR
}

type SP_DEVICE_INTERFACE_DETAIL_DATA struct {
	cbSize     uintptr;
	DevicePath [256]uint16;
};

type SP_DEVINFO_DATA struct {
	cbSize    uintptr
	ClassGuid syscall.GUID
	DevInst   uintptr; // DEVINST handle
	Reserved  uintptr;
}

func calltest4() {
	h := syscall.MustLoadDLL("libusb2.dll")
	//Hello := h.MustFindProc("AutomaticmonInit")

	Hello := h.MustFindProc("add")
	ret, _, err := Hello.Call()
	fmt.Println(ret, err)
}
func calltest3() {

	const (
		DIGCF_DEFAULT         = 0x00000001
		DIGCF_PRESENT         = 0x00000002
		DIGCF_ALLCLASSES      = 0x00000004
		DIGCF_PROFILE         = 0x00000008
		DIGCF_DEVICEINTERFACE = 0x00000010
	)

	h := syscall.MustLoadDLL("SetupAPI.dll")
	SetupDiGetClassDevs := h.MustFindProc("SetupDiGetClassDevsW")

	hDevInfo, _, err := SetupDiGetClassDevs.Call(uintptr(0), uintptr(0), uintptr(0), uintptr(DIGCF_ALLCLASSES|DIGCF_DEVICEINTERFACE))
	if hDevInfo == 0 {
		fmt.Println(err)
	} else {
		fmt.Println("one ", err, hDevInfo)
	}

	SetupDiEnumDeviceInterfaces := h.MustFindProc("SetupDiEnumDeviceInterfaces")
	icount := 0
	var USB_GUID = syscall.GUID{
		Data1: 0xA5DCBF10,
		Data2: 0x6530,
		Data3: 0x11D2,
		Data4: [8]byte{0x90, 0x1F, 0x00, 0xC0, 0x4F, 0xB9, 0x51, 0xED},
	}
	//{ 0xA5DCBF10L, 0x6530, 0x11D2, { 0x90, 0x1F, 0x00, 0xC0, 0x4F, 0xB9, 0x51, 0xED } };
	//uintptr(unsafe.Pointer(lpbi))
	//Data4 :=[80]byte{}
	deviceInterfaceData := new(SP_DEVICE_INTERFACE_DATA)
	deviceInterfaceData.cbSize = unsafe.Sizeof(*deviceInterfaceData)
	fmt.Println(unsafe.Sizeof(*deviceInterfaceData))
	//buffer := make([]byte ,10028)
	ret, _, err := SetupDiEnumDeviceInterfaces.Call(
		uintptr(hDevInfo),
		uintptr(0),
		uintptr(unsafe.Pointer(&USB_GUID)),
		uintptr(icount),
		uintptr(unsafe.Pointer(deviceInterfaceData)))
	fmt.Println("86 ",ret, err, deviceInterfaceData)

	SetupDiGetInterfaceDeviceDetail := h.MustFindProc("SetupDiGetDeviceInterfaceDetailW")
	var requiredLength uint64 = 2000
	ret, _, err = SetupDiGetInterfaceDeviceDetail.Call(uintptr(hDevInfo),
		uintptr(unsafe.Pointer(deviceInterfaceData)),
		0, 0,
		uintptr(unsafe.Pointer(&requiredLength)), 0);
	fmt.Println("Line 94 ",ret, err, requiredLength)

	deviceInterfaceDetailData := make([]byte ,1024)   // new(SP_DEVICE_INTERFACE_DETAIL_DATA)
	deviceInfoData := new(SP_DEVINFO_DATA)
	deviceInfoData.cbSize = unsafe.Sizeof(*deviceInfoData)
	ret, _, err = SetupDiGetInterfaceDeviceDetail.Call(uintptr(hDevInfo),
		uintptr(unsafe.Pointer(deviceInterfaceData)),
		0,
		0,
		0,
		0);
	fmt.Println(ret, err, requiredLength, deviceInterfaceDetailData)

	//SetupDiEnumDeviceInterfaces.Call( hDevInfo ,0 , 0 ,uintptr(icount) , )
	//icount := 0
	//for  {
	//
	//}
}

func calltest1() {
	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("CreateFileW")

	//lpTotalNumberOfBytes := 0x00000000

	r2, _, err := c.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("D:\\tmp.data"))),
		uintptr(0x40000000),
		uintptr(0x00000000),
		uintptr(0),
		uintptr(1),
		uintptr(0x00000080),
		uintptr(0),
	)
	if r2 != 0 {
		log.Println(r2, err)
	} else {
		log.Println(r2, err)
	}
}
func calltest() {
	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")
	lpFreeBytesAvailable := int64(0)
	lpTotalNumberOfBytes := int64(0)
	lpTotalNumberOfFreeBytes := int64(0)
	r2, _, err := c.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("C:"))),
		uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)))
	if r2 != 0 {
		log.Println(r2, err, lpFreeBytesAvailable/1024/1024, "MB")
	} else {
		log.Println(r2, err)
	}
}

func calltest2() {
	//首先,准备输入参数, GetDiskFreeSpaceEx需要4个参数, 可查MSDN
	dir := "D:"
	lpFreeBytesAvailable := int64(0) //注意类型需要跟API的类型相符 ,类型失败有时候会编译成功但是会发生内容不全的问题
	lpTotalNumberOfBytes := int64(0)
	lpTotalNumberOfFreeBytes := int64(0)

	//获取方法的引用
	kernel32, err := syscall.LoadLibrary("Kernel32.dll") // 严格来说需要加上
	defer syscall.FreeLibrary(kernel32)
	if err != nil {
		fmt.Println(err)
	}
	GetDiskFreeSpaceEx, err := syscall.GetProcAddress(syscall.Handle(kernel32), "GetDiskFreeSpaceExW")

	//执行之. 因为有4个参数,故取Syscall6才能放得下. 最后2个参数,自然就是0了
	r, _, _ := syscall.Syscall6(uintptr(GetDiskFreeSpaceEx), 4,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(dir))),
		uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)), 0, 0)

	// 注意, errno并非error接口的, 不可能是nil
	// 而且,根据MSDN的说明,返回值为0就fail, 不为0就是成功
	if r != 0 {
		log.Printf("Free %dmb ,Free %dmb ,Free %dmb ",
			lpFreeBytesAvailable/1024/1024,
			lpTotalNumberOfBytes/1024/1024,
			lpTotalNumberOfFreeBytes/1024/1024)
	}
}
