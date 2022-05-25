package flyweight

import (
	"fmt"
	"testing"
)

func ExampleFlyweight() {
	viewer := NewImageViewer("image1.png")
	viewer.Display()
	viewer.Data()
	// Output:
	// Display: image data image1.png
}

func TestFlyweight(t *testing.T) {
	viewer1 := NewImageViewer("image1.png")
	viewer2 := NewImageViewer("image1.png")
	fmt.Println(viewer2)
	viewer2.Display()
	if viewer1.ImageFlyweight != viewer2.ImageFlyweight {
		t.Fail()
	}
}
