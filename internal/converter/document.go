package converter

import (
	"fmt"
	"path/filepath"

	"os/exec"
)

type DocumentConverter struct {
	PDFEngine string
}

func (c *DocumentConverter) Convert(inputFile, outputFile string) error {
	ext := filepath.Ext(inputFile)
	outputExt := filepath.Ext(outputFile)

	if ext == ".docx" && outputExt == ".pdf" {
		return c.convertDocxToPdf(inputFile, outputFile, c.PDFEngine)
	}

	return fmt.Errorf("unsupported conversion: from %s to %s", ext, outputExt)
}

func (c *DocumentConverter) convertDocxToPdf(inputFile, outputFile, pdfEngine string) error {
	args := []string{"-s", inputFile, "-o", outputFile}
	if pdfEngine != "" && pdfEngine != "default" {
		args = append(args, "--pdf-engine", pdfEngine)
	}
	cmd := exec.Command("pandoc", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("pandoc conversion failed: %v\nOutput: %s", err, output)
	}
	return nil
}
