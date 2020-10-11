package hello

// #include <stdlib.h>
// #include <hello.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func Greet(name string) {
	cs := C.CString(name)
	size := C.size_t(len(name))
	msg := C.greet(cs, size)
	fmt.Println(C.GoString(msg))
	C.free(unsafe.Pointer(cs))
}
