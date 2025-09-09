package aphrodite

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"
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
Randomly choose a colour for you from: Black, Red, Green, Yellow, Blue, Purple, Cyan, White and prints it
*/
func Print(message string) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := rand.Intn(len(colour))

	keys := make([]string, 0, len(colour))
	keys = append(keys, keys...)

	var colourChoice string = colour[keys[randomIndex]]

	PrintColour(colourChoice, message)
}

/*
Prints any table that is passed into it - currently underdevelopment looking into how this could be more safely implimented
Pass in a map of any values and returns a printed table of the key / values
*/
func PrintTable(m any) {
	v := reflect.ValueOf(m)

	// Get key and value types
	for _, key := range v.MapKeys() {
		val := v.MapIndex(key)
		keyColour, _ := ReturnColour("Blue", fmt.Sprintf("%v", key.Interface()))
		valueColour, _ := ReturnColour("Green", fmt.Sprintf("%v", val.Interface()))
		fmt.Printf("%v: %v\n", keyColour, valueColour)
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
