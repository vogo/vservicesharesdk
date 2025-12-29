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

package accounts

import (
	"encoding/json"
	"fmt"

	"github.com/vogo/vservicesharesdk/cores"
)

// BalanceQueryRequest represents the request for querying account balance.
type BalanceQueryRequest struct {
	// ProviderID is the service provider ID (required)
	ProviderID int64 `json:"providerId"`

	// PaymentType is the account type (optional)
	// 0 = Bank Card, 1 = Alipay, 2 = WeChat
	PaymentType cores.PaymentType `json:"paymentType,omitempty"`
}

// BalanceQueryResponse represents the response for balance query.
type BalanceQueryResponse struct {
	// Balance is the account balance in fen (分)
	// Note: 1 yuan = 100 fen
	Balance int64 `json:"balance"`

	// ProviderID is the service provider ID
	ProviderID int64 `json:"providerId"`
}

// BalanceQuery queries the merchant account balance.
func (s *Service) BalanceQuery(req *BalanceQueryRequest) (*BalanceQueryResponse, error) {
	// Validate request
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.ProviderID == 0 {
		return nil, fmt.Errorf("providerId is required")
	}

	// Call API with function code 6003
	respData, err := s.client.Do(cores.FunCodeBalanceQuery, req)
	if err != nil {
		return nil, err
	}

	// Handle empty response
	if respData == "" {
		return nil, fmt.Errorf("empty response data")
	}

	// Unmarshal decrypted response
	var resp BalanceQueryResponse
	if err := json.Unmarshal([]byte(respData), &resp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &resp, nil
}
