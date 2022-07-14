package flyweight

import (
	"fmt"
	"reflect"
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
	fmt.Println(viewer1)
	viewer1.Display()
	viewer2 := NewImageViewer("image1.png")
	fmt.Println(viewer2)
	viewer2.Display()
	if viewer1 != viewer2 {
		t.Fail()
	}

	fmt.Println(reflect.DeepEqual(viewer1, viewer2))
	fmt.Println(reflect.DeepEqual(viewer1.data, viewer2.data))
	fmt.Println(reflect.DeepEqual(viewer1.ImageFlyweight, viewer2.ImageFlyweight))
}