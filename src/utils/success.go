package utils

/*
SuccesfullyRun prints a message to the terminal when the program has succesfully run

@param language string - The language you want to use for the project

@param complier string - The complier you want to use for the Project
*/
func SuccesfullyRun(language string, complier string) {
	println(`
Succesfully created ` + language + ` file with ` + complier + `.

-------------------------------------------

If you like the project please star it on github:
www.github.com/maduki-tech/templater
`)

}
