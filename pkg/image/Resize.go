package image

import "github.com/h2non/bimg"

func Resize(buffer []byte, options bimg.Options) ([]byte, error) {
	newImage, err := bimg.NewImage(buffer).Resize(800, 600)

	if err != nil {
		return nil, err
	}

	return newImage, nil
}
