package main

/*
#cgo LDFLAGS: -Lmylib/target/release -lmylib
#include "./mylib.h"
*/
import "C"

func main() {
	C.hello(C.CString("World"))
}
