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

	"github.com/vogo/vogo/vlog"
)

// SignCallback represents the data in the contract signing callback.
// Note: Fields are based on typical patterns. Please verify with actual documentation.
type SignCallback struct {
	// FlowId is the unique flow ID of the signing process
	FlowId string `json:"flowId"`
	// AccountId is the account ID of the freelancer
	AccountId string `json:"accountId"`
	// ContractId is the ID of the signed contract
	ContractId string `json:"contractId"`
	// Status is the signing status (e.g., "1" for success, "2" for fail)
	Status string `json:"status"`
	// SignTime is the time when the contract was signed
	SignTime string `json:"signTime"`
	// ErrMsg is the error message if signing failed
	ErrMsg string `json:"errMsg"`
	// ThirdTemplateId is the template ID used
	ThirdTemplateId string `json:"thirdTemplateId"`
}

// ParseSignCallback parses and validates the contract signing callback request.
// It takes the raw JSON body of the callback request.
func (s *Service) ParseSignCallback(body []byte) (*SignCallback, error) {
	// Verify and decrypt the notification
	decryptedData, err := s.client.VerifyAndDecryptNotification(body)
	if err != nil {
		return nil, fmt.Errorf("failed to verify and decrypt callback: %w", err)
	}

	vlog.Infof("service share contract sign callback | data: %s", decryptedData)

	if decryptedData == "" {
		return nil, fmt.Errorf("empty callback data")
	}

	// Unmarshal decrypted data
	var callback SignCallback
	if err := json.Unmarshal([]byte(decryptedData), &callback); err != nil {
		return nil, fmt.Errorf("failed to parse callback data: %w", err)
	}

	return &callback, nil
}
