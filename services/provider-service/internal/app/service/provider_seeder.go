package service

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/mehdihadeli/go-mediatr"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/app/command/addprovider"
)

func readProviderData(filename string) ([]*addprovider.AddProviderCommand, error) {

	filePath, err := GetRootFilePath(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to get file path: %v", err)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	var commands []*addprovider.AddProviderCommand
	if err := json.Unmarshal(data, &commands); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return commands, nil
}

func SeedutilityProviders() {
	commands, err := readProviderData("utility_provider_data.json")
	if err != nil {
		fmt.Printf("failed to read provider data: %v", err)
		return
	}

	for _, cmd := range commands {
		if err := cmd.Validate(); err != nil {
			fmt.Printf("Provider Seeding Validate Command Failed: %v", err)
			return
		}

		_, err := mediatr.Send[*addprovider.AddProviderCommand, *addprovider.AddProviderCommandResponse](context.Background(), cmd)
		if err != nil {
			fmt.Printf("Provider Seeding Failed: %v", err)
			return
		}
	}
	fmt.Println("Provider Seeding Completed")

}
