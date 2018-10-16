package main

/*
#cgo CFLAGS: -I./resource_c
#cgo LDFLAGS: -L./resource_c -ltest
#include<stdio.h>
#include<stdlib.h>
#include "./resource_c/test.h"
#cgo windows CFLAGS: -DCGO_OS_WINDOWS=1
#cgo darwin CFLAGS: -DCGO_OS_DARWIN=1
#cgo linux CFLAGS: -DCGO_OS_LINUX=1

#ifdef CGO_OS_WINDOWS
	extern char* os="windows";
#elif CGO_OS_DARWIN
	extern char* os="darwin";
#elif CGO_OS_LINUX
	extern char* os="linux";
#else
	#error(unknown os)
#endif
*/
import "C"
import "unsafe"

func main() {
	str := C.CString("songweishuai")
	defer C.free(unsafe.Pointer(str))
	//C.prints(C.CString("Hello,nihao!"))
	C.prints(str)

	println(C.GoString(C.os))
}
