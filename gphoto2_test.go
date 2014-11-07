package gp

import (
	"testing"
)

func TestGphoto2t (t *testing.T) {
	context := NewContext()
	
	camera, err := NewCamera()
	if err != nil {
		t.Fatal(err)
	}

	err = camera.Init(context)
	if err != nil {
		t.Fatal(err)
	}

	path, err := camera.Capture(CAPTURE_IMAGE, context)
	if err != nil {
		t.Fatal(err)
	}

	file, err := camera.File(path.Folder, path.Name, FILE_TYPE_NORMAL, context)
	if err != nil {
		t.Fatal(err)
	}

	err = file.Save("image.jpg")
	if err != nil {
		t.Fatal(err)
	}

	err = camera.Free()
	if err != nil {
		t.Fatal(err)
	}

	context.Free()
}
