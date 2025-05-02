package projectmenu

import (
	"fmt"

	"github.com/AsimWattoo/go-utilities/menu"
)

func ShowMainMenu() {
	mainMenu := []menu.MenuOption{
		{Name: "1. Create Project", Handler: createProjectHandler},
		// {Name: "2. Clean Playground", Handler: deleteProject},
		{Name: "2. Exit", Handler: exitHandler},
	}

	menu.MenuLoop(mainMenu)
}

func createProjectHandler() bool {
	projectsMenu := []menu.MenuOption{
		{Name: "1. Create Node Ts Project", Handler: CreateNodeTsProject},
		{Name: "2. Create React + Vite + TS + Tailwind Project", Handler: CreateViteTsProject},
		{Name: "3. Exit", Handler: exitHandler},
	}

	menu.MenuLoop(projectsMenu)
	return true
}

// func deleteProject() bool {

// 	entries, err := os.ReadDir(playground)
// 	if err != nil {
// 		fmt.Println("Error reading playground directory:", err)
// 		return true
// 	}

// 	for _, entry := range entries {
// 		path := filepath.Join(playground, entry.Name())
// 		err := os.RemoveAll(path)
// 		if err != nil {
// 			fmt.Println("Error deleting directory:", err)
// 			return true
// 		}
// 	}

// 	fmt.Println("Playground Cleared Successfully.")
// 	return true
// }

func exitHandler() bool {
	fmt.Println("Exiting...")
	return false
}
