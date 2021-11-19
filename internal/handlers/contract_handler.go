package handlers

import (
	"fmt"

	"github.com/latonaio/salesforce-data-models"
	"github.com/latonaio/aion-core/pkg/log"
)

func HandleContract(metadata map[string]interface{}) error {
	contracts, err := models.MetadataToContracts(metadata)
	if err != nil {
		return fmt.Errorf("failed to convert contracts: %v", err)
	}
	for _, contract := range contracts {
		if contract.SfContractID == nil {
			log.Print("SfContractID is null. skip")
			continue
		}
		c, err := models.ContractByID(*contract.SfContractID)
		if err != nil {
			log.Printf("failed to get contract: %v", err)
			continue
		}
		if c != nil {
			log.Printf("update contract: %s\n", *contract.SfContractID)
			if err := contract.Update(); err != nil {
				log.Printf("failed to update contract: %v", err)
				continue
			}
		} else {
			log.Printf("register contract: %s\n", *contract.SfContractID)
			if err := contract.Register(); err != nil {
				log.Printf("failed to register contract: %v", err)
			}
		}
	}
	return nil
}
