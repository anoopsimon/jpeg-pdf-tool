
# JPEG to PDF Converter

This tool converts multiple JPEG images into a single PDF file using Go. You can provide images by specifying a folder or a comma-separated list of filenames.

## Features

- Converts JPEG images to a PDF document
- Supports specifying either a folder containing images or a comma-separated list of files
- Automatically sorts images alphabetically to maintain order in the final PDF
- Customizable output PDF filename

## Prerequisites

- Go 1.16+ installed on your machine
- JPEG images to be converted

## Installation

Clone the repository, navigate to the project directory, and build the program:
```bash
git clone https://github.com/anoopsimon/image-to-pdf-go.git
cd pdf-tool
go mod tidy
go build -o jpeg-to-pdf
```

## Usage

Run the program using command-line flags:

```bash
jpeg-pdf-tool -folder=<folder_path> -output=<output_pdf_filename>
```
or
```bash
jpeg-pdf-tool -files=<file1.jpg,file2.jpg,...> -output=<output_pdf_filename>
```

### Command-Line Options

| Flag        | Description                                                                  |
|-------------|------------------------------------------------------------------------------|
| `-folder`   | Path to the folder containing JPEG files (optional if `-files` is used).    |
| `-files`    | Comma-separated list of JPEG filenames (optional if `-folder` is used).      |
| `-output`   | Name of the output PDF file (default: `images_output.pdf`).                 |

**Note**: Either `-folder` or `-files` must be provided. If both are provided, the `-folder` option will be used.

### Examples

1. **Convert all JPEG images in a folder to a PDF**:
   ```bash
   jpeg-pdf-tool -folder=imgs -output=output.pdf
   ```

2. **Convert specified JPEG files to a PDF**:
   ```bash
  jpeg-pdf-tool -files="imgs/image1.jpg,imgs/image2.jpg" -output=output.pdf
   ```

3. **Convert with default output filename**:
   ```bash
   jpeg-pdf-tool -folder=imgs
   ```

## Error Handling

If any issues occur, such as missing files or unsupported image formats, an error message will display in the terminal.
