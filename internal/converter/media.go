package converter

import ffmpeg "github.com/u2takey/ffmpeg-go"

type MediaConverter struct{}

func (c *MediaConverter) Convert(inputFile, outputFile string) error {
	return ffmpeg.Input(inputFile).
		Output(outputFile).
		OverWriteOutput().
		Run()
}