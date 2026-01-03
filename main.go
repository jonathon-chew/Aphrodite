package aphrodite

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

var bold = map[string]string{
	"Black":  "\x1b[1;30m",
	"Red":    "\x1b[1;31m",
	"Green":  "\x1b[1;32m",
	"Yellow": "\x1b[1;33m",
	"Blue":   "\x1b[1;34m",
	"Purple": "\x1b[1;35m",
	"Cyan":   "\x1b[1;36m",
	"White":  "\x1b[1;37m",
}

var underline = map[string]string{
	"Black":  "\x1b[4;30m",
	"Red":    "\x1b[4;31m",
	"Green":  "\x1b[4;32m",
	"Yellow": "\x1b[4;33m",
	"Blue":   "\x1b[4;34m",
	"Purple": "\x1b[4;35m",
	"Cyan":   "\x1b[4;36m",
	"White":  "\x1b[4;37m",
}

var background = map[string]string{
	"Black":  "\x1b[40m",
	"Red":    "\x1b[41m",
	"Green":  "\x1b[42m",
	"Yellow": "\x1b[43m",
	"Blue":   "\x1b[44m",
	"Purple": "\x1b[45m",
	"Cyan":   "\x1b[46m",
	"White":  "\x1b[47m",
}

var high_Intensity = map[string]string{
	"Black":  "\x1b[0;90m",
	"Red":    "\x1b[0;91m",
	"Green":  "\x1b[0;92m",
	"Yellow": "\x1b[0;93m",
	"Blue":   "\x1b[0;94m",
	"Purple": "\x1b[0;95m",
	"Cyan":   "\x1b[0;96m",
	"White":  "\x1b[0;97m",
}

var bold_High_Intensity = map[string]string{
	"Black":  "\x1b[1;90m",
	"Red":    "\x1b[1;91m",
	"Green":  "\x1b[1;92m",
	"Yellow": "\x1b[1;93m",
	"Blue":   "\x1b[1;94m",
	"Purple": "\x1b[1;95m",
	"Cyan":   "\x1b[1;96m",
	"White":  "\x1b[1;97m",
}

var high_Intensity_backgrounds = map[string]string{
	"Black":  "\x1b[0;100m",
	"Red":    "\x1b[0;101m",
	"Green":  "\x1b[0;102m",
	"Yellow": "\x1b[0;103m",
	"Blue":   "\x1b[0;104m",
	"Purple": "\x1b[0;105m",
	"Cyan":   "\x1b[0;106m",
	"White":  "\x1b[0;107m",
}

var colour = map[string]string{
	"Black":  "\x1b[0;30m",
	"Red":    "\x1b[0;31m",
	"Green":  "\x1b[0;32m",
	"Yellow": "\x1b[0;33m",
	"Blue":   "\x1b[0;34m",
	"Purple": "\x1b[0;35m",
	"Cyan":   "\x1b[0;36m",
	"White":  "\x1b[0;37m",
}

var reset string = "\x1b[0m"

