package main
/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: -L . -lcgo_wrap

#include "cgo_wrap.h"

int callOnMeGo_cgo(int in); // Forward declaration.
int testCallback(callback_fcn f);
*/
import "C"

import (
	"fmt"
	"unsafe"
)

//export callOnMeGo
func callOnMeGo(in int) int {
	fmt.Printf("Go.callOnMeGo(): called with arg = %d\n", in)
	return in + 1
}

//export callOnMeGo2
func callOnMeGo2(f C.callback_fcn) int {
	fmt.Printf("Go.callOnMeGo2(): called with arg = %d\n", f)
	C.testCallback(f)
	return 1
}

func main() {
	fmt.Printf("Go.main(): calling C function with callback to us\n")
	C.some_c_func((C.callback_fcn)(unsafe.Pointer(C.callOnMeGo_cgo)))
}