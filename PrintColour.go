package aphrodite

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

// (#5) TODO: Add the ability to pass in format strings to not rely on fmt.Sprtintf to make the string to pass into the function every time

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

	// (#4) TODO: This doesn't appear to work, sometimes/always appears to return index out of range
	//goroutine 1 [running]:
	// github.com/jonathon-chew/Aphrodite.Print({0x1400000e1f0, 0xb})
	// 	/Users/hunteradder626/go/pkg/mod/github.com/jonathon-chew/!aphrodite@v1.3.35/PrintColour.go:44 +0x1d8
	// main.main()
	// 	/Users/hunteradder626/Documents/Scripts/Go/Draft/Lines_Of_Code/main.go:86 +0x110
	// exit status 2
	var colourChoice string = colour[keys[randomIndex]]

	PrintColour(colourChoice, message)
}

/*
Uses the default colour of red - if you would like to determin you PrintError Colour use the function PrintColour instead
*/
func PrintError(message string){
	PrintColour("Red", message)
}

/*
Uses the default colour of Green- if you would like to determin you PrintError Colour use the function PrintColour instead
*/
func PrintInfo(message string){
	PrintColour("Green", message)
}

/*
This ignores warnings from a malformed message, to be used quickly when the message will be known prior to use to be safe!
*/
func PrintWarning(message string){
	PrintColour("Yellow", message)
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