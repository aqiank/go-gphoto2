package gp

import (
	"testing"
)

func TestGphoto2t (t *testing.T) {
	ctx := NewContext()
	cam, err := NewCamera()
	if err != nil {
		t.Fatal(err)
	}

	err = cam.Init(ctx)
	if err != nil {
		t.Fatal(err)
	}

	err = cam.Capture2(CAPTURE_IMAGE, "image.jpg", ".", ctx)
	if err != nil {
		t.Fatal(err)
	}

	err = cam.Free()
	if err != nil {
		t.Fatal(err)
	}

	ctx.Free()
}
