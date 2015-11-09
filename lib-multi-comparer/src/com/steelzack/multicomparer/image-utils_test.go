package multicomparer

import (
	"image"
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetChannelValue(t *testing.T) {
	fImg1, _ := os.Open("../../../../../mic-example-images/images1.jpeg")
	defer fImg1.Close()
	img1, _, _ := image.Decode(fImg1)
	assert.EqualValues(t, 6376, GetChannelValue(img1, 0, 0, R), "blah")	
}