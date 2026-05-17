# :rainbow: Aphrodite (Go)
<p align="center">
<img width="400" src="doc/images/Aphrodite.png" alt="Aphrodite" title="Aphrodite" />
</p>

A package for beautifying terminal output in Go.

## 🚀 Features

- Formats output strings by passing in two extra parameters
- Supports colour, bold, underline, background, high intensity, bold high intensity, and high intensity background
- Writes formatted output directly to the terminal as a drop-in replacement for `fmt.Printf`

## 🛠️ Prerequisites

- [Go](https://golang.org/dl/) installed

## 📁 Setup

1. In your Go module:

  ```bash
   go get github.com/jonathon-chew/aphrodite@latest
   ```

   Your `go.mod` file should then include something like this:

  ```bash
  require github.com/jonathon-chew/aphrodite v1.5.2 // indirect
  ```

2. Import it in your Go source file where needed:

  ```go
   import "github.com/jonathon-chew/aphrodite"
  ```

   Or alias it:

  ```go
   import aphrodite "github.com/jonathon-chew/aphrodite"
  ```

   Use whichever style fits your codebase.

## 📂 Output

Coloured, formatted text in your terminal.

## 🧠 Notes

Issues will be tracked in GitHub Issues.

## 📜 License

This project is licensed under the MIT License. See the LICENSE file for details.

### 🖌️ Attribution

The Go Gopher was originally designed by [Renee French](https://reneefrench.blogspot.com/).  
Used under the [Creative Commons Attribution 4.0 License](https://creativecommons.org/licenses/by/4.0/).  
