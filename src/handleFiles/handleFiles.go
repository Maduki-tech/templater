package handleFiles

import (
	"fmt"
	"os"
	"os/exec"
)

func HandleFiles(selectedItem string) {
	switch selectedItem {
	case "C":
		CreateCFile("makefile")
		cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()

		println(`
Succesfully created C file with makefile.
-------------------------------------------
To compile the project run: make
To run the project run: ./main

To clean the project run: make clean
To create a compile_commands.json file run: make compile_commands

-------------------------------------------
If you like the project please star it on github:
www.github.com/maduki-tech/templater
`)
	case "C++":
		fmt.Println("C++")
	}
}
