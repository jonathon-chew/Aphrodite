# 📸 Aphrodite (Go)

A package for colour terminal outputs in go

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
