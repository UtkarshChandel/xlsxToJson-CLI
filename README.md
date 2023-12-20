
# xlsxToJson-CLI

## About
`xlsxToJson-CLI` is a straightforward and efficient command-line tool that converts Excel (.xlsx) files into JSON format. It's designed to be user-friendly, with interactive prompts for file paths and built-in checks for file existence and format. This tool is perfect for anyone needing to quickly transform spreadsheet data into the widely-used JSON format.

## Features
- **Easy-to-use CLI**: Simple prompts guide the user through the conversion process.
- **File Validation**: Automatically checks if the provided Excel file exists and is in the correct format (.xlsx).
- **Efficient Conversion**: Quickly converts Excel data to structured JSON.

## Installation
To install `xlsxToJson-CLI`, you need to have Go installed on your machine. Follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/[YourUsername]/xlsxToJson-CLI.git
   ```
2. Navigate to the project directory:
   ```bash
   cd xlsxToJson-CLI
   ```

## Usage
After installation, run the program using the Go command:

```bash
go run main.go
```

Follow the interactive prompts to specify the input Excel file and the output JSON file.

## Building from Source
To build an executable from the source:

1. Navigate to the project directory.
2. Run the build command:
   ```bash
   go build -o xlsxToJson-CLI.exe
   ```
3. The executable `xlsxToJson-CLI.exe` can now be used directly.

## Contributing
Contributions to `xlsxToJson-CLI` are welcome! Feel free to open issues or submit pull requests.

## Acknowledgements
Special thanks to all the contributors and users of the `excelize` Go package, which made this project possible.
