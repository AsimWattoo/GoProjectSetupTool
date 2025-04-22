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
	rootDir, getWdErr := os.Getwd()

	if getWdErr != nil {
		fmt.Printf("Error getting the current directory. Error: %s\n", getWdErr)
		return true
	}

	executor := commands.NewCommandExecutor()

	executor.AddCommand("mkdir", directoryPath, "0755")
	executor.AddCommand("cd", directoryPath)
	executor.AddRawCommand("npm", "init", "-y")
	executor.AddRawCommand("npm", "install", "express", "dotenv", "cookie-parser", "tsconfig-paths", "cors", "express-session", "moment")
	executor.AddRawCommand("npm", "install", "@types/node", "@types/express", "ts-node", "typescript", "nodemon", "@types/cors", "@types/cookie-parser", "@types/express-session")
	executor.AddCommand("mkdir", "src", "0755")
	executor.AddCommand("mkdir", "uploads", "0755")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/node/index.ts"), "src/index.ts", "--replace", "// @ts-nocheck", "")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/node/tsconfig.txt"), "tsconfig.json", "--replace", "// @ts-nocheck", "")
	executor.AddCommand("json-update", "package.json", "--data", "scripts", "{\"start\": \"node src/index.ts\", \"dev\": \"nodemon --exec ts-node src/index.ts\"}")
	executor.AddCommand("mkdir", "src/middlewares", "0755")
	executor.AddCommand("mkdir", "src/data", "0755")
	executor.AddCommand("mkdir", "src/services", "0755")
	executor.AddCommand("mkdir", "src/settings", "0755")
	executor.AddCommand("mkdir", "src/types", "0755")
	executor.AddCommand("mkdir", "src/types/errors", "0755")
	executor.AddCommand("mkdir", "src/types/generic", "0755")
	executor.AddCommand("mkdir", "src/routes", "0755")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/node/middlewares/error-handler.ts"), "src/middlewares/error-handler.ts", "--replace", "// @ts-nocheck", "")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/node/middlewares/request-logger.ts"), "src/middlewares/request-logger.ts", "--replace", "// @ts-nocheck", "")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/node/data/DateFormats.ts"), "src/data/DateFormats.ts", "--replace", "// @ts-nocheck", "")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/node/services/ErrorService.ts"), "src/services/ErrorService.ts", "--replace", "// @ts-nocheck", "")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/node/services/Logger.ts"), "src/services/Logger.ts", "--replace", "// @ts-nocheck", "")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/node/settings/errors-config.ts"), "src/settings/errors-config.ts", "--replace", "// @ts-nocheck", "")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/node/types/errors/APIError.ts"), "src/types/errors/APIError.ts", "--replace", "// @ts-nocheck", "")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/node/types/errors/ErrorResponse.ts"), "src/types/errors/ErrorResponse.ts", "--replace", "// @ts-nocheck", "")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/node/types/generic/ApiResponse.ts"), "src/types/generic/ApiResponse.ts", "--replace", "// @ts-nocheck", "")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/node/types/generic/ValidationResult.ts"), "src/types/generic/ValidationResult.ts", "--replace", "// @ts-nocheck", "")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/node/routes/index.ts"), "src/routes/index.ts", "--replace", "// @ts-nocheck", "")

	executor.AddCommand("cd", rootDir)

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
