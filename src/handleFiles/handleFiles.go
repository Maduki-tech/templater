package handleFiles

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/maduki-tech/templater/src/utils"
)

func HandleFiles(selectedItem string) {
	switch selectedItem {
	case "C":
		CreateCFile("makefile")
		// clear the terminal
		utils.SuccesfullyRun("C", "makefile")
	case "C++":
		fmt.Println("C++")
	case "React":
		CreateProject(selectedItem)

	case "React Native":
		CreateProject(selectedItem)

	case "React with T3-Stack":
		CreateProject(selectedItem)
	}

}

func clearTerminal() {
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
