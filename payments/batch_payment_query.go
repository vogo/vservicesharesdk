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
)

const (
	// FunCodeBatchPaymentQuery is the function code for batch payment query API
	FunCodeBatchPaymentQuery = "6002"
)

// PaymentState represents the payment transaction state.
type PaymentState int

const (
	// PaymentStateProcessing indicates the payment is being processed
	PaymentStateProcessing PaymentState = 1
	// PaymentStateSuccess indicates the payment succeeded
	PaymentStateSuccess PaymentState = 3
	// PaymentStateFailed indicates the payment failed
	PaymentStateFailed PaymentState = 4
	// PaymentStatePendingConfirm indicates awaiting user confirmation
	PaymentStatePendingConfirm PaymentState = 6
	// PaymentStateCancelled indicates the payment was cancelled
	PaymentStateCancelled PaymentState = 7
)

// QueryItem represents a query filter for specific orders.
type QueryItem struct {
	// MerOrderId is the merchant order ID (optional)
	MerOrderId string `json:"merOrderId,omitempty"`

	// OrderNo is the platform order number (optional)
	OrderNo string `json:"orderNo,omitempty"`
}

// BatchPaymentQueryRequest represents the request for querying batch payment status.
type BatchPaymentQueryRequest struct {
	// MerBatchId is the batch number to query (required, max 32 chars)
	MerBatchId string `json:"merBatchId"`

	// QueryItems specifies which orders to query (optional)
	// If omitted, returns all orders in the batch
	QueryItems []QueryItem `json:"queryItems,omitempty"`
}

// PaymentQueryResult represents the detailed result of a payment query.
type PaymentQueryResult struct {
	// MerOrderId is the merchant order ID
	MerOrderId string `json:"merOrderId"`

	// OrderNo is the platform order number (use as primary transaction identifier)
	OrderNo string `json:"orderNo"`

	// State is the payment state
	State PaymentState `json:"state"`

	// Amt is the payment amount in fen
	Amt int64 `json:"amt"`

	// Fee is the service fee in fen
	Fee int64 `json:"fee"`

	// UserFee is the user's fee in fen
	UserFee int64 `json:"userFee"`

	// Tax is the tax amount in fen
	Tax int64 `json:"tax"`

	// UserDueAmt is the amount due to user in fen
	UserDueAmt int64 `json:"userDueAmt"`

	// ResCode is the result code
	ResCode string `json:"resCode"`

	// ResMsg is the result message
	ResMsg string `json:"resMsg"`
}

// BatchPaymentQueryResponse represents the response for batch payment query.
type BatchPaymentQueryResponse struct {
	// MerId is the merchant ID
	MerId string `json:"merId"`

	// MerBatchId is the merchant batch number
	MerBatchId string `json:"merBatchId"`

	// QueryItems contains the transaction details
	QueryItems []PaymentQueryResult `json:"queryItems"`
}

// QueryBatchPayment retrieves payment status for batch transactions.
//
// IMPORTANT NOTES:
// - Omitting QueryItems returns all orders in the batch
// - Error codes 6000 or 6042 indicate communication issues only, NOT transaction failures
// - Always use OrderNo as the primary transaction identifier to prevent duplicate processing
func (s *Service) QueryBatchPayment(req *BatchPaymentQueryRequest) (*BatchPaymentQueryResponse, error) {
	// Validate request
	if req == nil {
		return nil, fmt.Errorf("request cannot be nil")
	}
	if req.MerBatchId == "" {
		return nil, fmt.Errorf("merBatchId is required")
	}

	// Call API with function code 6002
	respData, err := s.client.Do(FunCodeBatchPaymentQuery, req)
	if err != nil {
		return nil, fmt.Errorf("query batch payment failed: %w", err)
	}

	// Handle empty response
	if respData == "" {
		return nil, fmt.Errorf("empty response data")
	}

	// Unmarshal decrypted response
	var resp BatchPaymentQueryResponse
	if err := json.Unmarshal([]byte(respData), &resp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &resp, nil
}
