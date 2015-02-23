package pkg

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L./my -lmy
#include "my/lib.h"
*/
import "C"
import "fmt"

func Print() {
	C.ObjectInit()
	fmt.Printf("Global: %d\n", C.Get(nil))
	fmt.Printf("Global: %d\n", C.Get(C.Global))
}
