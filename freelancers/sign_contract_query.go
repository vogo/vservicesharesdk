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

	"github.com/vogo/vservicesharesdk/cores"
)

// SignState represents the sign status.
type SignState int

const (
	SignStateUnsigned  SignState = 0 // freelancer has not signed
	SignStateSigned    SignState = 1 // freelancer has signed
	SignStateNotFound  SignState = 2 // freelancer record not found
	SignStatePending   SignState = 3 // sign is pending
	SignStateFailed    SignState = 4 // sign failed
	SignStateCancelled SignState = 5 // sign cancelled
)

// SignQueryRequest represents the request for querying freelancer sign status.
type SignQueryRequest struct {
	Name       string `json:"name"`       // the freelancer's name
	IdCard     string `json:"idCard"`     // the ID card number
	Mobile     string `json:"mobile"`     // the phone number registered with bank
	ProviderId int64  `json:"providerId"` // the service provider ID
}

// SignContractQuery queries the sign status of a freelancer.
//
// Note: After changing bank cards, no need to re-sign.
func (s *Service) SignContractQuery(req *SignQueryRequest) (*SignContractResult, error) {
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
	if req.ProviderId == 0 {
		return nil, fmt.Errorf("providerId is required")
	}

	// Call API with function code 6011
	respData, err := s.client.Do(cores.FunCodeSignContractQuery, req)
	if err != nil {
		return nil, err
	}

	// Handle empty response
	if respData == "" {
		return nil, fmt.Errorf("empty response data")
	}

	// Unmarshal decrypted response
	var resp SignContractResult
	if err := json.Unmarshal([]byte(respData), &resp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &resp, nil
}
