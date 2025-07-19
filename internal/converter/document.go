package converter

import (
	"fmt"
	"path/filepath"

	"os/exec"
)

type DocumentConverter struct{}

func (c *DocumentConverter) Convert(inputFile, outputFile string) error {
	ext := filepath.Ext(inputFile)
	outputExt := filepath.Ext(outputFile)

	if ext == ".docx" && outputExt == ".pdf" {
		return c.convertDocxToPdf(inputFile, outputFile)
	}

	return fmt.Errorf("unsupported conversion: from %s to %s", ext, outputExt)
}

func (c *DocumentConverter) convertDocxToPdf(inputFile, outputFile string) error {
	cmd := exec.Command("pandoc", "-s", inputFile, "-o", outputFile)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("pandoc conversion failed: %v\nOutput: %s", err, output)
	}
	return nil
}
