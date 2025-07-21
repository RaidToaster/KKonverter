package converter

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

type ImageConverter struct{}

func (c *ImageConverter) Convert(inputFile, outputFile string) error {
	reader, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		return err
	}

	writer, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer writer.Close()

	outputExt := strings.ToLower(filepath.Ext(outputFile))
	switch outputExt {
	case ".jpg", ".jpeg":
		return jpeg.Encode(writer, img, nil)
	case ".png":
		return png.Encode(writer, img)
	case ".gif":
		return gif.Encode(writer, img, nil)
	default:
		return fmt.Errorf("unsupported output image format: %s", outputExt)
	}
}
