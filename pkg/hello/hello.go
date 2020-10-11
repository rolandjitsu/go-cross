package hello

// #include <stdlib.h>
// #include <hello.h>
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/vmihailenco/msgpack/v5"
)

func Greet(name string) {
	cs := C.CString(name)
	size := C.size_t(len(name))
	msg := C.greet(cs, size)
	fmt.Println(C.GoString(msg))
	C.free(unsafe.Pointer(cs))
}

func GreetMsgPack(name string) (err error) {
	cs := C.CString(name)
	size := C.size_t(len(name))
	buf := C.msgpack_sbuffer{}

	C.msgpack_sbuffer_init(&buf)

	C.greet_msgpack(cs, size, &buf)

	data := C.GoBytes(unsafe.Pointer(buf.data), C.int(buf.size))

	var msg string
	err = msgpack.Unmarshal(data, &msg)
	if err != nil {
		return
	}
	fmt.Println(msg)

	C.msgpack_sbuffer_destroy(&buf)
	C.free(unsafe.Pointer(cs))

	return
}
