package handleFiles

import (
	"fmt"
	"os"

)

func handleFiles(selectedItem string) {
	switch selectedItem {
	case "C":
		fmt.Println("C")
		os.Exit(0)
	case "C++":
		fmt.Println("C++")
		os.Exit(0)
	}
}
