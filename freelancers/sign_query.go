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
	// FunCodeSignQuery is the function code for contract status query API
	FunCodeSignQuery = "6011"
)

// SignState represents the sign status.
type SignState int

const (
	// SignStateUnsigned indicates the freelancer has not signed
	SignStateUnsigned SignState = 0
	// SignStateSigned indicates the freelancer has signed
	SignStateSigned SignState = 1
	// SignStateNotFound indicates the freelancer record was not found
	SignStateNotFound SignState = 2
	// SignStatePending indicates the sign is pending
	SignStatePending SignState = 3
	// SignStateFailed indicates the sign failed
	SignStateFailed SignState = 4
	// SignStateCancelled indicates the sign was cancelled
	SignStateCancelled SignState = 5
)

// SignQueryRequest represents the request for querying freelancer sign status.
type SignQueryRequest struct {
	// Name is the freelancer's name (required, max 25 chars)
	Name string `json:"name"`

	// IdCard is the ID card number (required, max 18 chars)
	IdCard string `json:"idCard"`

	// Mobile is the phone number registered with bank (required, 11 chars)
	Mobile string `json:"mobile"`

	// ProviderId is the service provider ID (required, max 5 digits)
	ProviderId string `json:"providerId"`
}

// SignQueryResponse represents the response for contract query.
type SignQueryResponse struct {
	// Name is the freelancer's name
	Name string `json:"name"`

	// CardNo is the bank card number or payment account
	CardNo string `json:"cardNo"`

	// IdCard is the ID card number
	IdCard string `json:"idCard"`

	// Mobile is the phone number
	Mobile string `json:"mobile"`

	// State is the sign status
	State SignState `json:"state"`

	// ProviderId is the service provider ID
	ProviderId string `json:"providerId"`

	// RetMsg is the failure reason if applicable
	RetMsg string `json:"retMsg,omitempty"`
}

// QuerySign queries the sign status of a freelancer.
//
// Note: After changing bank cards, no need to re-sign.
func (s *Service) QuerySign(req *SignQueryRequest) (*SignQueryResponse, error) {
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
	respData, err := s.client.Do(FunCodeSignQuery, req)
	if err != nil {
		return nil, fmt.Errorf("query contract failed: %w", err)
	}

	// Handle empty response
	if respData == "" {
		return nil, fmt.Errorf("empty response data")
	}

	// Unmarshal decrypted response
	var resp SignQueryResponse
	if err := json.Unmarshal([]byte(respData), &resp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &resp, nil
}
