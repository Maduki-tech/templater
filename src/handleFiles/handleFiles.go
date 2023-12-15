package handleFiles

import (
	"fmt"
)

func HandleFiles(selectedItem string) {
	switch selectedItem {
	case "C":
	CreateCFile("makefile")
	case "C++":
		fmt.Println("C++")
	}
}
