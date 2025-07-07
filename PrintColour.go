package aphrodite

import (
	"fmt"
	"strings"
)

/*
Options: Black, Red, Green, Yellow, Blue, Purple, Cyan, White
Can error if colour not found
*/
func PrintColour(color, message string) {

	var colourChoice string = strings.ToUpper(string(color[0])) + color[1:]

	switch colourChoice {
	case "Rainbow":
		messageLength := len(message)
		for i := 0; i < messageLength; i++ {
			r, g, b := rainbow(i)
			fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m%s", r, g, b, message[i], reset)
		}
	case "Empty":
		fmt.Print(message)
	default:
		var colourPicked string = colour[colourChoice]
		fmt.Printf("%s%s%s", colourPicked, message, reset)
	}
}

/*
Options: r, g, b have to be 0 or bigger and under 255
Can error if the values are not the right size
*/
func PrintRGBColour(r, g, b int, message string) {
	fmt.Printf("\033[38;2;%d;%d;%dm%s\033[0m%s", r, g, b, message, reset)
}

/*
Options: Black, Red, Green, Yellow, Blue, Purple, Cyan, White
Can error if colour not found
*/
func PrintBold(colourChoice, message string) {
	var colourPicked string = bold[strings.ToUpper(string(colourChoice[0]))+colourChoice[1:]]

	fmt.Printf("%s%s%s", colourPicked, message, reset)
}

/*
Options: Black, Red, Green, Yellow, Blue, Purple, Cyan, White
Can error if colour not found
*/
func PrintUnderline(colourChoice, message string) {
	var colourPicked string = underline[strings.ToUpper(string(colourChoice[0]))+colourChoice[1:]]

	fmt.Printf("%s%s%s", colourPicked, message, reset)
}

/*
Options: Black, Red, Green, Yellow, Blue, Purple, Cyan, White
Can error if colour not found
*/
func PrintBackground(colourChoice, message string) {
	var colourPicked string = background[strings.ToUpper(string(colourChoice[0]))+colourChoice[1:]]

	fmt.Printf("%s%s%s%s", colourPicked, message, reset, reset)
}

/*
Options: Black, Red, Green, Yellow, Blue, Purple, Cyan, White
Can error if colour not found
*/
func PrintHighIntensity(colourChoice, message string) {
	var colourPicked string = high_Intensity[strings.ToUpper(string(colourChoice[0]))+colourChoice[1:]]

	fmt.Printf("%s%s%s", colourPicked, message, reset)
}

/*
Options: Black, Red, Green, Yellow, Blue, Purple, Cyan, White
Can error if colour not found
*/
func PrintBoldHighIntensity(colourChoice, message string) {
	var colourPicked string = bold_High_Intensity[strings.ToUpper(string(colourChoice[0]))+colourChoice[1:]]

	fmt.Printf("%s%s%s", colourPicked, message, reset)
}

/*
Options: Black, Red, Green, Yellow, Blue, Purple, Cyan, White
Can error if colour not found
*/
func PrintHighIntensityBackgrounds(colourChoice, message string) {
	var colourPicked string = high_Intensity_backgrounds[strings.ToUpper(string(colourChoice[0]))+colourChoice[1:]]

	fmt.Printf("%s%s%s", colourPicked, message, reset)
}
