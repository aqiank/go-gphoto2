package gp

// #cgo linux pkg-config: libgphoto2
// #include <gphoto2/gphoto2.h>
// #include <string.h>
import "C"
import "unsafe"

type Context C.GPContext

func NewContext() *Context {
	return (*Context)(unsafe.Pointer(C.gp_context_new()))
}

func (ctx *Context) Free() {
	C.gp_context_unref(ctx.c())
}

func (ctx *Context) c() *C.GPContext {
	return (*C.GPContext)(ctx)
}