package main

/*
#cgo LDFLAGS: -Lmylib/target/debug -lmylib
#include "./mylib.h"
*/
import "C"

func main() {
	C.hello(C.CString("World"))
}
