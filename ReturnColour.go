package aphrodite

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

/*
Options: Black, Red, Green, Yellow, Blue, Purple, Cyan, White
Can error if colour not found
*/
func ReturnColour(color, message string) (string, error) {

	var colourChoice string = strings.ToUpper(string(color[0])) + color[1:]

	if !checkColourInList(colourChoice) {
		return "", unacceptableColour(colourChoice)
	}

	switch colourChoice {
	case "Rainbow":
		messageLength := len(message)
		for i := 0; i < messageLength; i++ {
			r, g, b := rainbow(i)
			return fmt.Sprintf("\033[38;2;%d;%d;%dm%c\033[0m%s", r, g, b, message[i], reset), nil
		}
	case "Empty":
		return fmt.Sprint(message), nil
	default:
		var colourPicked string = colour[colourChoice]
		return fmt.Sprintf("%s%s%s", colourPicked, message, reset), nil
	}
	return "", nil
}

/*
Randomly choose a colour for you from: Black, Red, Green, Yellow, Blue, Purple, Cyan, White
Can error if colour not found
*/
func Return(message string) (string, error) {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	randomIndex := rand.Intn(len(colour))

	keys := make([]string, 0, len(colour))
	keys = append(keys, keys...)

	var colourChoice string = colour[keys[randomIndex]]

	return ReturnColour(colourChoice, message)
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
