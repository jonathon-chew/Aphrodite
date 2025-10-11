# :rainbow: Aphrodite (Go)
<p align="center">
<img width="400" src="doc/images/Aphrodite.png" alt="Aphrodite" title="Aphrodite" />
</p>


A package for beautifying terminal outputs in go

## 🚀 Features

- Formats output strings by passing in 2 extra paraemters 
- Supports Colour, Bold, Underline, Background, High Intensity, Bold High Intensity and High Intensity Background
- Serves the result directly out put to the terminal as a drop in 'replacement' for fmt.Printf 

## 🛠️ Prerequisites

- [Go] Go installed 

## 📁 Setup

1. In your go mod:

  ```bash
   go mod init
   go get -u github.com/jonathon-chew/aphrodite v1.0.2
   ```

  Should output in your go.mod file:

  ```bash
  require github.com/jonathon-chew/aphrodite v1.0.2 // indirect
  ```

2. In your go source file where required:

  ```go
   import github.com/jonathon-chew/aphrodite
  ```

  or 

  ```go
   import aphrodite "github.com/jonathon-chew/aphrodite"
  ```

  Depending if you want to name it something different when calling it.

## 📂 Output

Coloured/Formatted text in your terminal for convience

## 🧠 Notes

This is currently a work in progress with a few inovations planned for the future.
Issues will be tracked in Github issues.

## 📜 License

This project is licensed under the MIT License. See the LICENSE file for details.
### 🖌️ Attribution

The Go Gopher was originally designed by [Renee French](https://reneefrench.blogspot.com/).  
Used under the [Creative Commons Attribution 4.0 License](https://creativecommons.org/licenses/by/4.0/).  