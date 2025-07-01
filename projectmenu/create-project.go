package projectmenu

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/AsimWattoo/go-utilities/command"
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

	executor := command.NewCommandExecutor()

	executor.AddCommand("mkdir", directoryPath, "0755")
	executor.AddCommand("cd", directoryPath)
	executor.AddRawCommand("npm", true, "init", "-y")
	executor.AddRawCommand("npm", true, "install", "express", "dotenv", "cookie-parser", "tsconfig-paths", "cors", "express-session", "moment", "express-rate-limit")
	executor.AddRawCommand("npm", true, "install", "@types/node", "@types/express", "ts-node", "typescript", "nodemon", "@types/cors", "@types/cookie-parser", "@types/express-session")
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

	executor := command.NewCommandExecutor()

	executor.AddRawCommand("npm", true, "create", "vite@latest", projectName, "--", "--template", "react-ts")
	executor.AddCommand("cd", projectName)
	executor.AddRawCommand("npm", true, "install")
	executor.AddRawCommand("npm", true, "install", "@tanstack/react-router")
	executor.AddRawCommand("npm", true, "install", "-D", "tailwindcss@3", "postcss", "autoprefixer", "tsx", "@types/node", "postcss-cli", "postcss-import", "sass-embedded")
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

func CreateNodeLibraryProject() bool {

    fmt.Printf("Enter name of the project: ")
    var projectName string;
    var packageName string;

    _, inputErr := fmt.Scanln(&projectName)

    if inputErr != nil {
        fmt.Printf("Error reading project name: %s\n", inputErr)
        return true
    }

    fmt.Printf("Enter Package Name (Without Spaces): ")
    _, packageInputErr := fmt.Scanln(&packageName);

    if packageInputErr != nil {
        fmt.Printf("Error reading package name: %s\n", packageInputErr)
        return true;
    }

    var lowercasePackageName = strings.ToLower(packageName)
    directoryPath := filepath.Join(projectName)

    rootDir, getWdErr := getExecutableDir()

    if getWdErr != nil {
        fmt.Printf("Error getting the current directory. Error: %s\n", getWdErr)
        return true
    }

    executor := command.NewCommandExecutor()

    executor.AddCommand("mkdir", directoryPath, "0755")
    executor.AddCommand("cd", directoryPath)
    executor.AddRawCommand("npm", true, "init", "-y")
    executor.AddRawCommand("yarn", true, "add", "dotenv")
    executor.AddRawCommand("yarn", true, "add", "--dev", "@types/node", "ts-node", "typescript", "@types/chai", "@types/supertest", "chai", "supertest", "jsdom", "vite", "vitest", "vite-plugin-dts")
    executor.AddCommand("mkdir", "src", "0755")
    executor.AddCommand("mkdir", "test", "0755")
    executor.AddCommand("json-update", "package.json", "--data", "scripts", "{\"publish\": \"npm publish\", \"build\": \"tsc -b tsconfig.lib.json & vite build\", \"test\": \"vitest --run --config vitest.config.ts\"}");
    executor.AddCommand("json-update", "package.json", "--data", "name", fmt.Sprintf("@axontick/%s", lowercasePackageName))
    executor.AddCommand("json-update", "package.json", "--data", "description", "<DESCRIPTION>")
    executor.AddCommand("json-update", "package.json", "--data", "author", "Axontick", "--data", "type", "module")
    executor.AddCommand("json-update", "package.json", "--data", "main", fmt.Sprintf("./dist/%s.umd.cjs", lowercasePackageName), "--data", "module", fmt.Sprintf("./dist/%s.js", lowercasePackageName))
    executor.AddCommand("json-update", "package.json", "--data", "types", fmt.Sprintf("./dist/%s.d.ts", lowercasePackageName))
    executor.AddCommand("json-update", "package.json", "--data", "exports", fmt.Sprintf("{\".\": {\"types\": \"./dist/%s.d.ts\",\"import\": \"./dist/%s.js\",\"require\": \"./dist/%s.umd.cjs\"}}", lowercasePackageName, lowercasePackageName, lowercasePackageName))
    executor.AddCommand("json-update", "package.json", "--data", "files", "[\"dist\"]", "publishConfig", "{ \"registry\": \"https://npm.pkg.github.com\"}");
    executor.AddCommand("cp", filepath.Join(rootDir, "templates/library/index.ts"), "src/index.ts", "--replace", "// @ts-nocheck", "")
    executor.AddCommand("cp", filepath.Join(rootDir, "templates/library/tsconfig.json"), "tsconfig.json")
    executor.AddCommand("cp", filepath.Join(rootDir, "templates/library/tsconfig.lib.json"), "tsconfig.lib.json")
    executor.AddCommand("cp", filepath.Join(rootDir, "templates/library/tsconfig.test.json"), "tsconfig.test.json")
    executor.AddCommand("cp", filepath.Join(rootDir, "templates/library/vitest.config.ts"), "vitest.config.ts", "--replace", "// @ts-nocheck", "");
    executor.AddCommand("cp", filepath.Join(rootDir, "templates/library/vite.config.ts"), "vite.config.ts", "--replace", "// @ts-nocheck", "", "--replace", "<PACKAGE_NAME>", lowercasePackageName);
    executor.AddCommand("cp", filepath.Join(rootDir, "templates/library/.gitignore"), ".gitignore")
    executor.AddCommand("cp", filepath.Join(rootDir, "templates/library/.env"), ".env")

    executor.AddCommand("cd", rootDir)

    err := executor.ExecuteCommands()

    if err != nil {
        fmt.Printf("==> Command Execution Failed because of following error: %s\n", err)
    } else {
        fmt.Printf("==> Command Execution Successful\n")
    }

    return true
}