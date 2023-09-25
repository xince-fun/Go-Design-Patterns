package flyweight

import "fmt"

type ImageFlyweightFactory struct {
	maps map[string]*ImageFlyweight
}

func GetImageFlyweightFactory() *ImageFlyweightFactory {
	if imageFactory == nil {
		imageFactory = &ImageFlyweightFactory{
			maps: make(map[string]*ImageFlyweight, 0),
		}
	}
	return imageFactory
}

func (t *ImageFlyweightFactory) Get(filename string) *ImageFlyweight {
	image := t.maps[filename]
	if image == nil {
		image = NewImageFlyweight(filename)
		t.maps[filename] = image
	}

	return image
}

var imageFactory *ImageFlyweightFactory

type ImageFlyweight struct {
	data string
}

func NewImageFlyweight(filename string) *ImageFlyweight {
	// Load image file
	data := fmt.Sprintf("image data %s", filename)
	return &ImageFlyweight{
		data: data,
	}
}

func (i *ImageFlyweight) Data() string {
	return i.data
}

type ImageViewer struct {
	*ImageFlyweight
}

func NewImageViewer(filename string) *ImageViewer {
	image := GetImageFlyweightFactory().Get(filename)
	return &ImageViewer{
		ImageFlyweight: image,
	}
}

func (i *ImageViewer) Display() {
	fmt.Printf("Display: %s\n", i.Data())
}
