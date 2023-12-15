package handleFiles

import (
	"os"
	"os/exec"
)

func CreateProject(project string) {
	switch project {
	case "React":
		createReactProject()
	case "React Native":
		createExpoApp()
	case "React with T3-Stack":
		createReactT3StackProject()
	}

}

func createReactProject() {
	cmd := exec.Command("npx", "create-react-app", "./")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func createReactT3StackProject() {
	cmd := exec.Command("npx", "create-t3-app", "./")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func createExpoApp() {
	cmd := exec.Command("npx", "expo-cli", "init", "./")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
