package aphrodite

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// (#6) TODO: Add the ability to pass in format strings to not rely on fmt.Sprtintf to make the string to pass into the function every time

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
			return "\033[38;2;" + r + g + b + string(message[i]) + reset, nil
		}
	case "Empty":
		return message, nil
	default:
		var colourPicked string = colour[colourChoice]
		return colourPicked + message + reset, nil
	}
	return "", nil
}

/*
This ignores warnings from a malformed message, to be used quickly when the message will be known prior to use to be safe!
*/
func ReturnError(message string) string {
	returnMessage, _ := ReturnColour("Red", message)
	return returnMessage
}

/*
This ignores warnings from a malformed message, to be used quickly when the message will be known prior to use to be safe!
*/
func ReturnInfo(message string) string {
	returnMessage, _ := ReturnColour("Green", message)
	return returnMessage
}

/*
This ignores warnings from a malformed message, to be used quickly when the message will be known prior to use to be safe!
*/
func ReturnWarning(message string) string {
	returnMessage, _ := ReturnColour("Yellow", message)
	return returnMessage
}

/*
Randomly choose a colour for you from: Black, Red, Green, Yellow, Blue, Purple, Cyan, White
Can error if colour not found
*/
func Return(message string) (string, error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	keys := make([]string, 0, len(colour))
	for key := range colour {
		keys = append(keys, key) // Append keys from the map
	}

	// Ensure there are keys to choose from
	if len(keys) == 0 {
		return "", fmt.Errorf("no colors available")
	}

	// Generate a random index
	randomIndex := r.Intn(len(keys)) // Use len(keys) directly

	// Get the random color choice
	colourChoice := keys[randomIndex]

	// Print the chosen color for debugging
	// fmt.Println("Available colors:", keys)
	// // fmt.Println("Random index:", randomIndex)
	// fmt.Println("Chosen color:", colourChoice)

	// Call ReturnColour with the chosen color
	return ReturnColour(colourChoice, message)
}

/*
Options: r, g, b have to be 0 or bigger and under 255
Can error if the values are not the right size
*/
func ReturnRGBColour(r, g, b int, message string) (string, error) {
	if r >= 0 && r < 256 && g >= 0 && g < 265 && b >= 0 && b < 265 {
		var red string = strconv.Itoa(r)
		var green string = strconv.Itoa(g)
		var blue string = strconv.Itoa(b)

		return "\033[38;2;" + red + green + blue + "m" + message + reset, nil
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
	return colourPicked + message + reset, nil
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
	return colourPicked + message + reset, nil
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
	return colourPicked + message + reset, nil
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
	return colourPicked + message + reset, nil
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
	return colourPicked + message + reset, nil
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
	return colourPicked + message + reset, nil
}
