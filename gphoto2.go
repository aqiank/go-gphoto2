package gp

// #cgo linux pkg-config: libgphoto2
// #include <gphoto2/gphoto2.h>
// #include <string.h>
import "C"
import "unsafe"
import "errors"

const (
	CAPTURE_IMAGE = C.GP_CAPTURE_IMAGE
	CAPTURE_MOVIE = C.GP_CAPTURE_MOVIE
	CAPTURE_SOUND = C.GP_CAPTURE_SOUND
)

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

type Camera C.Camera

func NewCamera() (*Camera, error) {
	var _cam *C.Camera

	if ret := C.gp_camera_new(&_cam); ret != 0 {
		return nil, e(ret)
	}

	return (*Camera)(_cam), nil
}

func (cam *Camera) Init(ctx *Context) error {
	if ret := C.gp_camera_init(cam.c(), ctx.c()); ret != 0 {
		return e(ret)
	}

	return nil
}

func (cam *Camera) Capture(mode C.CameraCaptureType, path CameraFilePath, ctx *Context) error {
	if ret := C.gp_camera_capture(cam.c(), mode, path.c(), ctx.c()); ret != 0 {
		return e(ret)
	}

	return nil
}

func (cam *Camera) Capture2(mode C.CameraCaptureType, name, folder string, ctx *Context) error {
	var path CameraFilePath
	path.SetPath(name, folder)
	return cam.Capture(mode, path, ctx)

}

func (cam *Camera) Free() error {
	if ret := C.gp_camera_free(cam.c()); ret != 0 {
		return e(ret)
	}
	return nil
}

func (cam *Camera) c() *C.Camera {
	return (*C.Camera)(cam)
}

func e(ret C.int) error {
	return errors.New(C.GoString(C.gp_result_as_string(ret)))
}

type CameraFilePath C.CameraFilePath

func (path *CameraFilePath) SetPath(name, folder string) {
	_path := (*C.CameraFilePath) (path)
	C.strcpy(&_path.name[0], C.CString(name))
	C.strcpy(&_path.folder[0], C.CString(folder))
}

func (path *CameraFilePath) c() (*C.CameraFilePath) {
	return (*C.CameraFilePath)(path)
}
