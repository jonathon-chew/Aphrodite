package aphrodite 

import "fmt"

// Colour prints the message in the specified color.
func Colour(option, message string) {
    colour := map[string]string{
        "Black":  "\x1b[0;30m",
        "Red":    "\x1b[0;31m",
        "Green":  "\x1b[0;32m",
        "Yellow": "\x1b[0;33m",
        "Blue":   "\x1b[0;34m",
        "Purple": "\x1b[0;35m",
        "Cyan":   "\x1b[0;36m",
        "White":  "\x1b[0;37m",
    }
    reset := "\x1b[0m"
    colourOption := colour[option]
    fmt.Printf("%s%s%s\n", colourOption, message, reset)
}

