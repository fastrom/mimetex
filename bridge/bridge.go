// Package bridge implements a call to mimetex C API
package bridge

// #cgo CFLAGS: -DAA
// #cgo LDFLAGS: -lm
// #include "stdlib.h"
// #include "mimetex.h"
//
// unsigned char *
// allocate_buf(int bytes) {
//   return	(unsigned char *) malloc(bytes);
// }
import "C"
import "unsafe"
import "errors"

const MAXGIFSZ = 131072

var (
	EmptyStrErr      = errors.New("empty expression to render")
	IncorrectSizeErr = errors.New("size should be in range 1-12")
)

func RenderExprToGif(expr string, size int) ([]byte, error) {
	if len(expr) < 1 {
		return nil, EmptyStrErr
	} else if size < 1 || size > 12 {
		return nil, IncorrectSizeErr
	}

	outgif := C.allocate_buf(MAXGIFSZ)
	defer C.free(unsafe.Pointer(outgif))
	sexpr := C.CString(expr)
	defer C.free(unsafe.Pointer(sexpr))

	nbytes := C.RenderExprToGif(sexpr, C.int(size), outgif)
	data := C.GoBytes(unsafe.Pointer(outgif), nbytes)
	return data, nil
}
