package multicomparer

import (
	"image"
	_ "image/jpeg"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetChannelValue_Red(t *testing.T) {
	fImg1, _ := os.Open("../../../../../../mic-example-images/images1.jpeg")
	defer fImg1.Close()
	img1, _, _ := image.Decode(fImg1)
	assert.EqualValues(t, 6376, GetChannelValue(img1, 0, 0, R), "Wrong Red value found for image")
}

func TestGetChannelValue_Green(t *testing.T) {
	fImg1, _ := os.Open("../../../../../../mic-example-images/images1.jpeg")
	defer fImg1.Close()
	img1, _, _ := image.Decode(fImg1)
	assert.EqualValues(t, 34047, GetChannelValue(img1, 0, 0, G), "Wrong Red value found for image")
}

func TestGetChannelValue_Blue(t *testing.T) {
	fImg1, _ := os.Open("../../../../../../mic-example-images/images1.jpeg")
	defer fImg1.Close()
	img1, _, _ := image.Decode(fImg1)
	assert.EqualValues(t, 55941, GetChannelValue(img1, 0, 0, B), "Wrong Red value found for image")
}

func TestGetChannelValue_Alpha(t *testing.T) {
	fImg1, _ := os.Open("../../../../../../mic-example-images/images1.jpeg")
	defer fImg1.Close()
	img1, _, _ := image.Decode(fImg1)
	assert.EqualValues(t, 65535, GetChannelValue(img1, 0, 0, A), "Wrong Red value found for image")
}
