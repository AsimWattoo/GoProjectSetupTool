package menu

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"tool/node-backend-project/commands"
)

func getExecutableDir() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	exeDir := filepath.Dir(exePath)

	tempDir := os.TempDir()

	if strings.HasPrefix(exeDir, tempDir) {
		return os.Getwd()
	}

	return exeDir, nil
}

func CreateNodeTsProject() bool {

	fmt.Printf("Enter name of the project: ")
	var projectName string
	_, inputErr := fmt.Scanln(&projectName)

	if inputErr != nil {
		fmt.Printf("Error reading project name: %s", inputErr)
		return true
	}

	directoryPath := filepath.Join(projectName)

	rootDir, getWdErr := getExecutableDir()

	if getWdErr != nil {
		fmt.Printf("Error getting the current directory. Error: %s\n", getWdErr)
		return true
	}

	executor := commands.NewCommandExecutor()

	executor.AddCommand("mkdir", directoryPath, "0755")
	executor.AddCommand("cd", directoryPath)
	executor.AddRawCommand("npm", "init", "-y")
	executor.AddRawCommand("npm", "install", "express", "dotenv", "cookie-parser", "tsconfig-paths", "cors", "express-session", "moment", "express-rate-limit")
	executor.AddRawCommand("npm", "install", "@types/node", "@types/express", "ts-node", "typescript", "nodemon", "@types/cors", "@types/cookie-parser", "@types/express-session")
	executor.AddCommand("mkdir", "src", "0755")
	executor.AddCommand("mkdir", "uploads", "0755")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/node/index.ts"), "src/index.ts", "--replace", "// @ts-nocheck", "")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/node/tsconfig.txt"), "tsconfig.json", "--replace", "// @ts-nocheck", "")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/node/.env"), ".env", "--replace", "// @ts-nocheck", "")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/node/gitignore.txt"), ".gitignore", "--replace", "// @ts-nocheck", "")
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
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/node/middlewares/rate-limiter.ts"), "src/middlewares/rate-limiter.ts", "--replace", "// @ts-nocheck", "")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/node/middlewares/request-logger.ts"), "src/middlewares/request-logger.ts", "--replace", "// @ts-nocheck", "")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/node/data/DateFormats.ts"), "src/data/DateFormats.ts", "--replace", "// @ts-nocheck", "")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/node/services/ErrorService.ts"), "src/services/ErrorService.ts", "--replace", "// @ts-nocheck", "")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/node/services/Logger.ts"), "src/services/Logger.ts", "--replace", "// @ts-nocheck", "")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/node/settings/errors-config.ts"), "src/settings/errors-config.ts", "--replace", "// @ts-nocheck", "")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/node/settings/config.ts"), "src/settings/config.ts", "--replace", "// @ts-nocheck", "")
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

func CreateViteTsProject() bool {
	fmt.Printf("Enter name of the project: ")
	var projectName string = ""
	_, inputErr := fmt.Scanln(&projectName)

	if inputErr != nil {
		fmt.Printf("Error reading project name: %s", inputErr)
		return true
	}

	rootDir, getWdErr := getExecutableDir()

	if getWdErr != nil {
		fmt.Printf("Error getting the current directory. Error: %s\n", getWdErr)
		return true
	}

	executor := commands.NewCommandExecutor()

	executor.AddRawCommand("npm", "create", "vite@latest", projectName, "--", "--template", "react-ts")
	executor.AddCommand("cd", projectName)
	executor.AddRawCommand("npm", "install")
	executor.AddRawCommand("npm", "install", "@tanstack/react-router")
	executor.AddRawCommand("npm", "install", "-D", "tailwindcss@3", "postcss", "autoprefixer", "tsx", "@types/node", "postcss-cli", "postcss-import", "sass-embedded")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/react/tsconfig.json"), "tsconfig.json")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/react/tsconfig.app.json"), "tsconfig.app.json")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/react/vite.config.ts"), "vite.config.ts", "--replace", "// @ts-nocheck", "")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/react/postcss.config.ts"), "postcss.config.ts")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/react/tailwind.config.ts"), "tailwind.config.ts")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/react/src/index.scss"), "src/index.scss")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/react/src/App.scss"), "src/App.scss")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/react/src/main.tsx"), "src/main.tsx", "--replace", "// @ts-nocheck", "")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/react/src/App.tsx"), "src/App.tsx", "--replace", "// @ts-nocheck", "")

	executor.AddCommand("mkdir", "src/routes", "0755")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/react/src/routes/base-routes.tsx"), "src/routes/base-routes.tsx", "--replace", "// @ts-nocheck", "")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/react/src/routes/root-route.tsx"), "src/routes/root-route.tsx", "--replace", "// @ts-nocheck", "")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/react/src/routes/root.tsx"), "src/routes/root.tsx", "--replace", "// @ts-nocheck", "")

	executor.AddCommand("mkdir", "src/pages", "0755")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/react/src/pages/index.tsx"), "src/pages/index.tsx", "--replace", "// @ts-nocheck", "")

	executor.AddCommand("mkdir", "src/layouts", "0755")
	executor.AddCommand("cp", filepath.Join(rootDir, "templates/react/src/layouts/root-layout.tsx"), "src/layouts/root-layout.tsx", "--replace", "// @ts-nocheck", "")

	executor.AddCommand("cd", rootDir)

	err := executor.ExecuteCommands()

	if err != nil {
		fmt.Printf("==> Command Execution Failed because of following error: %s\n", err)
	} else {
		fmt.Printf("==> Command Execution Successful\n")
	}

	return true
}
