# KKonverter

A program i made cause i didn't want to go to random websites just to convert files (this is basically a ffmpeg & pandoc wrapper)

## Features

- **Document Conversion**: Convert `.docx` to `.pdf` using `pandoc`.
- **Media Conversion**: Convert media files using `ffmpeg-go`.
- **Batch Conversion**: Convert multiple files at once.
- **Drag and Drop**: Easily add files by dragging and dropping them onto the application window.
- **Custom Output Directory**: Select a custom directory for converted files.
- **Conversion Presets**: Choose from different quality presets for conversions.
- **Progress Display**: Monitor the conversion progress with a progress bar.

## Building and Running

### Prerequisites

- Go (version 1.16 or higher)
- Fyne dependencies (refer to [Fyne documentation](https://developer.fyne.io/started/)). For most systems, this involves installing development headers for OpenGL and X11.
- FFmpeg (for media conversion). Ensure `ffmpeg` is installed and accessible in your system's PATH.
- Pandoc (for document conversion). Ensure `pandoc` is installed and accessible in your system's PATH. You can download it from [https://pandoc.org/installing.html](https://pandoc.org/installing.html).

### Steps

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/your-repo/KKonverter.git
    cd KKonverter
    ```

2.  **Install Go modules:**
    ```bash
    go mod tidy
    ```

3.  **Run the application:**
    ```bash
    go run cmd/kkonverter/main.go
    ```

4.  **Build the application (optional):**
    ```bash
    go build -o kkonverter cmd/kkonverter/main.go
    ```
    This will create an executable named `kkonverter` (or `kkonverter.exe` on Windows) in the current directory.

## Usage

1.  **Add Files**: Click the "Add Files" button to select files, or drag and drop files directly onto the application window.
2.  **Select Output Format**: Choose the desired output format from the dropdown menu.
3.  **Select Output Directory**: Click "Browse..." to choose a custom output directory. By default, files will be converted to the current directory.
4.  **Select Preset**: Choose a conversion preset (e.g., "High Quality", "Medium Quality", "Low Quality") from the dropdown.
5.  **Convert**: Click the "Convert All" button to start the conversion process. A progress bar will indicate the status.

## Project Structure

```
.
├── cmd/
│   └── kkonverter/
│       └── main.go           # Application entry point
├── internal/
│   ├── app/
│   │   └── app.go            # Core application logic
│   ├── converter/
│   │   ├── converter.go      # Converter interface definition
│   │   ├── document.go       # Document conversion implementation
│   │   └── media.go          # Media conversion implementation
│   └── ui/
│       └── layouts.go        # Fyne GUI layout and logic
└── go.mod                    # Go module file
└── go.sum                    # Go module checksums
└── README.md                 # Project README
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
