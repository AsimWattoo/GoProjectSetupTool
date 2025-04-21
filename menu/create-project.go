package menu

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"tool/node-backend-project/commands"
)

var playground string = "playground"
var projectDir string = "test-project"

func CreateNodeTsProject() bool {
	directoryPath := filepath.Join(playground, projectDir)
	executor := commands.NewCommandExecutor()

	executor.AddCommand("mkdir", directoryPath, "0755")
	executor.AddCommand("cd", directoryPath)
	executor.AddRawCommand("npm", "init", "-y")
	executor.AddRawCommand("npm", "install", "express", "dotenv", "cookie-parser", "tsconfig-paths", "cors", "express-session")
	executor.AddRawCommand("npm", "install", "@types/node", "@types/express", "ts-node", "typescript", "nodemon", "@types/cors", "@types/cookie-parser", "@types/express-session")
	executor.AddCommand("mkdir", "src", "0755")
	executor.AddCommand("mkdir", "uploads", "0755")
	executor.AddCommand("cp", "../../templates/index.ts", "src/index.ts")
	executor.AddCommand("cp", "../../templates/tsconfig.txt", "tsconfig.json")
	executor.AddCommand("json-update", "package.json", "scripts", "{\"start\": \"node src/index.ts\", \"dev\": \"nodemon --exec ts-node src/index.ts\"}")

	executor.AddCommand("cd", "../..")

	err := executor.ExecuteCommands()

	if err != nil {
		fmt.Printf("==> Command Execution Failed because of following error: %s\n", err)
	} else {
		fmt.Printf("==> Command Execution Successful\n")
	}

	return true
}

func PrevCreateNodeTsProject() bool {

	fmt.Println("Running npm init -y")
	cmd := exec.Command("npm", "init", "-y")
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Error initializing npm:", err)
		return true
	}

	fmt.Println(string(output))

	fmt.Println("Installing Packages")
	result, err := installPackages([]string{})

	if !result && err != nil {
		fmt.Println("Error installing packages:", err)
		return true
	}

	fmt.Println("Installing Dev Packages")
	devResult, devErr := installDevPackages([]string{})

	if !devResult && devErr != nil {
		fmt.Println("Error installing dev packages:", devErr)
		return true
	}

	fmt.Println("Creating src directory")

	if mkdirErr := os.Mkdir("src", 0755); mkdirErr != nil {
		fmt.Println("Error creating src directory:", err)
		return true
	}

	fmt.Println("Creating uploads directory")
	if mkdirErr := os.Mkdir("uploads", 0755); mkdirErr != nil {
		fmt.Println("Error creating uploads directory:", err)
		return true
	}

	copyContent("../../templates/index.ts", "src/index.ts")
	copyContent("../../templates/tsconfig.txt", "tsconfig.json")

	updatePackageJson()

	os.Chdir("../..")
	return true
}

func copyContent(source string, destination string) bool {
	fmt.Printf("Creating %s\n", destination)

	templatePath := filepath.Join(source)

	content, err := os.ReadFile(templatePath)

	if err != nil {
		fmt.Println("Error reading template file:", err)
		return false
	}

	newFilePath := destination
	writeErr := os.WriteFile(newFilePath, content, 0644)

	if writeErr != nil {
		fmt.Println("Error writing to file:", err)
		return false
	}

	return true
}

func installPackages(packageNames []string) (bool, error) {
	for _, packageName := range packageNames {
		cmd := exec.Command("npm", "install", packageName)
		_, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Println("Error installing package:", packageName, err)
			return false, err
		}

		fmt.Printf("Successfully installed %s\n", packageName)
	}

	return true, nil
}

func installDevPackages(packageNames []string) (bool, error) {
	for _, packageName := range packageNames {
		cmd := exec.Command("npm", "install", "--save-dev", packageName)
		_, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Println("Error installing package:", packageName, err)
			return false, err
		}

		fmt.Printf("Successfully installed %s\n", packageName)
	}

	return true, nil
}

func updatePackageJson() {
	fmt.Println("Updating package.json")

	packageJsonPath := "package.json"

	packageJSON, err := os.ReadFile(packageJsonPath)

	if err != nil {
		fmt.Println("Error reading package.json:", err)
		return
	}

	var pkg map[string]interface{}

	if unMarshalErr := json.Unmarshal(packageJSON, &pkg); unMarshalErr != nil {
		fmt.Println("Error parsing package.json")
		return
	}

	scripts := map[string]string{}

	if pkg["scripts"] == nil {
		pkg["scripts"] = scripts
	} else {
		existingScripts := pkg["scripts"].(map[string]interface{})
		for key, value := range scripts {
			existingScripts[key] = value
		}
		pkg["scripts"] = existingScripts
	}

	updatedJson, err := json.MarshalIndent(pkg, "", "  ")

	if err != nil {
		fmt.Println("Error marshaling package.json: ", err)
		return
	}

	if err := os.WriteFile(packageJsonPath, updatedJson, 0644); err != nil {
		fmt.Println("Error writing package.json: ", err)
		return
	}
}
