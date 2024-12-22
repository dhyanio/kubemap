package terraform

import (
	"encoding/json"
	"fmt"
	"os"
)

type Node struct {
	ID       string   `json:"id"`       // Resource name
	Type     string   `json:"type"`     // Resource type (e.g., aws_instance)
	Provider string   `json:"provider"` // Provider (e.g., aws)
	Depends  []string `json:"depends"`  // Dependencies
}

type Graph struct {
	Nodes []Node `json:"nodes"`
}

// ParseTfState parses a Terraform state file and returns a list of resources as nodes.
func ParseTfState(filePath, account, region string) ([]Node, error) {
	// Read the Terraform state file
	stateFile, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read state file: %w", err)
	}

	// Parse the JSON
	var stateData map[string]interface{}
	if err := json.Unmarshal(stateFile, &stateData); err != nil {
		return nil, fmt.Errorf("failed to parse state file JSON: %w", err)
	}

	// Extract resources
	nodes := []Node{}
	modules, ok := stateData["modules"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid or missing 'modules' field in state file")
	}

	for _, mod := range modules {
		module, ok := mod.(map[string]interface{})
		if !ok {
			continue
		}

		resources, ok := module["resources"].(map[string]interface{})
		if !ok {
			continue
		}

		for resourceName, resourceData := range resources {
			resourceMap, ok := resourceData.(map[string]interface{})
			if !ok {
				continue
			}

			resourceType, ok := resourceMap["type"].(string)
			if !ok {
				continue
			}

			primary, ok := resourceMap["primary"].(map[string]interface{})
			if !ok {
				continue
			}

			resourceID, ok := primary["id"].(string)
			if !ok {
				continue
			}

			depends, _ := primary["depends_on"].([]string) // Safe to ignore missing depends_on

			node := Node{
				ID:       resourceID,
				Type:     resourceType,
				Provider: "aws", // Assuming AWS for simplicity
				Depends:  depends,
			}
			fmt.Printf("Parsed resource: %s\n", resourceName)
			nodes = append(nodes, node)
		}
	}

	return nodes, nil
}
