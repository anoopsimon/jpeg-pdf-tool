package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/signintech/gopdf"
)

// getImageDimensions retrieves the dimensions of a JPEG image
func getImageDimensions(filename string) (float64, float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		return 0, 0, err
	}

	bounds := img.Bounds()
	width := float64(bounds.Dx())
	height := float64(bounds.Dy())
	return width, height, nil
}

// jpegToPDF converts JPEG images to a PDF file
func jpegToPDF(jpegFiles []string, outputFile string) error {
	pdf := gopdf.GoPdf{}
	// Start the PDF with A4 page size
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	for _, jpegFile := range jpegFiles {
		// Get dimensions of the image
		imgWidth, imgHeight, err := getImageDimensions(jpegFile)
		if err != nil {
			return fmt.Errorf("failed to get dimensions of image %s: %v", jpegFile, err)
		}

		// Calculate scaling factors to fit the image in the A4 page size
		pageWidth, pageHeight := gopdf.PageSizeA4.W, gopdf.PageSizeA4.H
		scale := 1.0
		if imgWidth > pageWidth || imgHeight > pageHeight {
			widthRatio := pageWidth / imgWidth
			heightRatio := pageHeight / imgHeight
			scale = min(widthRatio, heightRatio)
		}

		// Calculate the new dimensions of the image
		newWidth := imgWidth * scale
		newHeight := imgHeight * scale

		// Center the image on the page
		xOffset := (pageWidth - newWidth) / 2
		yOffset := (pageHeight - newHeight) / 2

		// Add a new page for each image and insert the image
		pdf.AddPage()
		err = pdf.Image(jpegFile, xOffset, yOffset, &gopdf.Rect{W: newWidth, H: newHeight})
		if err != nil {
			return fmt.Errorf("failed to add image %s to PDF: %v", jpegFile, err)
		}
	}

	err := pdf.WritePdf(outputFile) // Write the PDF to the specified file
	if err != nil {
		return fmt.Errorf("failed to write PDF to file: %v", err)
	}

	fmt.Println("JPEGs converted to PDF successfully:", outputFile)
	return nil
}

// min returns the smaller of two float64 values
func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// getJPEGFiles retrieves all JPEG files from the specified folder
func getJPEGFiles(folder string) ([]string, error) {
	var jpegFiles []string
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Check if the file is a JPEG
		if !info.IsDir() && (filepath.Ext(path) == ".jpg" || filepath.Ext(path) == ".jpeg") {
			jpegFiles = append(jpegFiles, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Strings(jpegFiles) // Sort the JPEG files alphabetically
	return jpegFiles, nil
}

func main() {
	// Command-line flags
	folderPtr := flag.String("folder", "", "Path to the folder containing JPEG files")
	filesPtr := flag.String("files", "", "Comma-separated list of JPEG filenames")
	outputFile := flag.String("output", "merged_pdf.pdf", "Output PDF filename")

	flag.Parse()

	var jpegFiles []string
	var err error

	// Determine if we should use folder or list of files
	if *folderPtr != "" {
		jpegFiles, err = getJPEGFiles(*folderPtr)
		if err != nil {
			log.Fatal("Error retrieving JPEG files:", err)
		}
	} else if *filesPtr != "" {
		// Split the comma-separated list into an array
		jpegFiles = strings.Split(*filesPtr, ",")
		for i := range jpegFiles {
			jpegFiles[i] = strings.TrimSpace(jpegFiles[i]) // Trim whitespace
		}
		sort.Strings(jpegFiles) // Sort the files alphabetically
	} else {
		log.Fatal("Please provide either a folder path or a list of JPEG files.")
	}

	// Convert JPEG to PDF
	err = jpegToPDF(jpegFiles, *outputFile)
	if err != nil {
		log.Println("Error converting JPEGs to PDF:", err)
	}
}
