package aphrodite

import (
	"fmt";
	"strings"
)

// Colour prints the message in the specified color.
func Colour(option, color, message string) {
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
		"Black":  "x1b[1;30m",
		"Red":    "x1b[1;31m",
		"Green":  "x1b[1;32m",
		"Yellow": "1b[1;33m",
		"Blue":   "1b[1;34m",
		"Purple": "1b[1;35m",
		"Cyan":   "1b[1;36m",
		"White":  "1b[1;37m",
	}

	Underline := map[string]string{
		"Black":  "1b[4;30m",
		"Red":    "1b[4;31m",
		"Green":  "1b[4;32m",
		"Yellow": "1b[4;33m",
		"Blue":   "1b[4;34m",
		"Purple": "1b[4;35m",
		"Cyan":   "1b[4;36m",
		"White":  "1b[4;37m",
	}

	Background := map[string]string{
		"Black":  "1b[40m",
		"Red":    "1b[41m",
		"Green":  "1b[42m",
		"Yellow": "1b[43m",
		"Blue":   "1b[44m",
		"Purple": "1b[45m",
		"Cyan":   "1b[46m",
		"White":  "1b[47m",
	}

	High_Intensity := map[string]string{
		"Black":  "1b[0;90m",
		"Red":    "1b[0;91m",
		"Green":  "1b[0;92m",
		"Yellow": "1b[0;93m",
		"Blue":   "1b[0;94m",
		"Purple": "1b[0;95m",
		"Cyan":   "1b[0;96m",
		"White":  "1b[0;97m",
	}

	Bold_High_Intensity := map[string]string{
		"Black":  "1b[1;90m",
		"Red":    "1b[1;91m",
		"Green":  "1b[1;92m",
		"Yellow": "1b[1;93m",
		"Blue":   "1b[1;94m",
		"Purple": "1b[1;95m",
		"Cyan":   "1b[1;96m",
		"White":  "1b[1;97m",
	}

	High_Intensity_backgrounds := map[string]string{
		"Black":  "1b[0;100m",
		"Red":    "1b[0;101m",
		"Green":  "1b[0;102m",
		"Yellow": "1b[0;103m",
		"Blue":   "1b[0;104m",
		"Purple": "1b[0;105m",
		"Cyan":   "1b[0;106m",
		"White":  "1b[0;107m",
	}

	reset := "\x1b[0m"

	if strings.ToLower(option) == "color" || strings.ToLower(option) == "colour" {
		colourOption := colour[color]
		fmt.Printf("%s%s%s\n", colourOption, message, reset)
		return
	}

	if strings.ToLower(option) == "bold" {
		colourOption := Bold[color]
		fmt.Printf("%s%s%s\n", colourOption, message, reset)
		return
	}

	if strings.ToLower(option) == "underline" {
		colourOption := Underline[color]
		fmt.Printf("%s%s%s\n", colourOption, message, reset)
		return
	}

	if strings.ToLower(option) == "background" {
		colourOption := Background[color]
		fmt.Printf("%s%s%s\n", colourOption, message, reset)
		return
	}

	if strings.ToLower(option) == "high_intensity" {
		colourOption := High_Intensity[color]
		fmt.Printf("%s%s%s\n", colourOption, message, reset)
		return
	}

	if strings.ToLower(option) == "bold_high_intensity" {
		colourOption := Bold_High_Intensity[color]
		fmt.Printf("%s%s%s\n", colourOption, message, reset)
		return
	}

	if strings.ToLower(option) == "high_intensity_backgrounds" {
		colourOption := High_Intensity_backgrounds[color]
		fmt.Printf("%s%s%s\n", colourOption, message, reset)
		return
	}

}
