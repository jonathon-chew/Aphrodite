package aphrodite

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

// Colour prints the message in the specified color.
func Colour(option, color, message string) error {
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

	Bold := map[string]string{
		"Black":  "\x1b[1;30m",
		"Red":    "\x1b[1;31m",
		"Green":  "\x1b[1;32m",
		"Yellow": "\x1b[1;33m",
		"Blue":   "\x1b[1;34m",
		"Purple": "\x1b[1;35m",
		"Cyan":   "\x1b[1;36m",
		"White":  "\x1b[1;37m",
	}

	Underline := map[string]string{
		"Black":  "\x1b[4;30m",
		"Red":    "\x1b[4;31m",
		"Green":  "\x1b[4;32m",
		"Yellow": "\x1b[4;33m",
		"Blue":   "\x1b[4;34m",
		"Purple": "\x1b[4;35m",
		"Cyan":   "\x1b[4;36m",
		"White":  "\x1b[4;37m",
	}

	Background := map[string]string{
		"Black":  "\x1b[40m",
		"Red":    "\x1b[41m",
		"Green":  "\x1b[42m",
		"Yellow": "\x1b[43m",
		"Blue":   "\x1b[44m",
		"Purple": "\x1b[45m",
		"Cyan":   "\x1b[46m",
		"White":  "\x1b[47m",
	}

	High_Intensity := map[string]string{
		"Black":  "\x1b[0;90m",
		"Red":    "\x1b[0;91m",
		"Green":  "\x1b[0;92m",
		"Yellow": "\x1b[0;93m",
		"Blue":   "\x1b[0;94m",
		"Purple": "\x1b[0;95m",
		"Cyan":   "\x1b[0;96m",
		"White":  "\x1b[0;97m",
	}

	Bold_High_Intensity := map[string]string{
		"Black":  "\x1b[1;90m",
		"Red":    "\x1b[1;91m",
		"Green":  "\x1b[1;92m",
		"Yellow": "\x1b[1;93m",
		"Blue":   "\x1b[1;94m",
		"Purple": "\x1b[1;95m",
		"Cyan":   "\x1b[1;96m",
		"White":  "\x1b[1;97m",
	}

	High_Intensity_backgrounds := map[string]string{
		"Black":  "\x1b[0;100m",
		"Red":    "\x1b[0;101m",
		"Green":  "\x1b[0;102m",
		"Yellow": "\x1b[0;103m",
		"Blue":   "\x1b[0;104m",
		"Purple": "\x1b[0;105m",
		"Cyan":   "\x1b[0;106m",
		"White":  "\x1b[0;107m",
	}

	reset := "\x1b[0m"

	acceptableOptions := []string{"Colour", "Color", "Bold", "Underline", "Background", "High_intensity", "Bold_high_intesity", "High_Intensity_backgrounds"}
	acceptableColours := []string{"Black", "Red", "Green", "Yellow", "Blue", "Purple", "Cyan", "White", "Rainbow"}

	var optionChoice string = strings.Replace(option, " ", "_", -1)
	var colourChoice string = strings.ToUpper(string(color[0])) + color[1:]

	if !listContains(acceptableOptions, option){
		return errors.New(fmt.Sprint("Unable to recognise option: ", optionChoice))
	}

	if !listContains(acceptableColours, colourChoice) {
		return errors.New(fmt.Sprintf("Unfortunately %s isn't a useable colour", colourChoice))
	}

	if strings.ToLower(optionChoice) == "color" || strings.ToLower(optionChoice) == "colour" {
		if colourChoice == "Rainbow" {
			messageLength := len(message)
			for i := 0; i < messageLength; i++ {
				r, g, b := rainbow(i)
				fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m%s", r, g, b, message[i], reset)
			}
		} else {
			colourOfOptionPicked := colour[colourChoice]
			fmt.Printf("%s%s%s", colourOfOptionPicked, message, reset)
		}
		return nil
	}

	if strings.ToLower(optionChoice) == "bold" {
		colourOfOptionPicked := Bold[colourChoice]
		fmt.Printf("%s%s%s", colourOfOptionPicked, message, reset)
		return nil
	}

	if strings.ToLower(optionChoice) == "underline" {
		colourOfOptionPicked := Underline[colourChoice]
		fmt.Printf("%s%s%s", colourOfOptionPicked, message, reset)
		return nil
	}

	if strings.ToLower(optionChoice) == "background" {
		colourOfOptionPicked := Background[colourChoice]
		fmt.Printf("%s%s%s%s", colourOfOptionPicked, message, reset, reset)
		return nil
	}

	if strings.ToLower(optionChoice) == "high_intensity" {
		colourOfOptionPicked := High_Intensity[colourChoice]
		fmt.Printf("%s%s%s", colourOfOptionPicked, message, reset)
		return nil
	}

	if strings.ToLower(optionChoice) == "bold_high_intensity" {
		colourOfOptionPicked := Bold_High_Intensity[colourChoice]
		fmt.Printf("%s%s%s", colourOfOptionPicked, message, reset)
		return nil
	}

	if strings.ToLower(optionChoice) == "high_intensity_backgrounds" {
		colourOfOptionPicked := High_Intensity_backgrounds[colourChoice]
		fmt.Printf("%s%s%s", colourOfOptionPicked, message, reset)
		return nil
	}
	return nil
}

func listContains(s []string, comparison string) bool {
	for _, i := range s {
		if i == comparison {
			return true
		}
	}
	return false
}

func PrintPadR(s string, totalLength int) {
	padding := totalLength + len(s)
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

func PrintPadL(s string, totalLength int) {
	padding := totalLength + len(s)
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

func PrintPadRT(s string, totalLength int, flags []string){
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

func PrintPadLT(s string, totalLength int, flags []string) {
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

func rainbow(i int) (int, int, int) {
	var f = 0.1
	return int(math.Sin(f*float64(i)+0)*127 + 128),
		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}
