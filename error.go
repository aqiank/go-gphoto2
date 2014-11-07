package gp

// #cgo linux pkg-config: libgphoto2
// #include <gphoto2/gphoto2.h>
// #include <string.h>
import "C"
import "errors"

func e(ret C.int) error {
	return errors.New(C.GoString(C.gp_result_as_string(ret)))
}
