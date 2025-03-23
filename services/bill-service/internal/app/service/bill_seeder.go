package service

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/mehdihadeli/go-mediatr"
	"github.com/zelalem-12/bill-aggregation-system_onetab/bill-service/internal/app/command/createbill"
)

func readProviderData(filename string) ([]*createbill.CreateBillCommand, error) {

	filePath, err := GetRootFilePath(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to get file path: %v", err)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	var commands []*createbill.CreateBillCommand
	if err := json.Unmarshal(data, &commands); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return commands, nil
}

func SeedSampleBillData() {
	commands, err := readProviderData("sample_bill_data.json")
	if err != nil {
		fmt.Printf("failed to read bill data: %v", err)
		return
	}

	for _, cmd := range commands {
		if err := cmd.Validate(); err != nil {
			fmt.Printf("bill seeding validate command failed: %v", err)
			return
		}

		_, err := mediatr.Send[*createbill.CreateBillCommand, *createbill.CreateBillCommandResponse](context.Background(), cmd)
		if err != nil {
			fmt.Printf("bill seeding failed: %v", err)
			return
		}
	}
	fmt.Println("Bill Seeding Completed")

}
