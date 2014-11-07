package gp

// #cgo linux pkg-config: libgphoto2
// #include <gphoto2/gphoto2.h>
// #include <string.h>
import "C"
import "unsafe"

const (
	FILE_TYPE_PREVIEW = C.GP_FILE_TYPE_PREVIEW
	FILE_TYPE_NORMAL = C.GP_FILE_TYPE_NORMAL
	FILE_TYPE_RAW = C.GP_FILE_TYPE_RAW
	FILE_TYPE_AUDIO = C.GP_FILE_TYPE_AUDIO
	FILE_TYPE_EXIF = C.GP_FILE_TYPE_EXIF
	FILE_TYPE_METADATA = C.GP_FILE_TYPE_METADATA
)

type CameraFile C.CameraFile
type CameraFileType int

type CameraFilePath struct {
	Name string
	Folder string
}

func (file *CameraFile) Save(name string) error {
	_file := (*C.CameraFile)(unsafe.Pointer(file))
	_name := C.CString(name)
	if ret := C.gp_file_save(_file, _name); ret != 0 {
		return e(ret)
	}
	return nil
}