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

package payments

import (
	"encoding/json"
	"fmt"

	"github.com/vogo/vservicesharesdk/cores"
)

// PaymentQueryItem represents a query filter for specific orders.
type PaymentQueryItem struct {
	MerOrderId string `json:"merOrderId,omitempty"` // merchant order ID
	OrderNo    string `json:"orderNo,omitempty"`    // platform order number
}

// PaymentQueryRequest represents the request for querying batch payment status.
type PaymentQueryRequest struct {
	MerBatchId string             `json:"merBatchId"`           // merchant batch number
	QueryItems []PaymentQueryItem `json:"queryItems,omitempty"` // query items
}

// PaymentQuery retrieves payment status for batch transactions.
//
// IMPORTANT NOTES:
// - Omitting QueryItems returns all orders in the batch
// - Error codes 6000 or 6042 indicate communication issues only, NOT transaction failures
// - Always use OrderNo as the primary transaction identifier to prevent duplicate processing
func (s *Service) PaymentQuery(req *PaymentQueryRequest) (*PaymentBatchResult, error) {
	// Validate request
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.MerBatchId == "" {
		return nil, fmt.Errorf("merBatchId is required")
	}

	// Call API with function code 6002
	respData, err := s.client.Do(cores.FunCodePaymentQuery, req)
	if err != nil {
		return nil, err
	}

	// Handle empty response
	if respData == "" {
		return nil, fmt.Errorf("empty response data")
	}

	// Unmarshal decrypted response
	var resp PaymentBatchResult
	if err := json.Unmarshal([]byte(respData), &resp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &resp, nil
}
