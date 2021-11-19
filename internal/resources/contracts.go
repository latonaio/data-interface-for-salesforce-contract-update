package resources

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Contract struct
type Contract struct {
	method   string
	metadata map[string]interface{}
}

func (c *Contract) objectName() string {
	const obName = "Contract"
	return obName
}

// newContract writes that new Customer instance
func NewContract(metadata map[string]interface{}) (*Contract, error) {
	rawMethod, ok := metadata["method"]
	if !ok {
		return nil, errors.New("missing required parameters: method")
	}
	method, ok := rawMethod.(string)
	if !ok {
		return nil, errors.New("failed to convert interface{} to string")
	}
	return &Contract{
		method:   method,
		metadata: metadata,
	}, nil
}

// updateMetadata mold customer update metadata
func (c *Contract) updateMetadata() (map[string]interface{}, error) {
	paramsIF, paramsOk := c.metadata["params"]
	if !paramsOk {
		return nil, errors.New("params is required")
	}
	params, ok := paramsIF.(map[string]interface{})
	if !ok {
		return nil, errors.New("failed to convert interface{} to map[string]interface{}")
	}
	contractIdIF, contractIdOk := c.metadata["contract_id"]
	if !contractIdOk {
		return nil, errors.New("contract_id is required")
	}
	contractId, ok := contractIdIF.(string)
	if !ok {
		return nil, errors.New("failed to convert contract_id to string")
	}
	params["Id"] = contractId
	accountIdIF, accountIdOk := c.metadata["account_id"]
	if !accountIdOk {
		return nil, errors.New("account_id is required")
	}
	accountId, ok := accountIdIF.(string)
	if !ok {
		return nil, errors.New("failed to convert account_id to string")
	}
	params["AccountId"] = accountId
	bytes, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("JSON marshal error: %v", err)
	}
	body := string(bytes)
	return buildMetadata(c.method, c.objectName(), contractId, nil, body), nil
}

// BuildMetadata
func (c *Contract) BuildMetadata() (map[string]interface{}, error) {
	switch c.method {
	case "put":
		return c.updateMetadata()
	}
	return nil, fmt.Errorf("invalid method: %s", c.method)
}

func buildMetadata(method, object, pathParam string, queryParams map[string]string, body string) map[string]interface{} {
	metadata := map[string]interface{}{
		"method":         method,
		"object":         object,
		"connection_key": "contract_put",
	}
	if len(pathParam) > 0 {
		metadata["path_param"] = pathParam
	}
	if queryParams != nil {
		metadata["query_params"] = queryParams
	}
	if body != "" {
		metadata["body"] = body
	}
	return metadata
}
