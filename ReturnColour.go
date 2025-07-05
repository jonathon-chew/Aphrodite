package aphrodite

import (
	"fmt"
	"slices"
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

func checkColourInList(colourchoice string) bool {
	acceptableColours := []string{"Black", "Red", "Green", "Yellow", "Blue", "Purple", "Cyan", "White", "Rainbow"}
	return !slices.Contains(acceptableColours, colourchoice)
}

func unacceptableColour(color string) error {
	keys := make([]string, 0, len(colour))
	for k := range colour {
		keys = append(keys, k)
	}

	var options string = strings.Join(keys, ", ")

	return fmt.Errorf("%s isn't avaliable from the list %s", color, options)
}

/*
Options: Black, Red, Green, Yellow, Blue, Purple, Cyan, White
Can error if colour not found
*/
func ReturnColour(color, message string) (string, error) {

	var colourChoice string = strings.ToUpper(string(color[0])) + color[1:]

	if !checkColourInList(colourChoice) {
		return "", unacceptableColour(colourChoice)
	}

	if colourChoice == "Rainbow" {
		messageLength := len(message)
		for i := 0; i < messageLength; i++ {
			r, g, b := rainbow(i)
			return fmt.Sprintf("\033[38;2;%d;%d;%dm%c\033[0m%s", r, g, b, message[i], reset), nil
		}
	} else {
		var colourPicked string = colour[colourChoice]
		return fmt.Sprintf("%s%s%s", colourPicked, message, reset), nil
	}
	return "", nil
}

/*
Options: r, g, b have to be 0 or bigger and under 255
Can error if the values are not the right size
*/
func ReturnRGBColour(r, g, b int, message string) (string, error) {
	if r >= 0 && r < 256 && g >= 0 && g < 265 && b >= 0 && b < 265 {
		return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m%s", r, g, b, message, reset), nil
	}

	return "", fmt.Errorf("red: %d, green: %d, b: %d are the vales passed in, and should each be between 0 and 256", r, g, b)
}

/*
Options: Black, Red, Green, Yellow, Blue, Purple, Cyan, White
Can error if colour not found
*/
func ReturnBold(colourChoice, message string) (string, error) {
	var colourPicked string = bold[strings.ToUpper(string(colourChoice[0]))+colourChoice[1:]]
	if !checkColourInList(colourChoice) {
		return "", unacceptableColour(colourPicked)
	}
	return fmt.Sprintf("%s%s%s", colourPicked, message, reset), nil
}

/*
Options: Black, Red, Green, Yellow, Blue, Purple, Cyan, White
Can error if colour not found
*/
func ReturnUnderline(colourChoice, message string) (string, error) {
	var colourPicked string = underline[strings.ToUpper(string(colourChoice[0]))+colourChoice[1:]]
	if !checkColourInList(colourChoice) {
		return "", unacceptableColour(colourPicked)
	}
	return fmt.Sprintf("%s%s%s", colourPicked, message, reset), nil
}

/*
Options: Black, Red, Green, Yellow, Blue, Purple, Cyan, White
Can error if colour not found
*/
func ReturnBackground(colourChoice, message string) (string, error) {
	var colourPicked string = background[strings.ToUpper(string(colourChoice[0]))+colourChoice[1:]]
	if !checkColourInList(colourChoice) {
		return "", unacceptableColour(colourPicked)
	}
	return fmt.Sprintf("%s%s%s%s", colourPicked, message, reset, reset), nil
}

/*
Options: Black, Red, Green, Yellow, Blue, Purple, Cyan, White
Can error if colour not found
*/
func ReturnHighIntensity(colourChoice, message string) (string, error) {
	var colourPicked string = high_Intensity[strings.ToUpper(string(colourChoice[0]))+colourChoice[1:]]
	if !checkColourInList(colourChoice) {
		return "", unacceptableColour(colourPicked)
	}
	return fmt.Sprintf("%s%s%s", colourPicked, message, reset), nil
}

/*
Options: Black, Red, Green, Yellow, Blue, Purple, Cyan, White
Can error if colour not found
*/
func ReturnBoldHighIntensity(colourChoice, message string) (string, error) {
	var colourPicked string = bold_High_Intensity[strings.ToUpper(string(colourChoice[0]))+colourChoice[1:]]
	if !checkColourInList(colourChoice) {
		return "", unacceptableColour(colourPicked)
	}
	return fmt.Sprintf("%s%s%s", colourPicked, message, reset), nil
}

/*
Options: Black, Red, Green, Yellow, Blue, Purple, Cyan, White
Can error if colour not found
*/
func ReturnHighIntensityBackgrounds(colourChoice, message string) (string, error) {
	var colourPicked string = high_Intensity_backgrounds[strings.ToUpper(string(colourChoice[0]))+colourChoice[1:]]
	if !checkColourInList(colourChoice) {
		return "", unacceptableColour(colourPicked)
	}
	return fmt.Sprintf("%s%s%s", colourPicked, message, reset), nil
}
