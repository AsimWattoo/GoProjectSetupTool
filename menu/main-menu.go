package menu

import (
	"fmt"
	"os"
	"path/filepath"
	"tool/node-backend-project/utils"
)

func ShowMainMenu() {
	mainMenu := []utils.MenuOption{
		{Name: "1. Create Project", Handler: createProjectHandler},
		{Name: "2. Delete Project", Handler: deleteProject},
		{Name: "3. Exit", Handler: exitHandler},
	}

	utils.MenuLoop(mainMenu)
}

func createProjectHandler() bool {
	projectsMenu := []utils.MenuOption{
		{Name: "1. Create Node Ts Project", Handler: CreateNodeTsProject},
		{Name: "2. Exit", Handler: exitHandler},
	}

	utils.MenuLoop(projectsMenu)
	return true
}

func deleteProject() bool {
	directoryPath := filepath.Join(playground, projectDir)

	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		fmt.Println("Project directory does not exist.")
		return true
	}

	if err := os.RemoveAll(directoryPath); err != nil {
		fmt.Println("Error deleting project directory:", err)
		return true
	}

	fmt.Println("Project deleted successfully.")
	return true
}

func exitHandler() bool {
	fmt.Println("Exiting...")
	return false
}
