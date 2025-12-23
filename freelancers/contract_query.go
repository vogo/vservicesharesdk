/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package freelancers

import (
	"encoding/json"
	"fmt"
)

const (
	// FunCodeContractQuery is the function code for contract status query API
	FunCodeContractQuery = "6011"
)

// ContractState represents the contract status.
type ContractState int

const (
	// ContractStateUnsigned indicates the freelancer has not signed
	ContractStateUnsigned ContractState = 0
	// ContractStateSigned indicates the freelancer has signed
	ContractStateSigned ContractState = 1
	// ContractStateNotFound indicates the freelancer record was not found
	ContractStateNotFound ContractState = 2
	// ContractStatePending indicates the contract is pending
	ContractStatePending ContractState = 3
	// ContractStateFailed indicates the contract signing failed
	ContractStateFailed ContractState = 4
	// ContractStateCancelled indicates the contract was cancelled
	ContractStateCancelled ContractState = 5
)

// ContractQueryRequest represents the request for querying freelancer contract status.
type ContractQueryRequest struct {
	// Name is the freelancer's name (required, max 25 chars)
	Name string `json:"name"`

	// IdCard is the ID card number (required, max 18 chars)
	IdCard string `json:"idCard"`

	// Mobile is the phone number registered with bank (required, 11 chars)
	Mobile string `json:"mobile"`

	// ProviderId is the service provider ID (required, max 5 digits)
	ProviderId string `json:"providerId"`
}

// ContractQueryResponse represents the response for contract query.
type ContractQueryResponse struct {
	// Name is the freelancer's name
	Name string `json:"name"`

	// CardNo is the bank card number or payment account
	CardNo string `json:"cardNo"`

	// IdCard is the ID card number
	IdCard string `json:"idCard"`

	// Mobile is the phone number
	Mobile string `json:"mobile"`

	// State is the contract status
	// 0: unsigned, 1: signed, 2: not found, 3: pending, 4: failed, 5: cancelled
	State ContractState `json:"state"`

	// ProviderId is the service provider ID
	ProviderId string `json:"providerId"`

	// RetMsg is the failure reason if applicable
	RetMsg string `json:"retMsg,omitempty"`
}

// QueryContract queries the contract status of a freelancer.
//
// Note: After changing bank cards, no need to re-contract.
func (s *Service) QueryContract(req *ContractQueryRequest) (*ContractQueryResponse, error) {
	// Validate request
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.Name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if req.IdCard == "" {
		return nil, fmt.Errorf("idCard is required")
	}
	if req.Mobile == "" {
		return nil, fmt.Errorf("mobile is required")
	}
	if req.ProviderId == "" {
		return nil, fmt.Errorf("providerId is required")
	}

	// Call API with function code 6011
	respData, err := s.client.Do(FunCodeContractQuery, req)
	if err != nil {
		return nil, fmt.Errorf("query contract failed: %w", err)
	}

	// Handle empty response
	if respData == "" {
		return nil, fmt.Errorf("empty response data")
	}

	// Unmarshal decrypted response
	var resp ContractQueryResponse
	if err := json.Unmarshal([]byte(respData), &resp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &resp, nil
}
