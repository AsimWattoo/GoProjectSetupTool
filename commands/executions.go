package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var GlobalExecutions = map[string]func(args []string) error{
	"mkdir": func(args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("Missing required params")
		}

		permissions, err := strconv.ParseUint(args[1], 8, 32)

		if err != nil {
			return err
		}

		return os.Mkdir(args[0], os.FileMode(permissions))
	},
	"cd": func(args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("Missing required params")
		}
		return os.Chdir(args[0])
	},
	"cp": func(args []string) error {

		if len(args) < 2 {
			return fmt.Errorf("Missing required params")
		}

		content, err := os.ReadFile(args[0])

		if err != nil {
			return fmt.Errorf("Error reading file. Error: %s", err)
		}

		contentStr := string(content)

		if len(args) > 2 {
			for i := 2; i < len(args); i += 1 {
				command := args[i]

				if command == "--replace" {
					oldText := args[i+1]
					newText := args[i+2]

					contentStr = strings.ReplaceAll(contentStr, oldText, newText)

					i += 2
				}
			}
		}

		writeErr := os.WriteFile(args[1], []byte(contentStr), 0644)

		if writeErr != nil {
			return fmt.Errorf("Error writing to file. Error: %s", err)
		}

		return nil
	},
	"json-update": func(args []string) error {
		if len(args) < 3 {
			return fmt.Errorf("Missing required params")
		}

		jsonPath := args[0]

		jsonContent, err := os.ReadFile(jsonPath)

		if err != nil {
			return fmt.Errorf("Error reading json file. Error: %s", err)
		}

		var jsonData map[string]interface{}

		if unMarshalErr := json.Unmarshal(jsonContent, &jsonData); unMarshalErr != nil {
			return fmt.Errorf("Error parsing json file. Error: %s", unMarshalErr)
		}

		for i := 1; i < len(args); i += 1 {
			if args[i] == "--data" {
				block := args[i+1]
				data := args[i+2]
				blockData := map[string]string{}

				if jsonParseErr := json.Unmarshal([]byte(data), &blockData); jsonParseErr != nil {
					return fmt.Errorf("Error parsing data: %v", jsonParseErr)
				}

				if jsonData[block] == nil {
					jsonData[block] = blockData
				} else {
					existingData := jsonData[block].(map[string]interface{})
					for key, value := range blockData {
						existingData[key] = value
					}
					jsonData[block] = existingData
				}

				i += 2
			}
		}

		updatedJson, err := json.MarshalIndent(jsonData, "", "  ")

		if err != nil {
			return fmt.Errorf("Error marshaling json file. Error: %s", err)
		}

		if err := os.WriteFile(jsonPath, updatedJson, 0644); err != nil {
			return fmt.Errorf("Error writing json file. Error %s", err)
		}

		return nil
	},
}
