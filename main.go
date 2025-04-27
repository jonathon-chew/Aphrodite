package aphrodite 

import "fmt"

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
			"x1b[1;30m":"Black",
			"x1b[1;31m":"Red",
			"x1b[1;32m":"Green",
			"1b[1;33m":"Yellow",
			"1b[1;34m":"Blue",
			"1b[1;35m":"Purple",
			"1b[1;36m":"Cyan",
			"1b[1;37m":"White",
		}

		Underline := map[string]string{
			"1b[4;30m":"Black",
			"1b[4;31m":"Red",
			"1b[4;32m":"Green",
			"1b[4;33m":"Yellow",
			"1b[4;34m":"Blue",
			"1b[4;35m":"Purple",
			"1b[4;36m":"Cyan",
			"1b[4;37m":"White",
		}

		Background := map[string]string{
			"1b[40m":"Black",
			"1b[41m":"Red",
			"1b[42m":"Green",
			"1b[43m":"Yellow",
			"1b[44m":"Blue",
			"1b[45m":"Purple",
			"1b[46m":"Cyan",
			"1b[47m":"White",
		}

		High_Intensity := map[string]string{
			"1b[0;90m":"Black",
			"1b[0;91m":"Red",
			"1b[0;92m":"Green",
			"1b[0;93m":"Yellow",
			"1b[0;94m":"Blue",
			"1b[0;95m":"Purple",
			"1b[0;96m":"Cyan",
			"1b[0;97m":"White",
		}

		Bold_High_Intensity := map[string]string{
			"1b[1;90m":"Black",
			"1b[1;91m":"Red",
			"1b[1;92m":"Green",
			"1b[1;93m":"Yellow",
			"1b[1;94m":"Blue",
			"1b[1;95m":"Purple",
			"1b[1;96m":"Cyan",
			"1b[1;97m":"White",
		}

		High_Intensity_backgrounds := map[string]string{
			"1b[0;100m":"Black",
			"1b[0;101m":"Red",
			"1b[0;102m":"Green",
			"1b[0;103m":"Yellow",
			"1b[0;104m":"Blue",
			"1b[0;105m":"Purple",
			"1b[0;106m":"Cyan",
			"1b[0;107m":"White",
		}

		reset := "\x1b[0m"

		if option == "Color" || option == "color" || option == "Colour" || option == "colour"{
			colourOption := colour[color]
			fmt.Printf("%s%s%s\n", colourOption, message, reset)
		}
		
		if option == "Bold" {
			fmt.Printf("%s", Bold[color])
		}

		if option == "Underline" {
			fmt.Printf("%s", Underline[color])
		}

		if option == "Background" {
			fmt.Printf("%s", Background[color])
		}

	if option == "High_Intensity" {
			fmt.Printf("%s", High_Intensity[color])
		}

		if option == "Bold_High_Intensity" {
			fmt.Printf("%s", Bold_High_Intensity[color])
		}

		if option == "High_Intensity_backgrounds" {
			fmt.Printf("%s", High_Intensity_backgrounds[color])
		}

}

