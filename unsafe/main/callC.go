package main

// #cgo CFLAGS: -I${SRCDIR}/library
// #cgo LDFLAGS: ${SRCDIR}/callC.a
// #include <stdlib.h>
// #include <callC.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("going to call a C function!")

	C.cHello()

	fmt.Println("going to call another C function!")

	myMessage := C.CString("This is a string!")
	defer C.free(unsafe.Pointer(myMessage))

	C.printMessage(myMessage)

	fmt.Println("all perfectly done!")
}
