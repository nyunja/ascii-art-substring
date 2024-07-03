# ASCII Banner Highlighter

This program reads an ASCII banner file and highlights specific sequences within the provided text. The sequences to be highlighted are specified by the user. The highlighted text is then printed to the console with the specified color.

## Features

- Reads an ASCII banner file to generate ASCII art.
- Highlights specified sequences within the text.
- Outputs the highlighted ASCII art to the console.

## Usage

### Prerequisites

- Go 1.16+ installed on your machine.

### Instructions

1. Place your ASCII banner file (e.g., `standard.txt`) in the same directory as the Go program.
2. Run the Go program with the command:

    ```sh
    go run main.go
    ```

### Example

Given the input text `hello\nthere\nthere` and the sequence to highlight `re\nthe`, the program will output the ASCII art representation of the text with the specified sequences highlighted.

## Functions

### `readBannerFile(bannerFile string) ([]string, error)`

Reads the ASCII banner file and returns its lines as a slice of strings.

### `getOutputSlice(s string) []Output`

Converts the input string into a slice of `Output` structs, separating words and newline characters.

### `indicesToColor(s1, s2 string) map[int]bool`

Determines the indices of characters in `s1` that should be highlighted based on the sequences in `s2`.

### `getIndicesToColor(i, j int, indices map[int]bool) map[int]bool`

Helper function to mark the indices that should be highlighted.

### `getFinalOutput(s1, s2 string, lines []string) string`

Generates the final output string with highlighted sequences based on the input text and ASCII banner lines.

## Data Structures

### `Chars`

A struct representing a character and its position in the string.

```go
type Chars struct {
    index     int
    character rune
}
```

### `Output`

A struct representing a line of output, including newlines and words.

```go
type Output struct {
    newLine string
    word    []Chars
}
```

## Example Output

```
hello
there
there
```

Highlighted sequences (e.g., `re`, `the`) will appear in green.

## License

This project is licensed under the MIT License.

## Author

Nyunja

---

Feel free to customize this `README.md` file according to your specific needs.