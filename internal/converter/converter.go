package converter

type Converter interface {
	Convert(inputFile, outputFile string) error
}