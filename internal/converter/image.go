package converter

import (
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

type ImageConverter struct{}

func (c *ImageConverter) Convert(inputFile, outputFile string) error {
	inputExt := strings.ToLower(getFileExtension(inputFile))
	outputExt := strings.ToLower(getFileExtension(outputFile))

	if inputExt == "jpg" || inputExt == "jpeg" {
		if outputExt == "png" {
			return c.jpegToPng(inputFile, outputFile)
		}
	} else if inputExt == "png" {
		if outputExt == "jpg" || outputExt == "jpeg" {
			return c.pngToJpeg(inputFile, outputFile)
		}
	}
	return nil // Or return an error for unsupported conversion
}

func (c *ImageConverter) jpegToPng(inputFile, outputFile string) error {
	reader, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer reader.Close()

	img, err := jpeg.Decode(reader)
	if err != nil {
		return err
	}

	outputFilePNG, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outputFilePNG.Close()

	return png.Encode(outputFilePNG, img)
}

func (c *ImageConverter) pngToJpeg(inputFile, outputFile string) error {
	reader, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer reader.Close()

	img, err := png.Decode(reader)
	if err != nil {
		return err
	}

	outputFileJPG, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outputFileJPG.Close()

	return jpeg.Encode(outputFileJPG, img, nil)
}

func getFileExtension(filename string) string {
	parts := strings.Split(filename, ".")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}
	return ""
}