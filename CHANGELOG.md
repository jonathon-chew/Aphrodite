# CHANGELOG

	## v1.6.2, tag: v1.6.1, tag: v1.6.0 

		### UPDATES
		1. README clean up

	## v1.5.1 

		### NEW
		1. ANSI check, it'll only output ansi if the output is a terminal
		### UPDATES
		1. minor changes to README
		1. README cleanup

	## v1.5.0 

		### UPDATES
		1. scripts to match other repos for some uniformality
		### DELETES
		1. removed the work in progress from the README as this is being used in several of my projects

	## v1.4.0 

		### NEW
		1. adding print bytes as an option

	## v1.3.36 

		### NEW
		1. CICD script for updating and checking!
		1. Adding an image to the Readme
		1. Adding an image to the Readme
		1. todos based on useage
		1. todos based on useage
		### UPDATES
		1. making the use of Thoth from $PATH not from inside the local folder
		1. README.md to inlcude credit for the origional gopher design
		1. docs folder to doc
		1. README.md to inlcude credit for the origional gopher design
		1. docs folder to doc
		### REFACTORS
		1. remove SprintF for better performance
		### MISC
		1. Merge pull request #7 from jonathon-chew/Self_Formatting_Strings
		1. Merge branch 'main' into Self_Formatting_Strings

	## v1.3.35 

		### UPDATES
		1. adding error message below the Thoth todo comment
		1. adding error message below the Thoth todo comment

	## v1.3.34 

		### REFACTORS
		1. fixed a bug with warning, and added tests for warning/info/error functions
		1. fixed a bug with warning, and added tests for warning/info/error functions

	## v1.3.33, tag: v1.3.32 

		### NEW
		1. added the ability to quickly print or return with default colours for warning / info and error
		1. added the ability to quickly print or return with default colours for warning / info and error

	## v1.2.31 

		### MISC
		1. fix randomness in return and updated from rand.Seed to rand.New to match good practice
		1. fix randomness in return and updated from rand.Seed to rand.New to match good practice

	## v1.2.3 

		### MISC
		1. updated checkColourInList as double negative was returning incorrectly
		1. updated checkColourInList as double negative was returning incorrectly

	## v1.2.2 

		### MISC
		1. Add the ability to pretty pretty any table/map
		1. Adding the ability to not care what colour is printed/returned and just picks a random one for your
		1. empty as an option of acceptable colours, refactor: convert some if statements to case statements
		1. errors to return options, moved maps to main.go so removed from here

	## v1.2.1 

		### MISC
		1. converted PrintColour to match ReturnColour, moved generic functions and maps to the main file
		1. formatting changes

	## v1.1.2 

		### MISC
		1. massive refactor to remove option parameter from ReturnColour function, and now in seperate functions to call, add in a feature to set their own colour for text colour

	## v1.1.1 

		### MISC
		1. options in the comments directly above the function for support with LSPs that support that

	## v1.1.0, tag: v0.1.0 

		### MISC
		1. Added options in the comments directly below the function for support with LSPs that support that

	## v1.0.25 

		### REFACTORS
		1. Moved a copy of a large amount of the function to unique files

	## v1.0.24 

		### MISC
		1. Added a better name for the variable, to aid understanding of the intention of the function when looking at the source code - whether the function would be adding the length of spaces on to what ever length of string is being passed in OR whether it will return a string the length specified

	## v1.0.23 

		### MISC
		1. Small refactor and allow multple ways for commond useage to be called

	## v1.0.22 

		### MISC
		1. Adding in functions for return to be seperate from print
		1. Adding better named functions, all current functions print and user knows from function name instead of passing in a variable

	## v1.0.21 

		### MISC
		1. Adding error messages and checking whether the user picked options in the slots are acceptable options!

	## v1.0.20 

		### MISC
		1. Adding an additional reset to rainbow

	## v1.0.19 

		### MISC
		1. Fixing an issue where this would repeat the message because it's not getting the character!

	## v1.0.18 

		### MISC
		1. Adding a rainbow option

	## v1.0.17 

		### MISC
		1. Added flags to total padding!

	## v1.0.16, tag: v1.0.15 

		### MISC
		1. Removing an imported package by gopls that shouldn't have been added!

	## v1.0.14 

		### MISC
		1. Attempting to add contional returns, print padtotal if flag present, else return the string so it can be used / assigned
		1. Formatting changes to the source code

	## v1.0.13 

		### MISC
		1. Removing new line characters to better reflect the idea that this is a drop in replacement, doesn't put on unexpected behaviour, pass in the new line if wanted

	## v1.0.12 

		### MISC
		1. Adding 2 functions to allow for total, so it will only add spacing if the number is greater than the length of string that was passed in

	## v1.0.11 

		### MISC
		1. Changing how padding works so the declaration of length isn't the final length, it's the amount of padding you want to add

	## v1.0.10 

		### MISC
		1. Trying new string as doesn't appear to be padding correctly in certain terminals

	## v1.0.9 

		### MISC
		1. Formatting, converting returns in padding to fmt.Printf to match the behaviour of Colour

	## v1.0.8 

		### MISC
		1. Adding Padding to Aphrodite
		1. Removing symbolic link to test folder
		1. Removing testing dir

	## v1.0.7 

		### MISC
		1. Fix an issue with escape characters
		1. Adding a rainbow to README.md
		1. Formatting the README.md for better readability
		1. Adding a git ignore
		1. Updated all references to Aphrodite with a capital as per personal preference

	## v1.0.6 

		### MISC
		1. Updating go.mod and seeing if it breaks when capatlaised

	## v1.0.4 

		### MISC
		1. Added a way to deal with captalisation as strings.ToTitle and strings.Upper do effectively the same thing, not capalitising the first character
		1. Adding option picking debug message

	## v1.0.3 

		### MISC
		1. Adding a debug message

	## v1.0.2 

		### MISC
		1. ReAdding in README and LICSENCE
		1. Revert "Added a README.md and License"
		1. Added a README.md and License

	## v1.0.1 

		### MISC
		1. Making the colour choice case insensative when picking
		1. whether the users calls by capital or not is ignored
		1. Format changes
		1. Adding early returns to stop looping through the function when done if an earlier choice is picked
		1. Added the right returns, from each option call, but currently only as print statements not as return statements
		1. Flipping order in the map for all maps other than colour so user calls the colour not the escape sequence
		1. Removing temp.go as not needed any more
		1. Formatting change
		1. Adding the majority of options, tag not updated

	## v1.0.0 

		### MISC
		1. Added an option package to main.go