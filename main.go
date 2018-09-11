package main

/*
#cgo CFLAGS: -I./resource_c
#cgo LDFLAGS: -L./resource_c -ltest
#include "./resource_c/test.h"
*/
import "C"

func main() {
	C.prints(C.CString("Hello,nihao!"))
}
