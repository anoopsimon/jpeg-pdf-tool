package main

import (
	"testing"
)

func TestJpegToPDF(t *testing.T) {
	// Example call with test file(s) for basic verification
	jpegFiles := []string{"testdata/sample.jpg"} // Ensure a sample test file exists
	outputFile := "test_output.pdf"
	err := jpegToPDF(jpegFiles, outputFile)

	if err != nil {
		t.Fatalf("Failed to convert JPEG to PDF: %v", err)
	}
}