/*
Options: Black, Red, Green, Yellow, Blue, Purple, Cyan, White
*/
func Colour(option, color, message string) error {

	acceptableOptions := []string{"Colour", "Color", "Bold", "Underline", "Background", "High_intensity", "Bold_high_intesity", "High_Intensity_backgrounds"}
	acceptableColours := []string{"Black", "Red", "Green", "Yellow", "Blue", "Purple", "Cyan", "White", "Rainbow"}

	var optionChoice string = strings.Replace(option, " ", "_", -1)
	var colourChoice string = strings.ToUpper(string(color[0])) + color[1:]

	var allAcceptableOptions string = strings.Join(acceptableOptions, ", ")
	var allAcceptableColours string = strings.Join(acceptableColours, ", ")

	if !slices.Contains(acceptableOptions, option) {
		return fmt.Errorf("unable to recognise option: %s the options are: %s", optionChoice, allAcceptableOptions)
	}

	if !slices.Contains(acceptableColours, colourChoice) {
		return fmt.Errorf("unfortunately %s isn't a useable colour the options are: %s", colourChoice, allAcceptableColours)
	}

	var command string = strings.ToLower(optionChoice)

	if command == "color" || command == "colour" || command == "text" || command == "rune" {
		if colourChoice == "Rainbow" {
			messageLength := len(message)
			for i := 0; i < messageLength; i++ {
				r, g, b := rainbow(i)
				fmt.Printf("\033[38;2;%s;%s;%sm%c\033[0m%s", r, g, b, message[i], reset)
			}
		} else {
			colourOfOptionPicked := colour[colourChoice]
			fmt.Printf("%s%s%s", colourOfOptionPicked, message, reset)
		}
		return nil
	}

	if command == "bold" {
		colourOfOptionPicked := bold[colourChoice]
		fmt.Printf("%s%s%s", colourOfOptionPicked, message, reset)
		return nil
	}

	if command == "underline" {
		colourOfOptionPicked := underline[colourChoice]
		fmt.Printf("%s%s%s", colourOfOptionPicked, message, reset)
		return nil
	}

	if command == "background" {
		colourOfOptionPicked := background[colourChoice]
		fmt.Printf("%s%s%s%s", colourOfOptionPicked, message, reset, reset)
		return nil
	}

	if command == "high_intensity" {
		colourOfOptionPicked := high_Intensity[colourChoice]
		fmt.Printf("%s%s%s", colourOfOptionPicked, message, reset)
		return nil
	}

	if command == "bold_high_intensity" {
		colourOfOptionPicked := bold_High_Intensity[colourChoice]
		fmt.Printf("%s%s%s", colourOfOptionPicked, message, reset)
		return nil
	}

	if command == "high_intensity_backgrounds" {
		colourOfOptionPicked := high_Intensity_backgrounds[colourChoice]
		fmt.Printf("%s%s%s", colourOfOptionPicked, message, reset)
		return nil
	}
	return nil
}

func PrintPadR(s string, addLength int) {
	padding := addLength + len(s)
	var i = 0
	newString := s
	if padding > 0 {
		for i < padding {
			newString = newString + " "
			i++
		}
	}

	fmt.Printf("%s", newString)
}

func checkColourInList(colourchoice string) bool {
	acceptableColours := []string{"Black", "Red", "Green", "Yellow", "Blue", "Purple", "Cyan", "White", "Rainbow", "Empty"}
	return slices.Contains(acceptableColours, colourchoice)
}

func unacceptableColour(color string) error {
	keys := make([]string, 0, len(colour))
	for k := range colour {
		keys = append(keys, k)
	}

	var options string = strings.Join(keys, ", ")

	return fmt.Errorf("%s isn't avaliable from the list %s", color, options)
}

func PrintPadL(s string, addLength int) {
	padding := addLength + len(s)
	var i = 0
	newString := s
	if padding > 0 {
		for i < padding {
			newString = " " + newString
			i++
		}
	}
	fmt.Printf("%s", newString)
}

func PrintPadRT(s string, totalLength int) {
	padding := totalLength - len(s)
	var i = 0
	newString := s
	if padding > 0 {
		for i < padding {
			newString = newString + " "
			i++
		}
	}

	fmt.Printf("%s", newString)
}

func PrintPadLT(s string, totalLength int) {
	padding := totalLength - len(s)
	var i = 0
	newString := s
	if padding > 0 {
		for i < padding {
			newString = " " + newString
			i++
		}
	}

	fmt.Printf("%s", newString)
}

func ReturnPadR(s string, addLength int) string {
	padding := addLength + len(s)
	var i = 0
	newString := s
	if padding > 0 {
		for i < padding {
			newString = newString + " "
			i++
		}
	}

	return newString
}

func ReturnPadL(s string, addLength int) string {
	padding := addLength + len(s)
	var i = 0
	newString := s
	if padding > 0 {
		for i < padding {
			newString = " " + newString
			i++
		}
	}
	return newString
}

func ReturnPadRT(s string, totalLength int) string {
	padding := totalLength - len(s)
	var i = 0
	newString := s
	if padding > 0 {
		for i < padding {
			newString = newString + " "
			i++
		}
	}

	return newString
}

func ReturnPadLT(s string, totalLength int) string {
	padding := totalLength - len(s)
	var i = 0
	newString := s
	if padding > 0 {
		for i < padding {
			newString = " " + newString
			i++
		}
	}

	return newString
}

func rainbow(i int) (string, string, string) {
	var f = 0.1

	// temp := int(math.Sin(f * float64(i)))
	// suffix := 127 + 128 // 255

	return strconv.Itoa(int(math.Sin(f*float64(i)+0)*127 + 128)),
		strconv.Itoa(int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128)),
		strconv.Itoa(int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128))
}
